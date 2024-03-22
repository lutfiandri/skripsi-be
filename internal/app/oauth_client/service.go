package oauth_client

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreateOAuthClient(c *fiber.Ctx, request CreateOAuthClientRequest) OAuthClientResponse
	GetOAuthClients(c *fiber.Ctx) []OAuthClientResponse
	GetOAuthClient(c *fiber.Ctx, request GetOAuthClientRequest) OAuthClientResponse
	UpdateOAuthClient(c *fiber.Ctx, request UpdateOAuthClientRequest) OAuthClientResponse
	DeleteOAuthClient(c *fiber.Ctx, request DeleteOAuthClientRequest)

	GetOAuthClientPublic(c *fiber.Ctx, request GetOAuthClientRequest) OAuthClientPublicResponse
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreateOAuthClient(c *fiber.Ctx, request CreateOAuthClientRequest) OAuthClientResponse {
	clientSecret, err := helper.GenerateClientSecret(64)
	helper.PanicIfErr(err)

	now := time.Now()
	oauthClient := domain.OAuthClient{
		Id:           uuid.New(),
		Secret:       clientSecret,
		Name:         request.Name,
		RedirectUris: request.RedirectUris,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	err = service.repository.CreateOAuthClient(c.Context(), oauthClient)
	helper.PanicIfErr(err)

	response := NewResponse(oauthClient)

	return response
}

func (service service) GetOAuthClients(c *fiber.Ctx) []OAuthClientResponse {
	oauthClients, err := service.repository.GetOAuthClients(c.Context())
	helper.PanicIfErr(err)

	response := NewResponses(oauthClients)

	return response
}

func (service service) GetOAuthClient(c *fiber.Ctx, request GetOAuthClientRequest) OAuthClientResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewResponse(oauthClient)

	return response
}

func (service service) UpdateOAuthClient(c *fiber.Ctx, request UpdateOAuthClientRequest) OAuthClientResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthClient.Name = request.Name
	oauthClient.RedirectUris = request.RedirectUris
	oauthClient.UpdatedAt = time.Now()

	err = service.repository.UpdateOAuthClient(c.Context(), oauthClient)
	helper.PanicIfErr(err)

	response := NewResponse(oauthClient)

	return response
}

func (service service) DeleteOAuthClient(c *fiber.Ctx, request DeleteOAuthClientRequest) {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	_, err = service.repository.GetOAuthClientById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	err = service.repository.DeleteOAuthClient(c.Context(), id)
	helper.PanicIfErr(err)
}

func (service service) GetOAuthClientPublic(c *fiber.Ctx, request GetOAuthClientRequest) OAuthClientPublicResponse {
	id, err := uuid.Parse(request.Id)
	helper.PanicErrIfErr(err, ErrNotFound)

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	helper.PanicErrIfErr(err, ErrNotFound)

	response := NewPublicResponse(oauthClient)

	return response
}
