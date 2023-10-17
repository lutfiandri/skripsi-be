package rest

import "time"

type DeviceResponse struct {
	Id           string    `json:"id"`
	UserId       *string   `json:"user_id"`
	DeviceTypeId string    `json:"device_type_id"`
	Name         string    `json:"name"`
	Room         string    `json:"room"`
	HwVersion    string    `json:"hw_version"`
	SwVersion    string    `json:"sw_version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateDeviceRequest struct {
	DeviceTypeId string `json:"device_type_id" validate:"required"`
	HwVersion    string `json:"hw_version" validate:"required"`
	SwVersion    string `json:"sw_version" validate:"required"`
}

type GetDevicesRequest struct{}

type GetDeviceRequest struct {
	Id string `params:"id"`
}

// for admin
type UpdateDeviceVersionRequest struct {
	Id        string `params:"id"`
	HwVersion string `json:"hw_version" validate:"required"`
	SwVersion string `json:"sw_version" validate:"required"`
}

// for user
type UpdateDeviceRequest struct {
	Id   string `params:"id"`
	Name string `json:"name" validate:"required"`
	Room string `json:"room" validate:"required"`
}

type DeleteDeviceRequest struct {
	Id string `params:"id"`
}

// user acquiring a device
type AcquireDeviceRequest struct {
	Id string `params:"id"`
}
