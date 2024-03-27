package domain

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	Id           uuid.UUID      `bson:"_id"`
	UserId       *uuid.UUID     `bson:"user_id"`
	DeviceTypeId string         `bson:"device_type_id"`
	LastState    map[string]any `bson:"last_state"`
	Name         string         `bson:"name"`
	Room         string         `bson:"room"`
	HwVersion    string         `bson:"hw_version"`
	SwVersion    string         `bson:"sw_version"`
	CreatedAt    time.Time      `bson:"created_at"`
	UpdatedAt    time.Time      `bson:"updated_at"`
	DeletedAt    *time.Time     `bson:"deleted_at"`
}
