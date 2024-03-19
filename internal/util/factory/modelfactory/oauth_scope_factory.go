package modelfactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
)

// DB to Rest

func OAuthScopeDbToRest(oauthScope domain.OAuthScope) rest.OAuthScopeResponse {
	result := rest.OAuthScopeResponse{
		Id:          oauthScope.Id,
		Section:     oauthScope.Section,
		Code:        oauthScope.Code,
		Description: oauthScope.Description,
		CreatedAt:   oauthScope.CreatedAt,
		UpdatedAt:   oauthScope.UpdatedAt,
	}

	return result
}

func OAuthScopeDbToRestMany(oauthScopes []domain.OAuthScope) []rest.OAuthScopeResponse {
	var results []rest.OAuthScopeResponse
	for _, oauthScope := range oauthScopes {
		results = append(results, OAuthScopeDbToRest(oauthScope))
	}
	return results
}

// Rest to DB

func CreateOAuthScopeRestToDb(oauthScope rest.CreateOAuthScopeRequest) domain.OAuthScope {
	result := domain.OAuthScope{
		Section:     oauthScope.Section,
		Code:        oauthScope.Code,
		Description: oauthScope.Description,
	}
	return result
}

func UpdateOAuthScopeRestToDb(oauthScope rest.UpdateOAuthScopeRequest) domain.OAuthScope {
	result := domain.OAuthScope{
		Id:          oauthScope.Id,
		Section:     oauthScope.Section,
		Code:        oauthScope.Code,
		Description: oauthScope.Description,
	}
	return result
}
