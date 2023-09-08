package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceState struct {
	Id        primitive.ObjectID `bson:"_id"`
	DeviceId  primitive.ObjectID `bson:"device_id"`
	State     DeviceStateData    `bson:"state"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type DeviceStateData struct {
	Online bool `bson:"online"`
	On     bool `bson:"on"`
}
