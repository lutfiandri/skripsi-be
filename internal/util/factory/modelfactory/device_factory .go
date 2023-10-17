package modelfactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
)

// DB to Rest

func DeviceDomainToRest(device domain.Device) rest.DeviceResponse {
	result := rest.DeviceResponse{
		Id:           device.Id,
		UserId:       device.UserId,
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

func DeviceDomainToRestMany(devices []domain.Device) []rest.DeviceResponse {
	var results []rest.DeviceResponse
	for _, device := range devices {
		results = append(results, DeviceDomainToRest(device))
	}
	return results
}

// Rest to DB

func CreateDeviceRestToDomain(device rest.CreateDeviceRequest) domain.Device {
	result := domain.Device{
		DeviceTypeId: device.DeviceTypeId,
		HwVersion:    device.HwVersion,
		SwVersion:    device.SwVersion,
	}
	return result
}

func UpdateDeviceRestToDomain(device rest.UpdateDeviceRequest) domain.Device {
	result := domain.Device{
		Id:   device.Id,
		Name: device.Name,
		Room: device.Room,
	}
	return result
}

func UpdateDeviceVersionRestToDomain(device rest.UpdateDeviceVersionRequest) domain.Device {
	result := domain.Device{
		Id:        device.Id,
		HwVersion: device.HwVersion,
		SwVersion: device.SwVersion,
	}
	return result
}
