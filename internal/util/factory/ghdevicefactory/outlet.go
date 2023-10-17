package ghdevicefactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/googlehome"
)

func NewOutlet(device domain.Device) googlehome.Device {
	ghDevice := NewDevice(device)
	ghDevice.Type = "action.devices.types.OUTLET"
	ghDevice.Traits = []string{
		"action.devices.traits.OnOff",
	}
	ghDevice.Attributes = map[string]interface{}{}

	return ghDevice
}
