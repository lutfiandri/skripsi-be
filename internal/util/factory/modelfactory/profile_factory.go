package modelfactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
)

// db to rest

func ProfileDbToRest(user domain.User) rest.ProfileResponse {
	result := rest.ProfileResponse{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return result
}

func ProfileDbToRestMany(users []domain.User) []rest.ProfileResponse {
	var results []rest.ProfileResponse
	for _, user := range users {
		results = append(results, ProfileDbToRest(user))
	}
	return results
}

// rest to db
