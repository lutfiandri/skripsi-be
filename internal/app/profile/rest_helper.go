package profile

import (
	"skripsi-be/internal/domain"
)

func NewResponse(user domain.User) ProfileResponse {
	result := ProfileResponse{
		Id:        user.Id.String(),
		Email:     user.Email,
		Name:      user.Name,
		RoleId:    user.RoleId.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return result
}

func NewResponses(users []domain.User) []ProfileResponse {
	var results []ProfileResponse
	for _, user := range users {
		results = append(results, NewResponse(user))
	}
	return results
}
