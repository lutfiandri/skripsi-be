package service

import (
	"context"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory/modelfactory"
	"skripsi-be/internal/util/helper"

	"github.com/google/uuid"
)

type OAuthClientService interface {
	CreateOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.CreateOAuthClientRequest) (rest.OAuthClientResponse, error)
	GetOAuthClients(ctx context.Context, claims rest.JWTClaims) ([]rest.OAuthClientResponse, error)
	GetOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.GetOAuthClientRequest) (rest.OAuthClientResponse, error)
	UpdateOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.UpdateOAuthClientRequest) (rest.OAuthClientResponse, error)
	DeleteOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.DeleteOAuthClientRequest) error

	GetOAuthClientPublic(ctx context.Context, request rest.GetOAuthClientRequest) (rest.OAuthClientPublicResponse, error)
}

type oauthClientService struct {
	oauthClientRepository repository.OAuthClientRepository
}

func NewOAuthClientService(oauthClientRepository repository.OAuthClientRepository) OAuthClientService {
	return &oauthClientService{
		oauthClientRepository: oauthClientRepository,
	}
}

func (service *oauthClientService) CreateOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.CreateOAuthClientRequest) (rest.OAuthClientResponse, error) {
	oauthClient := modelfactory.CreateOAuthClientRestToDb(request)

	clientSecret, err := helper.GenerateClientSecret(32)
	if err != nil {
		return rest.OAuthClientResponse{}, err
	}

	now := time.Now()
	oauthClient.Id = uuid.NewString()
	oauthClient.Secret = clientSecret
	oauthClient.CreatedAt = now
	oauthClient.UpdatedAt = now

	err = service.oauthClientRepository.UpsertOAuthClient(ctx, oauthClient.Id, oauthClient)
	if err != nil {
		return rest.OAuthClientResponse{}, err
	}

	response := modelfactory.OAuthClientDbToRest(oauthClient)

	return response, nil
}

func (service *oauthClientService) GetOAuthClients(ctx context.Context, claims rest.JWTClaims) ([]rest.OAuthClientResponse, error) {
	oauthClients, err := service.oauthClientRepository.GetOAuthClients(ctx)
	if err != nil {
		return []rest.OAuthClientResponse{}, err
	}

	response := modelfactory.OAuthClientDbToRestMany(oauthClients)
	return response, nil
}

func (service *oauthClientService) GetOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.GetOAuthClientRequest) (rest.OAuthClientResponse, error) {
	oauthClient, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.Id)
	if err != nil {
		return rest.OAuthClientResponse{}, err
	}

	response := modelfactory.OAuthClientDbToRest(oauthClient)

	return response, nil
}

func (service *oauthClientService) UpdateOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.UpdateOAuthClientRequest) (rest.OAuthClientResponse, error) {
	oauthClient, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.Id)
	if err != nil {
		return rest.OAuthClientResponse{}, err
	}

	oauthClient.Name = request.Name
	oauthClient.RedirectUris = request.RedirectUris
	oauthClient.UpdatedAt = time.Now()

	if err := service.oauthClientRepository.UpsertOAuthClient(ctx, request.Id, oauthClient); err != nil {
		return rest.OAuthClientResponse{}, err
	}

	response := modelfactory.OAuthClientDbToRest(oauthClient)

	return response, nil
}

func (service *oauthClientService) DeleteOAuthClient(ctx context.Context, claims rest.JWTClaims, request rest.DeleteOAuthClientRequest) error {
	_, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.Id)
	if err != nil {
		return err
	}

	err = service.oauthClientRepository.DeleteOAuthClient(ctx, request.Id)
	return err
}

func (service *oauthClientService) GetOAuthClientPublic(ctx context.Context, request rest.GetOAuthClientRequest) (rest.OAuthClientPublicResponse, error) {
	oauthClient, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.Id)
	if err != nil {
		return rest.OAuthClientPublicResponse{}, err
	}

	response := modelfactory.OAuthClientPublicDbToRest(oauthClient)

	return response, nil
}
