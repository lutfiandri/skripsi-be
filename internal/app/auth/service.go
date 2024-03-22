package auth

import (
	"time"

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

	now := time.Now()
	user := domain.User{
		Id:        uuid.New(),
		Email:     request.Email,
		Password:  string(hashedPassword),
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = service.repository.CreateUser(ctx, user)
	helper.PanicIfErr(err)

	accessToken, err := helper.GenerateJwt(user)
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
	ctx := c.Context()

	user, err := service.repository.GetUserByEmail(ctx, request.Email)
	helper.PanicErrIfNotErr(err, ErrInvalidCredentials)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfErr(err)

	accessToken, err := helper.GenerateJwt(user)
	helper.PanicIfErr(err)

	refreshToken, err := helper.GenerateRefreshJwt(user)
	helper.PanicIfErr(err)

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response
}
