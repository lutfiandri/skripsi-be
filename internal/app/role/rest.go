package role

import "time"

type RoleResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetRoleRequest struct {
	Id string `params:"id"`
}

type CreateRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions"`
}

type UpdateRoleRequest struct {
	Id          string   `params:"id"`
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions"`
}

type DeleteRoleRequest struct {
	Id string `params:"id"`
}
