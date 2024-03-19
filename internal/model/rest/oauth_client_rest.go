package rest

import "time"

type OAuthClientResponse struct {
	Id           string    `json:"id"`
	Secret       string    `json:"secret"`
	Name         string    `json:"name"`
	RedirectUris []string  `json:"redirect_uris"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OAuthClientPublicResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetOAuthClientRequest struct {
	Id string `params:"id" validate:"required"`
}

type CreateOAuthClientRequest struct {
	Name         string   `json:"name" validate:"required"`
	RedirectUris []string `json:"redirect_uris" validate:"min=0"`
}

type UpdateOAuthClientRequest struct {
	Id           string   `params:"id" validate:"required"`
	Name         string   `json:"name" validate:"required"`
	RedirectUris []string `json:"redirect_uris" validate:"min=0"`
}

type DeleteOAuthClientRequest struct {
	Id string `params:"id" validate:"required"`
}
