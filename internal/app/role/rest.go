package role

import "time"

type RoleResponse struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	PermissionIds []string  `json:"permission_ids"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetRoleRequest struct {
	Id string `params:"id"`
}

type CreateRoleRequest struct {
	Name          string   `json:"name" validate:"required"`
	PermissionIds []string `json:"permission_ids"`
}

type UpdateRoleRequest struct {
	Id            string   `params:"id"`
	Name          string   `json:"name" validate:"required"`
	PermissionIds []string `json:"permission_ids"`
}

type DeleteRoleRequest struct {
	Id string `params:"id"`
}
