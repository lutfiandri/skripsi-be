package auth

import (
	"log"
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

	accessToken, err := helper.GenerateJwt(user, permissions)
	helper.PanicIfErr(err)

	refreshToken, err := helper.GenerateRefreshJwt(user)
	helper.PanicIfErr(err)

	response := RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response
}

func (service service) Login(c *fiber.Ctx, request LoginRequest) LoginResponse {
	log.Println("login")
	ctx := c.Context()

	user, err := service.repository.GetUserByEmail(ctx, request.Email)
	// log.Println(err.Error())
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	permissions, err := service.repository.GetPermissionsByRoleId(ctx, user.RoleId)
	helper.PanicIfErr(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	accessToken, err := helper.GenerateJwt(user, permissions)
	helper.PanicIfErr(err)

	refreshToken, err := helper.GenerateRefreshJwt(user)
	helper.PanicIfErr(err)

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response
}
