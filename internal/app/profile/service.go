package profile

import (
	"time"

	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetProfile(c *fiber.Ctx) ProfileResponse
	UpdateProfile(c *fiber.Ctx, request UpdateProfileRequest) ProfileResponse
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) GetProfile(c *fiber.Ctx) ProfileResponse {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)

	user, err := service.repository.GetUserByEmail(c.Context(), claims.User.Email)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(user)
	return response
}

func (service service) UpdateProfile(c *fiber.Ctx, request UpdateProfileRequest) ProfileResponse {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)

	user, err := service.repository.GetUserByEmail(c.Context(), claims.User.Email)
	helper.PanicErrIfErr(err, ErrNotFound)

	user.Name = request.Name
	user.UpdatedAt = time.Now()

	err = service.repository.UpdateUser(c.Context(), user)
	helper.PanicIfErr(err)

	response := NewResponse(user)
	return response
}
