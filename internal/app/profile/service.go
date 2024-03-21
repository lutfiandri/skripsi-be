package profile

import (
	"time"

	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetProfile(c *fiber.Ctx) (ProfileResponse, error)
	UpdateProfile(c *fiber.Ctx, request UpdateProfileRequest) (ProfileResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) GetProfile(c *fiber.Ctx) (ProfileResponse, error) {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)

	user, err := service.repository.GetUserByEmail(c.Context(), claims.User.Email)
	if err != nil {
		return ProfileResponse{}, ErrNotFound
	}

	response := NewResponse(user)

	return response, nil
}

func (service service) UpdateProfile(c *fiber.Ctx, request UpdateProfileRequest) (ProfileResponse, error) {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)

	user, err := service.repository.GetUserByEmail(c.Context(), claims.User.Email)
	if err != nil {
		return ProfileResponse{}, ErrNotFound
	}

	user.Name = request.Name
	user.UpdatedAt = time.Now()

	if err := service.repository.UpdateUser(c.Context(), user); err != nil {
		return ProfileResponse{}, err
	}

	response := NewResponse(user)

	return response, nil
}
