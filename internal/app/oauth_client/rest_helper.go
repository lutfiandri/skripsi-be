package oauth_client

import (
	"skripsi-be/internal/domain"
)

func NewResponse(oauthClient domain.OAuthClient) OAuthClientResponse {
	result := OAuthClientResponse{
		Id:           oauthClient.Id.String(),
		Secret:       oauthClient.Secret,
		Name:         oauthClient.Name,
		RedirectUris: oauthClient.RedirectUris,
		ScopeIds:     oauthClient.ScopeIds.Strings(),
		CreatedAt:    oauthClient.CreatedAt,
		UpdatedAt:    oauthClient.UpdatedAt,
	}

	return result
}

func NewResponses(oauthClients []domain.OAuthClient) []OAuthClientResponse {
	var results []OAuthClientResponse
	for _, oauthClient := range oauthClients {
		results = append(results, NewResponse(oauthClient))
	}
	return results
}

func NewPublicResponse(oauthClient domain.OAuthClient) OAuthClientPublicResponse {
	result := OAuthClientPublicResponse{
		Id:       oauthClient.Id.String(),
		Name:     oauthClient.Name,
		ScopeIds: oauthClient.ScopeIds.Strings(),
	}

	return result
}
