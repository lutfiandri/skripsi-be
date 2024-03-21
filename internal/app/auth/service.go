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
	Register(c *fiber.Ctx, request RegisterRequest) (RegisterResponse, error)
	Login(c *fiber.Ctx, request LoginRequest) (LoginResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) Register(c *fiber.Ctx, request RegisterRequest) (RegisterResponse, error) {
	ctx := c.Context()

	_, err := service.repository.GetUserByEmail(c.Context(), request.Email)
	if err == nil {
		return RegisterResponse{}, ErrDuplicateEmail
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
	}

	now := time.Now()
	user := domain.User{
		Id:        uuid.New(),
		Email:     request.Email,
		Password:  string(hashedPassword),
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := service.repository.CreateUser(ctx, user); err != nil {
		return RegisterResponse{}, err
	}

	accessToken, err := helper.GenerateJwt(user)
	if err != nil {
		return RegisterResponse{}, nil
	}

	refreshToken, err := helper.GenerateRefreshJwt(user)
	if err != nil {
		return RegisterResponse{}, nil
	}

	response := RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (service service) Login(c *fiber.Ctx, request LoginRequest) (LoginResponse, error) {
	ctx := c.Context()

	user, err := service.repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	accessToken, err := helper.GenerateJwt(user)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := helper.GenerateRefreshJwt(user)
	if err != nil {
		return LoginResponse{}, err
	}

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
