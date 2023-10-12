package domain

import (
	"time"
)

type DeviceState struct {
	Id        string          `bson:"_id"`
	DeviceId  string          `bson:"device_id"`
	State     DeviceStateData `bson:"state"`
	CreatedAt time.Time       `bson:"created_at"`
	UpdatedAt time.Time       `bson:"updated_at"`
}

type DeviceStateData struct {
	Online bool `bson:"online"`
	On     bool `bson:"on"`
}
