package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	Id          uuid.UUID   `bson:"_id,omitempty"`
	Name        string      `bson:"name,omitempty"`
	Permissions []uuid.UUID `bson:"permissions,omitempty"`
	CreatedAt   time.Time   `bson:"created_at,omitempty"`
	UpdatedAt   time.Time   `bson:"updated_at,omitempty"`
}
