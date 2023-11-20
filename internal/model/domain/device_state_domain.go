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
