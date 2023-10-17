package modelfactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
)

// DB to Rest

func DeviceTypeDbToRest(deviceType domain.DeviceType) rest.DeviceTypeResponse {
	result := rest.DeviceTypeResponse{
		Id:               deviceType.Id,
		Name:             deviceType.Name,
		GoogleDeviceType: deviceType.GoogleDeviceType,
		Description:      deviceType.Description,
		CreatedAt:        deviceType.CreatedAt,
		UpdatedAt:        deviceType.UpdatedAt,
	}

	return result
}

func DeviceTypeDbToRestMany(deviceTypes []domain.DeviceType) []rest.DeviceTypeResponse {
	var results []rest.DeviceTypeResponse
	for _, deviceType := range deviceTypes {
		results = append(results, DeviceTypeDbToRest(deviceType))
	}
	return results
}

// Rest to DB

func CreateDeviceTypeRestToDb(deviceType rest.CreateDeviceTypeRequest) domain.DeviceType {
	result := domain.DeviceType{
		Name:             deviceType.Name,
		GoogleDeviceType: deviceType.GoogleDeviceType,
		Description:      deviceType.Description,
	}
	return result
}

func UpdateDeviceTypeRestToDb(deviceType rest.UpdateDeviceTypeRequest) domain.DeviceType {
	result := domain.DeviceType{
		Id:               deviceType.Id,
		Name:             deviceType.Name,
		GoogleDeviceType: deviceType.GoogleDeviceType,
		Description:      deviceType.Description,
	}
	return result
}
