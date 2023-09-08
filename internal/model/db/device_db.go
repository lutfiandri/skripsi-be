package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	Id           primitive.ObjectID
	DeviceTypeId primitive.ObjectID
	UserId       primitive.ObjectID
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
