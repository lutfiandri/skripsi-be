package domain

import (
	"time"
)

type DeviceType struct {
	Id               string    `bson:"_id"`
	Name             string    `bson:"name"`
	GoogleDeviceType string    `bson:"google_device_type"`
	Description      string    `bson:"description"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
}
