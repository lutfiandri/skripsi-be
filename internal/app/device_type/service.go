package device_type

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	GetDeviceTypes(c *fiber.Ctx) []DeviceTypeResponse
	GetDeviceType(c *fiber.Ctx, request GetDeviceTypeRequest) DeviceTypeResponse
	CreateDeviceType(c *fiber.Ctx, request CreateDeviceTypeRequest) DeviceTypeResponse
	UpdateDeviceType(c *fiber.Ctx, request UpdateDeviceTypeRequest) DeviceTypeResponse
	DeleteDeviceType(c *fiber.Ctx, request DeleteDeviceTypeRequest)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) GetDeviceTypes(c *fiber.Ctx) []DeviceTypeResponse {
	result, err := service.repository.GetDeviceTypes(c.Context())
	helper.PanicIfErr(err)

	responses := NewResponses(result)

	return responses
}

func (service service) GetDeviceType(c *fiber.Ctx, request GetDeviceTypeRequest) DeviceTypeResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	result, err := service.repository.GetDeviceTypeById(c.Context(), id)
	helper.PanicIfErr(err)

	response := NewResponse(result)

	return response
}

func (service service) CreateDeviceType(c *fiber.Ctx, request CreateDeviceTypeRequest) DeviceTypeResponse {
	now := time.Now()

	deviceType := domain.DeviceType{
		Id:          uuid.New(),
		Name:        request.Name,
		GoogleHome:  domain.GoogleHome(request.GoogleHome),
		Description: request.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := service.repository.CreateDeviceType(c.Context(), deviceType)
	helper.PanicIfErr(err)

	response := NewResponse(deviceType)

	return response
}

func (service service) UpdateDeviceType(c *fiber.Ctx, request UpdateDeviceTypeRequest) DeviceTypeResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	prev, err := service.repository.GetDeviceTypeById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	deviceType := domain.DeviceType{
		Id:          id,
		Name:        request.Name,
		GoogleHome:  domain.GoogleHome(request.GoogleHome),
		Description: request.Description,
		UpdatedAt:   time.Now(),

		CreatedAt: prev.CreatedAt,
	}

	err = service.repository.UpdateDeviceType(c.Context(), deviceType)
	helper.PanicIfErr(err)

	response := NewResponse(deviceType)

	return response
}

func (service service) DeleteDeviceType(c *fiber.Ctx, request DeleteDeviceTypeRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetDeviceTypeById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeleteDeviceType(c.Context(), id)
	helper.PanicIfErr(err)
}
