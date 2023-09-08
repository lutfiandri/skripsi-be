package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceType struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	ResetGuide  string             `bson:"reset_guide"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
