package oauth_scope

import (
	"skripsi-be/internal/domain"
)

func NewResponse(oauthScope domain.OAuthScope) OAuthScopeResponse {
	result := OAuthScopeResponse{
		Id:          oauthScope.Id.String(),
		Section:     oauthScope.Section,
		Code:        oauthScope.Code,
		Description: oauthScope.Description,
		CreatedAt:   oauthScope.CreatedAt,
		UpdatedAt:   oauthScope.UpdatedAt,
	}

	return result
}

func NewResponses(oauthScopes []domain.OAuthScope) []OAuthScopeResponse {
	var results []OAuthScopeResponse
	for _, oauthScope := range oauthScopes {
		results = append(results, NewResponse(oauthScope))
	}
	return results
}
