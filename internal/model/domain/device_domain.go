package domain

import (
	"time"
)

type Device struct {
	Id           string
	DeviceTypeId string
	UserId       string
	Name         string
	Room         string
	HwVersion    string
	SwVersion    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
