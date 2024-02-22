package modelfactory

import (
	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
)

// DB to Rest

func OAuthClientDbToRest(oauthClient domain.OAuthClient) rest.OAuthClientResponse {
	result := rest.OAuthClientResponse{
		Id:           oauthClient.Id,
		Secret:       oauthClient.Secret,
		Name:         oauthClient.Name,
		RedirectUris: oauthClient.RedirectUris,
		CreatedAt:    oauthClient.CreatedAt,
		UpdatedAt:    oauthClient.UpdatedAt,
	}

	return result
}

func OAuthClientDbToRestMany(oauthClients []domain.OAuthClient) []rest.OAuthClientResponse {
	var results []rest.OAuthClientResponse
	for _, oauthClient := range oauthClients {
		results = append(results, OAuthClientDbToRest(oauthClient))
	}
	return results
}

// Rest to DB

func CreateOAuthClientRestToDb(oauthClient rest.CreateOAuthClientRequest) domain.OAuthClient {
	result := domain.OAuthClient{
		Name:         oauthClient.Name,
		RedirectUris: oauthClient.RedirectUris,
	}
	return result
}

func UpdateOAuthClientRestToDb(oauthClient rest.UpdateOAuthClientRequest) domain.OAuthClient {
	result := domain.OAuthClient{
		Id:           oauthClient.Id,
		Name:         oauthClient.Name,
		RedirectUris: oauthClient.RedirectUris,
	}
	return result
}
