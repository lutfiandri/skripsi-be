package profile

import (
	"time"

	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetProfile(c *fiber.Ctx) ProfileResponse
	UpdateProfile(c *fiber.Ctx, request UpdateProfileRequest) ProfileResponse
	UpdatePassword(c *fiber.Ctx, request UpdatePasswordRequest)
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

func (service service) UpdatePassword(c *fiber.Ctx, request UpdatePasswordRequest) {
	ctx := c.Context()

	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId := uuid.MustParse(claims.User.Id)

	user, err := service.repository.GetUserByEmail(ctx, claims.User.Email)
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.CurrentPassword))
	helper.PanicErrIfErr(err, ErrInvalidCredentials)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfErr(err)

	err = service.repository.UpdatePassword(ctx, userId, string(hashedPassword))
	helper.PanicIfErr(err)
}
