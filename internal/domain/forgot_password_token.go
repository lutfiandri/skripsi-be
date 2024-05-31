package domain

import (
	"time"
)

type ForgotPasswordToken struct {
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
}
