package googlehomefactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/googlehome"
)

func NewLight(device domain.Device) googlehome.Device {
	ghDevice := NewDevice(device)
	ghDevice.Type = "action.devices.types.LIGHT"
	ghDevice.Traits = []string{
		"action.devices.traits.OnOff",
		"action.devices.traits.Brightness",
		"action.devices.traits.ColorSetting",
	}
	ghDevice.Attributes = map[string]interface{}{
		"colorModel": "rgb",
		"colorTemperatureRange": map[string]int{
			"temperatureMinK": 2000,
			"temperatureMaxK": 10000,
		},
	}

	return ghDevice
}
