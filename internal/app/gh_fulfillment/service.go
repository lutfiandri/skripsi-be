package gh_fulfillment

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/gh_builder"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	Sync(c *fiber.Ctx, request Request) SyncResponse
	Query(c *fiber.Ctx, request Request) Response
	Execute(c *fiber.Ctx, request Request) Response
	Disconnect(c *fiber.Ctx, request Request) Response
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) Sync(c *fiber.Ctx, request Request) SyncResponse {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.User.Id)
	helper.PanicIfErr(err)

	devices, err := service.repository.GetDevices(c.Context(), userId)
	helper.PanicIfErr(err)

	ghDevices := []gh_builder.Device{}
	for _, device := range devices {
		ghDevice := gh_builder.NewLightBuilder().
			SetAttributes(device.LastState).
			SetID(device.Id.String()).
			SetName([]string{device.Name}, device.Name, []string{device.Name}).
			SetDeviceInfo("lutfi-smarthome", device.DeviceTypeId, device.HwVersion, device.SwVersion).
			Build()
		ghDevices = append(ghDevices, ghDevice)
	}

	response := SyncResponse{
		RequestId: request.RequestId,
		Payload: SyncPayloadResponse{
			AgentUserId: claims.User.Id,
			Devices:     ghDevices,
		},
	}
	return response
}

func (service service) Query(c *fiber.Ctx, request Request) Response {
	panic("unimplemented")
}

func (service service) Execute(c *fiber.Ctx, request Request) Response {
	panic("unimplemented")
}

func (service service) Disconnect(c *fiber.Ctx, request Request) Response {
	panic("unimplemented")
}
