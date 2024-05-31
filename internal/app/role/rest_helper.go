package role

import "skripsi-be/internal/domain"

func NewResponse(role domain.Role) RoleResponse {
	permissions := []string{}
	for _, p := range role.PermissionIds {
		permissions = append(permissions, p.String())
	}

	result := RoleResponse{
		Id:            role.Id.String(),
		Name:          role.Name,
		PermissionIds: permissions,
		CreatedAt:     role.CreatedAt,
		UpdatedAt:     role.UpdatedAt,
	}

	return result
}

func NewResponses(roles []domain.Role) []RoleResponse {
	var results []RoleResponse
	for _, role := range roles {
		results = append(results, NewResponse(role))
	}
	return results
}
