package rest

import (
	"time"
)

type DeviceType struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ResetGuide  string    `json:"reset_guide"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetDeviceTypeRequest struct {
	Id string `params:"id"`
}

type CreateDeviceTypeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ResetGuide  string `json:"reset_guide" validate:"required"`
}

type UpdateDeviceTypeRequest struct {
	Id          string `params:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ResetGuide  string `json:"reset_guide" validate:""`
}

type DeleteDeviceTypeRequest struct {
	Id string `params:"id"`
}
