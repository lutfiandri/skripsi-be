package domain

import "time"

type OAuthClient struct {
	Id           string    `bson:"_id"`
	Secret       string    `bson:"secret"`
	Name         string    `bson:"name"`
	RedirectUris []string  `bson:"redirect_uris"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
