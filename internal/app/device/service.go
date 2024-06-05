package device

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/helper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	GetDevices(c *fiber.Ctx) []DeviceResponse
	GetDevice(c *fiber.Ctx, request GetDeviceRequest) DeviceResponse
	CreateDevice(c *fiber.Ctx, request CreateDeviceRequest) DeviceResponse
	UpdateDevice(c *fiber.Ctx, request UpdateDeviceRequest) DeviceResponse
	UpdateDeviceVersion(c *fiber.Ctx, request UpdateDeviceVersionRequest) DeviceResponse
	DeleteDevice(c *fiber.Ctx, request DeleteDeviceRequest)

	AcquireDevice(c *fiber.Ctx, request AcquireDeviceRequest) DeviceResponse

	CommandDevice(c *fiber.Ctx, request CommandDeviceRequest)
}

type service struct {
	repository Repository
	mqttClient mqtt.Client
}

func NewService(repository Repository, mqttClient mqtt.Client) Service {
	return &service{
		repository: repository,
		mqttClient: mqttClient,
	}
}

func (service service) GetDevices(c *fiber.Ctx) []DeviceResponse {
	result, err := service.repository.GetDevices(c.Context())
	helper.PanicIfErr(err)

	responses := NewResponses(result)

	return responses
}

func (service service) GetDevice(c *fiber.Ctx, request GetDeviceRequest) DeviceResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	result, err := service.repository.GetDeviceById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(result)

	return response
}

func (service service) CreateDevice(c *fiber.Ctx, request CreateDeviceRequest) DeviceResponse {
	now := time.Now()

	device := domain.Device{
		Id:           uuid.New(),
		HwVersion:    request.HwVersion,
		SwVersion:    request.SwVersion,
		DeviceTypeId: uuid.MustParse(request.DeviceTypeId),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := service.repository.CreateDevice(c.Context(), device)
	helper.PanicIfErr(err)

	response := NewResponse(device)

	return response
}

func (service service) UpdateDevice(c *fiber.Ctx, request UpdateDeviceRequest) DeviceResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	device, err := service.repository.GetDeviceById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	device.Name = request.Name
	device.Room = request.Room
	device.UpdatedAt = time.Now()

	err = service.repository.UpdateDevice(c.Context(), device)
	helper.PanicIfErr(err)

	response := NewResponse(device)

	return response
}

func (service service) UpdateDeviceVersion(c *fiber.Ctx, request UpdateDeviceVersionRequest) DeviceResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	device, err := service.repository.GetDeviceById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	device.HwVersion = request.HwVersion
	device.SwVersion = request.SwVersion
	device.UpdatedAt = time.Now()

	err = service.repository.UpdateDevice(c.Context(), device)
	helper.PanicIfErr(err)

	response := NewResponse(device)

	return response
}

func (service service) AcquireDevice(c *fiber.Ctx, request AcquireDeviceRequest) DeviceResponse {
	// get device
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	device, err := service.repository.GetDeviceById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	// get user
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.User.Id)
	helper.PanicIfErr(err)

	// update device
	device.UserId = &userId
	device.UpdatedAt = time.Now()

	err = service.repository.UpdateDevice(c.Context(), device)
	helper.PanicIfErr(err)

	response := NewResponse(device)

	return response
}

func (service service) DeleteDevice(c *fiber.Ctx, request DeleteDeviceRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetDeviceById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeleteDevice(c.Context(), id)
	helper.PanicIfErr(err)
}

func (service service) CommandDevice(c *fiber.Ctx, request CommandDeviceRequest) {
	ctx := c.Context()

	id := uuid.MustParse(request.Id)
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId := uuid.MustParse(claims.User.Id)

	_, err := service.repository.GetDeviceByIdAndUserId(ctx, id, userId)
	helper.PanicErrIfErr(err, ErrNotFound)

	helper.CommandDeviceWithMqtt(service.mqttClient, request.Id, request.Params)
}
