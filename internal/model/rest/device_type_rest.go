package rest

import (
	"time"
)

type DeviceTypeResponse struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	GoogleDeviceType string    `json:"google_device_type"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type GetDeviceTypeRequest struct {
	Id string `params:"id"`
}

type CreateDeviceTypeRequest struct {
	Name             string `json:"name" validate:"required"`
	GoogleDeviceType string `json:"google_device_type" validate:"required"`
	Description      string `json:"description" validate:"required"`
}

type UpdateDeviceTypeRequest struct {
	Id               string `params:"id"`
	Name             string `json:"name" validate:"required"`
	GoogleDeviceType string `json:"google_device_type" validate:"required"`
	Description      string `json:"description" validate:"required"`
}

type DeleteDeviceTypeRequest struct {
	Id string `params:"id"`
}
