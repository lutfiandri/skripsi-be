package factory

import (
	"skripsi-be/internal/model/db"
	"skripsi-be/internal/model/rest"
)

// DB to Rest

func DeviceTypeDbToRest(deviceType db.DeviceType) rest.DeviceType {
	result := rest.DeviceType{
		Id:          deviceType.Id,
		Name:        deviceType.Name,
		Description: deviceType.Description,
		ResetGuide:  deviceType.ResetGuide,
		CreatedAt:   deviceType.CreatedAt,
		UpdatedAt:   deviceType.UpdatedAt,
	}

	return result
}

func DeviceTypeDbToRestMany(deviceTypes []db.DeviceType) []rest.DeviceType {
	var results []rest.DeviceType
	for _, deviceType := range deviceTypes {
		results = append(results, DeviceTypeDbToRest(deviceType))
	}
	return results
}

// Rest to DB

func CreateDeviceTypeRestToDb(deviceType rest.CreateDeviceTypeRequest) db.DeviceType {
	result := db.DeviceType{
		Name:        deviceType.Name,
		Description: deviceType.Description,
		ResetGuide:  deviceType.ResetGuide,
	}
	return result
}

func UpdateDeviceTypeRestToDb(deviceType rest.UpdateDeviceTypeRequest) db.DeviceType {
	result := db.DeviceType{
		Id:          deviceType.Id,
		Name:        deviceType.Name,
		Description: deviceType.Description,
		ResetGuide:  deviceType.ResetGuide,
	}
	return result
}
