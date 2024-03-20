package device

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	GetDevices(c *fiber.Ctx) ([]DeviceResponse, error)
	GetDevice(c *fiber.Ctx, request GetDeviceRequest) (DeviceResponse, error)
	CreateDevice(c *fiber.Ctx, request CreateDeviceRequest) (DeviceResponse, error)
	UpdateDevice(c *fiber.Ctx, request UpdateDeviceRequest) (DeviceResponse, error)
	UpdateDeviceVersion(c *fiber.Ctx, request UpdateDeviceVersionRequest) (DeviceResponse, error)
	DeleteDevice(c *fiber.Ctx, request DeleteDeviceRequest) error

	AcquireDevice(c *fiber.Ctx, request AcquireDeviceRequest) (DeviceResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) GetDevices(c *fiber.Ctx) ([]DeviceResponse, error) {
	result, err := service.repository.GetDevices(c.Context())
	if err != nil {
		return []DeviceResponse{}, err
	}

	responses := NewResponses(result)

	return responses, nil
}

func (service service) GetDevice(c *fiber.Ctx, request GetDeviceRequest) (DeviceResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	result, err := service.repository.GetDeviceById(c.Context(), id)
	if err != nil {
		return DeviceResponse{}, err
	}

	response := NewResponse(result)

	return response, nil
}

func (service service) CreateDevice(c *fiber.Ctx, request CreateDeviceRequest) (DeviceResponse, error) {
	now := time.Now()

	device := domain.Device{
		Id:           uuid.New(),
		HwVersion:    request.HwVersion,
		SwVersion:    request.SwVersion,
		DeviceTypeId: request.DeviceTypeId,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := service.repository.CreateDevice(c.Context(), device); err != nil {
		return DeviceResponse{}, err
	}

	response := NewResponse(device)

	return response, nil
}

func (service service) UpdateDevice(c *fiber.Ctx, request UpdateDeviceRequest) (DeviceResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	device, err := service.repository.GetDeviceById(c.Context(), id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	device.Name = request.Name
	device.Room = request.Room
	device.UpdatedAt = time.Now()

	if err := service.repository.UpdateDevice(c.Context(), device); err != nil {
		return DeviceResponse{}, err
	}

	response := NewResponse(device)

	return response, nil
}

func (service service) UpdateDeviceVersion(c *fiber.Ctx, request UpdateDeviceVersionRequest) (DeviceResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	device, err := service.repository.GetDeviceById(c.Context(), id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	device.HwVersion = request.HwVersion
	device.SwVersion = request.SwVersion
	device.UpdatedAt = time.Now()

	if err := service.repository.UpdateDevice(c.Context(), device); err != nil {
		return DeviceResponse{}, err
	}

	response := NewResponse(device)

	return response, nil
}

func (service service) AcquireDevice(c *fiber.Ctx, request AcquireDeviceRequest) (DeviceResponse, error) {
	// get device
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeviceResponse{}, ErrNotFound
	}

	device, err := service.repository.GetDeviceById(c.Context(), id)
	if err != nil {
		return DeviceResponse{}, err
	}

	// get user
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.Id)
	if err != nil {
		return DeviceResponse{}, err
	}

	// update device
	device.UserId = &userId
	device.UpdatedAt = time.Now()

	if err := service.repository.UpdateDevice(c.Context(), device); err != nil {
		return DeviceResponse{}, err
	}

	response := NewResponse(device)

	return response, nil
}

func (service service) DeleteDevice(c *fiber.Ctx, request DeleteDeviceRequest) error {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return ErrNotFound
	}
	if _, err := service.repository.GetDeviceById(c.Context(), id); err != nil {
		return err
	}

	if err := service.repository.DeleteDevice(c.Context(), id); err != nil {
		return err
	}

	return nil
}
