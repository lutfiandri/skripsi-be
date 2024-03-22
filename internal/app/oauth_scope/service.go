package oauth_scope

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreateOAuthScope(c *fiber.Ctx, request CreateOAuthScopeRequest) OAuthScopeResponse
	GetOAuthScopes(c *fiber.Ctx) []OAuthScopeResponse
	GetOAuthScope(c *fiber.Ctx, request GetOAuthScopeRequest) OAuthScopeResponse
	UpdateOAuthScope(c *fiber.Ctx, request UpdateOAuthScopeRequest) OAuthScopeResponse
	DeleteOAuthScope(c *fiber.Ctx, request DeleteOAuthScopeRequest)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreateOAuthScope(c *fiber.Ctx, request CreateOAuthScopeRequest) OAuthScopeResponse {
	now := time.Now()

	oauthScope := domain.OAuthScope{
		Section:     request.Section,
		Code:        request.Code,
		Description: request.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := service.repository.CreateOAuthScope(c.Context(), oauthScope)
	helper.PanicIfErr(err)

	response := NewResponse(oauthScope)

	return response
}

func (service service) GetOAuthScopes(c *fiber.Ctx) []OAuthScopeResponse {
	oauthScopes, err := service.repository.GetOAuthScopes(c.Context())
	helper.PanicIfErr(err)

	response := NewResponses(oauthScopes)
	return response
}

func (service service) GetOAuthScope(c *fiber.Ctx, request GetOAuthScopeRequest) OAuthScopeResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthScope, err := service.repository.GetOAuthScopeById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(oauthScope)

	return response
}

func (service service) UpdateOAuthScope(c *fiber.Ctx, request UpdateOAuthScopeRequest) OAuthScopeResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	prev, err := service.repository.GetOAuthScopeById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthScope := domain.OAuthScope{
		Id:          id,
		Code:        request.Code,
		Section:     request.Section,
		Description: request.Description,
		UpdatedAt:   time.Now(),
		CreatedAt:   prev.CreatedAt,
	}

	err = service.repository.UpdateOAuthScope(c.Context(), oauthScope)
	helper.PanicIfErr(err)

	response := NewResponse(oauthScope)

	return response
}

func (service service) DeleteOAuthScope(c *fiber.Ctx, request DeleteOAuthScopeRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetOAuthScopeById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeleteOAuthScope(c.Context(), id)
	helper.PanicIfErr(err)
}
