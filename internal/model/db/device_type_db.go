package db

import (
	"time"
)

type DeviceType struct {
	Id          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	ResetGuide  string    `bson:"reset_guide"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
