package permission

import "skripsi-be/internal/domain"

func NewResponse(permission domain.Permission) PermissionResponse {
	result := PermissionResponse{
		Id:          permission.Id.String(),
		Code:        permission.Code,
		Group:       permission.Group,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}

	return result
}

func NewResponses(permissions []domain.Permission) []PermissionResponse {
	var results []PermissionResponse
	for _, permission := range permissions {
		results = append(results, NewResponse(permission))
	}
	return results
}
