package domain

import (
	"time"

	"github.com/google/uuid"
)

type OAuthAuthCode struct {
	Id        uuid.UUID `bson:"_id"`
	UserId    uuid.UUID `bson:"user_id"`
	AuthCode  string    `bson:"auth_code"`
	CreatedAt time.Time `bson:"created_at"`
}
