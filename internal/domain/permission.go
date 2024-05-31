package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	Id          uuid.UUID `bson:"_id,omitempty"`
	Code        string    `bson:"code,omitempty"`        // example: read:device_state, write:device_state
	Group       string    `bson:"group,omitempty"`       // example: device_state
	Description string    `bson:"description,omitempty"` // example: Create, update, and delete device information
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
