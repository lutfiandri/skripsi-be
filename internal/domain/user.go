package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `bson:"_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
