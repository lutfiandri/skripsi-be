package domain

import (
	"time"
)

type DeviceState struct {
	Id        string      `bson:"_id"`
	DeviceId  string      `bson:"device_id"`
	State     interface{} `bson:"state"`
	CreatedAt time.Time   `bson:"created_at"`
}

type SmartPlugLitState struct {
	On           bool    `bson:"on"`
	Volt         float64 `bson:"volt"`
	MilliAmpere  float64 `bson:"milli_ampere"`
	Watt         float64 `bson:"watt"`
	KiloWattHour float64 `bson:"kilo_watt_hour"`
	NumOfSensor  int     `bson:"num_of_sensor"`
}
