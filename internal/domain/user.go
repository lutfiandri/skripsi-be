package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `bson:"_id"`
	RoleId    uuid.UUID  `bson:"role_id"`
	Email     string     `bson:"email"`
	Password  string     `bson:"password"`
	Name      string     `bson:"name"`
	ClientIds uuid.UUIDs `bson:"client_ids"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
}
