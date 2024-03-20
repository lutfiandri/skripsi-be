package domain

import (
	"time"

	"github.com/google/uuid"
)

type DeviceType struct {
	Id               uuid.UUID  `bson:"_id"`
	Name             string     `bson:"name"`
	GoogleDeviceType string     `bson:"google_device_type"`
	Description      string     `bson:"description"`
	CreatedAt        time.Time  `bson:"created_at"`
	UpdatedAt        time.Time  `bson:"updated_at"`
	DeletedAt        *time.Time `bson:"deleted_at"`
}
