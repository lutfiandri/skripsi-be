package rest

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceType struct {
	Id          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	ResetGuide  string             `json:"reset_guide"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetDeviceTypeRequest struct {
	Id primitive.ObjectID `params:"id"`
}

type CreateDeviceTypeRequest struct{}

type UpdateDeviceTypeRequest struct {
	Id primitive.ObjectID `params:"id"`
}

type DeleteDeviceTypeRequest struct {
	Id primitive.ObjectID `params:"id"`
}
