package rest

import "time"

type OAuthScopeResponse struct {
	Id          string    `json:"id"`
	Code        string    `json:"code"`        // example: read:device_state, write:device_state
	Section     string    `json:"section"`     // example: device_state
	Description string    `json:"description"` // example: Create, update, and delete device information
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetOAuthScopeRequest struct {
	Id string `params:"id" validate:"required"`
}

type CreateOAuthScopeRequest struct {
	Code        string `json:"code" validate:"required"`
	Section     string `json:"section" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateOAuthScopeRequest struct {
	Id          string `params:"id" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Section     string `json:"section" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteOAuthScopeRequest struct {
	Id string `params:"id" validate:"required"`
}
