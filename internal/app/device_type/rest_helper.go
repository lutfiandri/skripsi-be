package device_type

import (
	"skripsi-be/internal/domain"
)

func NewResponse(deviceType domain.DeviceType) DeviceTypeResponse {
	result := DeviceTypeResponse{
		Id:               deviceType.Id.String(),
		Name:             deviceType.Name,
		GoogleDeviceType: deviceType.GoogleDeviceType,
		Description:      deviceType.Description,
		CreatedAt:        deviceType.CreatedAt,
		UpdatedAt:        deviceType.UpdatedAt,
	}

	return result
}

func NewResponses(deviceTypes []domain.DeviceType) []DeviceTypeResponse {
	var results []DeviceTypeResponse
	for _, deviceType := range deviceTypes {
		results = append(results, NewResponse(deviceType))
	}
	return results
}
