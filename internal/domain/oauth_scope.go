package domain

import (
	"time"

	"github.com/google/uuid"
)

type OAuthScope struct {
	Id            uuid.UUID  `bson:"_id,omitempty"`
	Description   string     `bson:"description,omitempty"` // example: Create, update, and delete device information
	PermissionIds uuid.UUIDs `bson:"permission_ids,omitempty"`
	CreatedAt     time.Time  `bson:"created_at,omitempty"`
	UpdatedAt     time.Time  `bson:"updated_at,omitempty"`
}
