package permission

import "time"

type PermissionResponse struct {
	Id          string    `json:"id"`
	Code        string    `json:"code"`
	Group       string    `json:"group"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetPermissionRequest struct {
	Id string `params:"id"`
}

type CreatePermissionRequest struct {
	Code        string `json:"code" validate:"required"`
	Group       string `json:"group" validate:"required"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	Id          string `params:"id"`
	Code        string `json:"code" validate:"required"`
	Group       string `json:"group" validate:"required"`
	Description string `json:"description"`
}

type DeletePermissionRequest struct {
	Id string `params:"id"`
}
