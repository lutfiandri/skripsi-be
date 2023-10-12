package domain

import "time"

type User struct {
	Id        string    `bson:"_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
