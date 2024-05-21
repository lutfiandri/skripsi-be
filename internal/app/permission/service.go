package permission

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreatePermission(c *fiber.Ctx, request CreatePermissionRequest) PermissionResponse
	GetPermissions(c *fiber.Ctx) []PermissionResponse
	GetPermission(c *fiber.Ctx, request GetPermissionRequest) PermissionResponse
	UpdatePermission(c *fiber.Ctx, request UpdatePermissionRequest) PermissionResponse
	DeletePermission(c *fiber.Ctx, request DeletePermissionRequest)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreatePermission(c *fiber.Ctx, request CreatePermissionRequest) PermissionResponse {
	now := time.Now()

	role := domain.Permission{
		Id:          uuid.New(),
		Code:        request.Code,
		Group:       request.Group,
		Description: request.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := service.repository.CreatePermission(c.Context(), role)
	helper.PanicIfErr(err)

	response := NewResponse(role)

	return response
}

func (service service) GetPermissions(c *fiber.Ctx) []PermissionResponse {
	roles, err := service.repository.GetPermissions(c.Context())
	helper.PanicIfErr(err)

	response := NewResponses(roles)

	return response
}

func (service service) GetPermission(c *fiber.Ctx, request GetPermissionRequest) PermissionResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	role, err := service.repository.GetPermissionById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(role)

	return response
}

func (service service) UpdatePermission(c *fiber.Ctx, request UpdatePermissionRequest) PermissionResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	role, err := service.repository.GetPermissionById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	// fields to update
	role.Code = request.Code
	role.Group = request.Group
	role.Description = request.Description
	role.UpdatedAt = time.Now()

	err = service.repository.UpdatePermission(c.Context(), role)
	helper.PanicIfErr(err)

	response := NewResponse(role)

	return response
}

func (service service) DeletePermission(c *fiber.Ctx, request DeletePermissionRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetPermissionById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeletePermission(c.Context(), id)
	helper.PanicIfErr(err)
}
