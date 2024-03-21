package device_type

import (
	"time"

	"skripsi-be/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	GetDeviceTypes(c *fiber.Ctx) ([]DeviceTypeResponse, error)
	GetDeviceType(c *fiber.Ctx, request GetDeviceTypeRequest) (DeviceTypeResponse, error)
	CreateDeviceType(c *fiber.Ctx, request CreateDeviceTypeRequest) (DeviceTypeResponse, error)
	UpdateDeviceType(c *fiber.Ctx, request UpdateDeviceTypeRequest) (DeviceTypeResponse, error)
	DeleteDeviceType(c *fiber.Ctx, request DeleteDeviceTypeRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) GetDeviceTypes(c *fiber.Ctx) ([]DeviceTypeResponse, error) {
	result, err := service.repository.GetDeviceTypes(c.Context())
	if err != nil {
		return []DeviceTypeResponse{}, err
	}

	responses := NewResponses(result)

	return responses, nil
}

func (service service) GetDeviceType(c *fiber.Ctx, request GetDeviceTypeRequest) (DeviceTypeResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceTypeResponse{}, ErrNotFound
	}

	result, err := service.repository.GetDeviceTypeById(c.Context(), id)
	if err != nil {
		return DeviceTypeResponse{}, err
	}

	response := NewResponse(result)

	return response, nil
}

func (service service) CreateDeviceType(c *fiber.Ctx, request CreateDeviceTypeRequest) (DeviceTypeResponse, error) {
	now := time.Now()

	deviceType := domain.DeviceType{
		Id:               uuid.New(),
		Name:             request.Name,
		GoogleDeviceType: request.GoogleDeviceType,
		Description:      request.Description,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := service.repository.CreateDeviceType(c.Context(), deviceType); err != nil {
		return DeviceTypeResponse{}, err
	}

	response := NewResponse(deviceType)

	return response, nil
}

func (service service) UpdateDeviceType(c *fiber.Ctx, request UpdateDeviceTypeRequest) (DeviceTypeResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceTypeResponse{}, ErrNotFound
	}

	prev, err := service.repository.GetDeviceTypeById(c.Context(), id)
	if err != nil {
		return DeviceTypeResponse{}, ErrNotFound
	}

	deviceType := domain.DeviceType{
		Id:               id,
		Name:             request.Name,
		GoogleDeviceType: request.GoogleDeviceType,
		Description:      request.Description,
		UpdatedAt:        time.Now(),

		CreatedAt: prev.CreatedAt,
	}

	if err := service.repository.UpdateDeviceType(c.Context(), deviceType); err != nil {
		return DeviceTypeResponse{}, err
	}

	response := NewResponse(deviceType)

	return response, nil
}

func (service service) DeleteDeviceType(c *fiber.Ctx, request DeleteDeviceTypeRequest) error {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return ErrNotFound
	}

	if _, err := service.repository.GetDeviceTypeById(c.Context(), id); err != nil {
		return ErrNotFound
	}

	if err := service.repository.DeleteDeviceType(c.Context(), id); err != nil {
		return err
	}

	return nil
}
