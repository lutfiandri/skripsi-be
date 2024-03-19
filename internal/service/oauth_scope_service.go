package service

import (
	"context"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory/modelfactory"

	"github.com/google/uuid"
)

type OAuthScopeService interface {
	CreateOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.CreateOAuthScopeRequest) (rest.OAuthScopeResponse, error)
	GetOAuthScopes(ctx context.Context, claims rest.JWTClaims) ([]rest.OAuthScopeResponse, error)
	GetOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.GetOAuthScopeRequest) (rest.OAuthScopeResponse, error)
	UpdateOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.UpdateOAuthScopeRequest) (rest.OAuthScopeResponse, error)
	DeleteOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.DeleteOAuthScopeRequest) error
}

type oauthScopeService struct {
	oauthScopeRepository repository.OAuthScopeRepository
}

func NewOAuthScopeService(oauthScopeRepository repository.OAuthScopeRepository) OAuthScopeService {
	return &oauthScopeService{
		oauthScopeRepository: oauthScopeRepository,
	}
}

func (service *oauthScopeService) CreateOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.CreateOAuthScopeRequest) (rest.OAuthScopeResponse, error) {
	oauthScope := modelfactory.CreateOAuthScopeRestToDb(request)

	now := time.Now()
	oauthScope.Id = uuid.NewString()
	oauthScope.CreatedAt = now
	oauthScope.UpdatedAt = now

	err := service.oauthScopeRepository.UpsertOAuthScope(ctx, oauthScope.Id, oauthScope)
	if err != nil {
		return rest.OAuthScopeResponse{}, err
	}

	response := modelfactory.OAuthScopeDbToRest(oauthScope)

	return response, nil
}

func (service *oauthScopeService) GetOAuthScopes(ctx context.Context, claims rest.JWTClaims) ([]rest.OAuthScopeResponse, error) {
	oauthScopes, err := service.oauthScopeRepository.GetOAuthScopes(ctx)
	if err != nil {
		return []rest.OAuthScopeResponse{}, err
	}

	response := modelfactory.OAuthScopeDbToRestMany(oauthScopes)
	return response, nil
}

func (service *oauthScopeService) GetOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.GetOAuthScopeRequest) (rest.OAuthScopeResponse, error) {
	oauthScope, err := service.oauthScopeRepository.GetOAuthScopeById(ctx, request.Id)
	if err != nil {
		return rest.OAuthScopeResponse{}, err
	}

	response := modelfactory.OAuthScopeDbToRest(oauthScope)

	return response, nil
}

func (service *oauthScopeService) UpdateOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.UpdateOAuthScopeRequest) (rest.OAuthScopeResponse, error) {
	oauthScope, err := service.oauthScopeRepository.GetOAuthScopeById(ctx, request.Id)
	if err != nil {
		return rest.OAuthScopeResponse{}, err
	}

	oauthScope = modelfactory.UpdateOAuthScopeRestToDb(request)
	oauthScope.UpdatedAt = time.Now()

	if err := service.oauthScopeRepository.UpsertOAuthScope(ctx, request.Id, oauthScope); err != nil {
		return rest.OAuthScopeResponse{}, err
	}

	response := modelfactory.OAuthScopeDbToRest(oauthScope)

	return response, nil
}

func (service *oauthScopeService) DeleteOAuthScope(ctx context.Context, claims rest.JWTClaims, request rest.DeleteOAuthScopeRequest) error {
	_, err := service.oauthScopeRepository.GetOAuthScopeById(ctx, request.Id)
	if err != nil {
		return err
	}

	err = service.oauthScopeRepository.DeleteOAuthScope(ctx, request.Id)
	return err
}
