package oauth_scope

import (
	"time"

	"skripsi-be/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreateOAuthScope(c *fiber.Ctx, request CreateOAuthScopeRequest) (OAuthScopeResponse, error)
	GetOAuthScopes(c *fiber.Ctx) ([]OAuthScopeResponse, error)
	GetOAuthScope(c *fiber.Ctx, request GetOAuthScopeRequest) (OAuthScopeResponse, error)
	UpdateOAuthScope(c *fiber.Ctx, request UpdateOAuthScopeRequest) (OAuthScopeResponse, error)
	DeleteOAuthScope(c *fiber.Ctx, request DeleteOAuthScopeRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreateOAuthScope(c *fiber.Ctx, request CreateOAuthScopeRequest) (OAuthScopeResponse, error) {
	now := time.Now()

	oauthScope := domain.OAuthScope{
		Section:     request.Section,
		Code:        request.Code,
		Description: request.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := service.repository.CreateOAuthScope(c.Context(), oauthScope)
	if err != nil {
		return OAuthScopeResponse{}, err
	}

	response := NewResponse(oauthScope)

	return response, nil
}

func (service service) GetOAuthScopes(c *fiber.Ctx) ([]OAuthScopeResponse, error) {
	oauthScopes, err := service.repository.GetOAuthScopes(c.Context())
	if err != nil {
		return []OAuthScopeResponse{}, err
	}

	response := NewResponses(oauthScopes)
	return response, nil
}

func (service service) GetOAuthScope(c *fiber.Ctx, request GetOAuthScopeRequest) (OAuthScopeResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return OAuthScopeResponse{}, ErrNotFound
	}

	oauthScope, err := service.repository.GetOAuthScopeById(c.Context(), id)
	if err != nil {
		return OAuthScopeResponse{}, ErrNotFound
	}

	response := NewResponse(oauthScope)

	return response, nil
}

func (service service) UpdateOAuthScope(c *fiber.Ctx, request UpdateOAuthScopeRequest) (OAuthScopeResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return OAuthScopeResponse{}, ErrNotFound
	}

	prev, err := service.repository.GetOAuthScopeById(c.Context(), id)
	if err != nil {
		return OAuthScopeResponse{}, ErrNotFound
	}

	oauthScope := domain.OAuthScope{
		Id:          id,
		Code:        request.Code,
		Section:     request.Section,
		Description: request.Description,
		UpdatedAt:   time.Now(),
		CreatedAt:   prev.CreatedAt,
	}

	if err := service.repository.UpdateOAuthScope(c.Context(), oauthScope); err != nil {
		return OAuthScopeResponse{}, err
	}
	response := NewResponse(oauthScope)

	return response, nil
}

func (service service) DeleteOAuthScope(c *fiber.Ctx, request DeleteOAuthScopeRequest) error {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return ErrNotFound
	}

	if _, err := service.repository.GetOAuthScopeById(c.Context(), id); err != nil {
		return ErrNotFound
	}

	err = service.repository.DeleteOAuthScope(c.Context(), id)
	return err
}
