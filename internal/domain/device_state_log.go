package domain

import (
	"time"

	"github.com/google/uuid"
)

type DeviceStateLog[T any] struct {
	Id        uuid.UUID `bson:"_id"`
	DeviceId  uuid.UUID `bson:"device_id"`
	State     T         `bson:"state"`
	CreatedAt time.Time `bson:"created_at"`
}
