package role

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreateRole(c *fiber.Ctx, request CreateRoleRequest) RoleResponse
	GetRoles(c *fiber.Ctx) []RoleResponse
	GetRole(c *fiber.Ctx, request GetRoleRequest) RoleResponse
	UpdateRole(c *fiber.Ctx, request UpdateRoleRequest) RoleResponse
	DeleteRole(c *fiber.Ctx, request DeleteRoleRequest)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreateRole(c *fiber.Ctx, request CreateRoleRequest) RoleResponse {
	now := time.Now()

	permissions := []uuid.UUID{}
	for _, p := range request.PermissionIds {
		permission, err := uuid.Parse(p)
		helper.PanicErrIfErr(err, ErrPermissionNotFound)

		permissions = append(permissions, permission)
	}

	role := domain.Role{
		Id:            uuid.New(),
		Name:          request.Name,
		PermissionIds: permissions,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	err := service.repository.CreateRole(c.Context(), role)
	helper.PanicIfErr(err)

	response := NewResponse(role)

	return response
}

func (service service) GetRoles(c *fiber.Ctx) []RoleResponse {
	roles, err := service.repository.GetRoles(c.Context())
	helper.PanicIfErr(err)

	response := NewResponses(roles)

	return response
}

func (service service) GetRole(c *fiber.Ctx, request GetRoleRequest) RoleResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	role, err := service.repository.GetRoleById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(role)

	return response
}

func (service service) UpdateRole(c *fiber.Ctx, request UpdateRoleRequest) RoleResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	// parse permissions
	permissions := []uuid.UUID{}
	for _, p := range request.PermissionIds {
		permission, err := uuid.Parse(p)
		helper.PanicErrIfErr(err, ErrPermissionNotFound)

		permissions = append(permissions, permission)
	}

	role, err := service.repository.GetRoleById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	// fields to update
	role.Name = request.Name
	role.PermissionIds = permissions
	role.UpdatedAt = time.Now()

	err = service.repository.UpdateRole(c.Context(), role)
	helper.PanicIfErr(err)

	response := NewResponse(role)

	return response
}

func (service service) DeleteRole(c *fiber.Ctx, request DeleteRoleRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetRoleById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeleteRole(c.Context(), id)
	helper.PanicIfErr(err)
}
