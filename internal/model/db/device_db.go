package db

import (
	"time"
)

type Device struct {
	Id           string
	DeviceTypeId string
	UserId       string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
