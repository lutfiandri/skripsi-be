package oauth_scope

import "time"

type OAuthScopeResponse struct {
	Id            string    `json:"id"`
	Description   string    `json:"description"` // example: Create, update, and delete device information
	PermissionIds []string  `json:"permission_ids"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OAuthScopePublicResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"` // example: Create, update, and delete device information
}

type GetOAuthScopeRequest struct {
	Id string `params:"id" validate:"required"`
}

type CreateOAuthScopeRequest struct {
	Description   string   `json:"description" validate:"required"`
	PermissionIds []string `json:"permission_ids" validate:"required"`
}

type UpdateOAuthScopeRequest struct {
	Id            string   `params:"id" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	PermissionIds []string `json:"permission_ids" validate:"required"`
}

type DeleteOAuthScopeRequest struct {
	Id string `params:"id" validate:"required"`
}
