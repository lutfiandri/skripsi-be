package service

import (
	"context"
	"sync"
	"time"

	"skripsi-be/internal/model/db"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, request rest.RegisterRequest) (rest.RegisterResponse, error)
	Login(ctx context.Context, request rest.LoginRequest) (rest.LoginResponse, error)
}

type authService struct {
	sync.Mutex
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) Register(ctx context.Context, request rest.RegisterRequest) (rest.RegisterResponse, error) {
	_, err := service.userRepository.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return rest.RegisterResponse{}, fiber.NewError(fiber.StatusConflict, "email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return rest.RegisterResponse{}, err
	}

	userId := uuid.NewString()
	now := time.Now()
	user := db.User{
		Id:        userId,
		Email:     request.Email,
		Password:  string(hashedPassword),
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := service.userRepository.UpsertUser(ctx, userId, user); err != nil {
		return rest.RegisterResponse{}, err
	}

	userClaimsData := rest.JWTUserClaimsData{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	accessToken, err := helper.GenerateJwt(userClaimsData)
	if err != nil {
		return rest.RegisterResponse{}, nil
	}

	response := rest.RegisterResponse{
		AccessToken: accessToken,
	}

	return response, nil
}

func (service *authService) Login(ctx context.Context, request rest.LoginRequest) (rest.LoginResponse, error) {
	user, err := service.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return rest.LoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return rest.LoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	userClaimsData := rest.JWTUserClaimsData{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	accessToken, err := helper.GenerateJwt(userClaimsData)
	if err != nil {
		return rest.LoginResponse{}, err
	}

	response := rest.LoginResponse{
		AccessToken: accessToken,
	}

	return response, nil
}
