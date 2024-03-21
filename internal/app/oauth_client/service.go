package oauth_client

import (
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	CreateOAuthClient(c *fiber.Ctx, request CreateOAuthClientRequest) (OAuthClientResponse, error)
	GetOAuthClients(c *fiber.Ctx) ([]OAuthClientResponse, error)
	GetOAuthClient(c *fiber.Ctx, request GetOAuthClientRequest) (OAuthClientResponse, error)
	UpdateOAuthClient(c *fiber.Ctx, request UpdateOAuthClientRequest) (OAuthClientResponse, error)
	DeleteOAuthClient(c *fiber.Ctx, request DeleteOAuthClientRequest) error

	GetOAuthClientPublic(c *fiber.Ctx, request GetOAuthClientRequest) (OAuthClientPublicResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) CreateOAuthClient(c *fiber.Ctx, request CreateOAuthClientRequest) (OAuthClientResponse, error) {
	clientSecret, err := helper.GenerateClientSecret(64)
	if err != nil {
		return OAuthClientResponse{}, err
	}

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
	if err != nil {
		return OAuthClientResponse{}, err
	}

	response := NewResponse(oauthClient)

	return response, nil
}

func (service service) GetOAuthClients(c *fiber.Ctx) ([]OAuthClientResponse, error) {
	oauthClients, err := service.repository.GetOAuthClients(c.Context())
	if err != nil {
		return []OAuthClientResponse{}, err
	}

	response := NewResponses(oauthClients)

	return response, nil
}

func (service service) GetOAuthClient(c *fiber.Ctx, request GetOAuthClientRequest) (OAuthClientResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return OAuthClientResponse{}, ErrNotFound
	}

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	if err != nil {
		return OAuthClientResponse{}, ErrNotFound
	}

	response := NewResponse(oauthClient)

	return response, nil
}

func (service service) UpdateOAuthClient(c *fiber.Ctx, request UpdateOAuthClientRequest) (OAuthClientResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return OAuthClientResponse{}, ErrNotFound
	}

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	if err != nil {
		return OAuthClientResponse{}, ErrNotFound
	}

	oauthClient.Name = request.Name
	oauthClient.RedirectUris = request.RedirectUris
	oauthClient.UpdatedAt = time.Now()

	if err := service.repository.UpdateOAuthClient(c.Context(), oauthClient); err != nil {
		return OAuthClientResponse{}, err
	}

	response := NewResponse(oauthClient)

	return response, nil
}

func (service service) DeleteOAuthClient(c *fiber.Ctx, request DeleteOAuthClientRequest) error {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return ErrNotFound
	}

	if _, err := service.repository.GetOAuthClientById(c.Context(), id); err != nil {
		return ErrNotFound
	}

	err = service.repository.DeleteOAuthClient(c.Context(), id)
	return err
}

func (service service) GetOAuthClientPublic(c *fiber.Ctx, request GetOAuthClientRequest) (OAuthClientPublicResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return OAuthClientPublicResponse{}, ErrNotFound
	}

	oauthClient, err := service.repository.GetOAuthClientById(c.Context(), id)
	if err != nil {
		return OAuthClientPublicResponse{}, ErrNotFound
	}

	response := NewPublicResponse(oauthClient)

	return response, nil
}
