package device

import (
	"skripsi-be/internal/domain"
)

func NewResponse(device domain.Device) DeviceResponse {
	userId := device.UserId.String()

	result := DeviceResponse{
		Id:           device.Id.String(),
		UserId:       &userId,
		DeviceTypeId: device.DeviceTypeId,
		Name:         device.Name,
		Room:         device.Room,
		HwVersion:    device.HwVersion,
		SwVersion:    device.SwVersion,
		CreatedAt:    device.CreatedAt,
		UpdatedAt:    device.UpdatedAt,
	}

	return result
}

func NewResponses(devices []domain.Device) []DeviceResponse {
	var results []DeviceResponse
	for _, device := range devices {
		results = append(results, NewResponse(device))
	}
	return results
}
