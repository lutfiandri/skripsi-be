package domain

import (
	"time"
)

type Device struct {
	Id           string     `bson:"_id"`
	DeviceTypeId string     `bson:"device_type_id"`
	UserId       string     `bson:"user_id"`
	Name         string     `bson:"name"`
	Room         string     `bson:"room"`
	HwVersion    string     `bson:"hw_version"`
	SwVersion    string     `bson:"sw_version"`
	CreatedAt    time.Time  `bson:"created_at"`
	UpdatedAt    time.Time  `bson:"updated_at"`
	DeletedAt    *time.Time `bson:"deleted_at"`
}
