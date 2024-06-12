package auth

import (
	"fmt"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(c *fiber.Ctx, request RegisterRequest) RegisterResponse
	Login(c *fiber.Ctx, request LoginRequest) LoginResponse
	Token(c *fiber.Ctx, request TokenRequest) TokenResponse
	ForgotPassword(c *fiber.Ctx, request ForgotPasswordRequest)
	ResetPassword(c *fiber.Ctx, request ResetPasswordRequest)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) Register(c *fiber.Ctx, request RegisterRequest) RegisterResponse {
	ctx := c.Context()

	_, err := service.repository.GetUserByEmail(c.Context(), request.Email)
	helper.PanicErrIfNotErr(err, ErrDuplicateEmail)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfErr(err)

	roleCustomerId, _ := uuid.Parse(constant.RoleCustomerId)

	now := time.Now()
	user := domain.User{
		Id:        uuid.New(),
		Email:     request.Email,
		Password:  string(hashedPassword),
		Name:      request.Name,
		RoleId:    roleCustomerId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = service.repository.CreateUser(ctx, user)
	helper.PanicIfErr(err)

	permissions, err := service.repository.GetPermissionsByRoleId(ctx, user.RoleId)
	helper.PanicIfErr(err)

	accessToken, err := helper.GenerateJwt(user, permissions, nil)
	helper.PanicIfErr(err)

	refreshToken, err := helper.GenerateRefreshJwt(user, nil)
	helper.PanicIfErr(err)

	response := RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response
}

func (service service) Login(c *fiber.Ctx, request LoginRequest) LoginResponse {
	ctx := c.Context()

	user, err := service.repository.GetUserByEmail(ctx, request.Email)
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	permissions, err := service.repository.GetPermissionsByRoleId(ctx, user.RoleId)
	helper.PanicIfErr(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	accessToken, err := helper.GenerateJwt(user, permissions, nil)
	helper.PanicIfErr(err)

	refreshToken, err := helper.GenerateRefreshJwt(user, nil)
	helper.PanicIfErr(err)

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response
}

func (service service) Token(c *fiber.Ctx, request TokenRequest) TokenResponse {
	ctx := c.Context()

	claims, err := helper.ParseRefreshJwt(request.RefreshToken)
	helper.PanicIfErr(err)
	if claims.ClientId != nil {
		panic(ErrInvalidGrant)
	}

	user, err := service.repository.GetUserByEmail(ctx, claims.User.Email)
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	permissions, err := service.repository.GetPermissionsByRoleId(ctx, user.RoleId)
	helper.PanicIfErr(err)

	accessToken, err := helper.GenerateJwt(user, permissions, nil)
	helper.PanicIfErr(err)

	response := TokenResponse{
		AccessToken: accessToken,
	}

	return response
}

func (service service) ForgotPassword(c *fiber.Ctx, request ForgotPasswordRequest) {
	ctx := c.Context()

	token := uuid.NewString()
	err := service.repository.SetForgotPasswordToken(ctx, request.Email, token)
	helper.PanicIfErr(err)

	emailTo := []string{request.Email}
	emailCc := []string{}
	emailSubject := "Lutfi's Smarthome Forgot Password"
	emailMessage := fmt.Sprintf("Here is your token: %s\n", token)

	err = helper.SendMail(emailTo, emailCc, emailSubject, emailMessage)
	helper.PanicIfErr(err)
}

func (service service) ResetPassword(c *fiber.Ctx, request ResetPasswordRequest) {
	ctx := c.Context()

	_, err := service.repository.GetForgotPasswordToken(ctx, request.Email, request.Token)
	helper.PanicIfErr(err)

	err = service.repository.DeleteForgotPasswordToken(ctx, request.Email, request.Token)
	helper.PanicIfErr(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfErr(err)

	err = service.repository.UpdatePassword(ctx, request.Email, string(hashedPassword))
	helper.PanicIfErr(err)
}
