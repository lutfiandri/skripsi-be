package domain

import "time"

type OAuthScope struct {
	Id          string    `bson:"_id"`
	Code        string    `bson:"code"`        // example: read:device_state, write:device_state
	Section     string    `bson:"section"`     // example: device_state
	Description string    `bson:"description"` // example: Create, update, and delete device information
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
