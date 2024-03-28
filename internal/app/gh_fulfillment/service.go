package gh_fulfillment

import (
	"context"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/gh_builder"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	Sync(c *fiber.Ctx, request Request) SyncResponse
	Query(c *fiber.Ctx, request Request) QueryResponse
	Execute(c *fiber.Ctx, request Request) Response
	Disconnect(c *fiber.Ctx, request Request) Response
}

type service struct {
	repository    Repository
	deviceTypeMap map[string]domain.DeviceType
}

func NewService(repository Repository) Service {
	deviceTypes, err := repository.GetDeviceTypes(context.TODO())
	helper.PanicIfErr(err)

	deviceTypeMap := map[string]domain.DeviceType{}
	for _, deviceType := range deviceTypes {
		deviceTypeMap[deviceType.Id.String()] = deviceType
	}

	return &service{
		repository:    repository,
		deviceTypeMap: deviceTypeMap,
	}
}

func (service service) Sync(c *fiber.Ctx, request Request) SyncResponse {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.User.Id)
	helper.PanicIfErr(err)

	devices, err := service.repository.GetDevices(c.Context(), userId, nil)
	helper.PanicIfErr(err)

	ghDevices := []gh_builder.Device{}
	for _, device := range devices {
		deviceType := service.deviceTypeMap[device.DeviceTypeId]

		ghDevice := gh_builder.NewBaseDeviceBuilder().
			SetID(device.Id.String()).
			SetType(deviceType.GoogleHome.Type).
			AddTraits(deviceType.GoogleHome.Traits...).
			SetWillReportState(deviceType.GoogleHome.WillReportState).
			SetAttributes(device.LastState).
			SetRoomHint(device.Room).
			SetName([]string{device.Name}, device.Name, []string{device.Name}).
			SetDeviceInfo("lutfi-smart-home", device.DeviceTypeId, device.HwVersion, device.SwVersion).
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

func (service service) Query(c *fiber.Ctx, request Request) QueryResponse {
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.User.Id)
	helper.PanicIfErr(err)

	deviceIds := []uuid.UUID{}
	for _, device := range request.Inputs[0].Payload.Devices {
		id, err := uuid.Parse(device.Id)
		if err != nil {
			continue
		}

		deviceIds = append(deviceIds, id)
	}

	devices, err := service.repository.GetDevices(c.Context(), userId, &deviceIds)
	helper.PanicIfErr(err)

	ghDevices := map[string]any{}
	// "123": {
	// 	"on": true,
	// 	"online": true
	// }

	for _, device := range devices {
		ghDevices[device.Id.String()] = device.LastState
	}

	response := QueryResponse{
		RequestId: request.RequestId,
		Payload: QueryPayloadResponse{
			Devices: ghDevices,
		},
	}

	return response
}

func (service service) Execute(c *fiber.Ctx, request Request) Response {
	panic("unimplemented")
}

func (service service) Disconnect(c *fiber.Ctx, request Request) Response {
	panic("unimplemented")
}
