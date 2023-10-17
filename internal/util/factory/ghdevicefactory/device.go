package ghdevicefactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/googlehome"
)

func NewDevice(device domain.Device) googlehome.Device {
	return googlehome.Device{
		ID: device.Id,
		Name: googlehome.DeviceName{
			DefaultNames: []string{device.Name},
			Name:         device.Name,
			Nicknames:    []string{device.Name},
		},
		WillReportState: true,
		RoomHint:        device.Room,
		DeviceInfo: googlehome.DeviceInfo{
			Manufacturer: "home.ltf",
			Model:        device.DeviceTypeId,
			HwVersion:    device.HwVersion,
			SwVersion:    device.SwVersion,
		},
		OtherDeviceIDs: []googlehome.OtherDeviceID{},
		CustomData: googlehome.CustomData{
			UserId: device.UserId,
		},
	}
}
