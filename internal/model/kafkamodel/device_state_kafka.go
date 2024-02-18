package kafkamodel

import "time"

type DeviceState struct {
	Id        string      `json:"_id"`
	DeviceId  string      `json:"device_id"`
	State     interface{} `json:"state"`
	CreatedAt time.Time   `json:"created_at"`
}

type SmartPlugLitState struct {
	On           bool    `json:"on"`
	Volt         float64 `json:"volt"`
	MilliAmpere  float64 `json:"milli_ampere"`
	Watt         float64 `json:"watt"`
	KiloWattHour float64 `json:"kilo_watt_hour"`
	NumOfSensor  int     `json:"num_of_sensor"`
}
