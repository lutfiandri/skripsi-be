package domain

import "time"

type OAuthAuthCode struct {
	Id        string    `bson:"_id"`
	AuthCode  string    `bson:"auth_code"`
	UserId    string    `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}
