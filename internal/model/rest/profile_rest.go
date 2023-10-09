package rest

import "time"

type ProfileResponse struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EditProfileRequest struct {
	Name string `json:"name" validate:"required"`
}
