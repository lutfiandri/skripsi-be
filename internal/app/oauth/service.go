package oauth

import (
	"fmt"
	"log"
	"net/url"
	"slices"
	"time"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	Authorize(c *fiber.Ctx, request OAuthAuthorizeRequest) (OAuthAuthorizeResponse, error)
	Token(c *fiber.Ctx, request OAuthTokenRequest) (OAuthTokenResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service service) Authorize(c *fiber.Ctx, request OAuthAuthorizeRequest) (OAuthAuthorizeResponse, error) {
	// check response_type must be code
	if request.ResponseType != "code" {
		return OAuthAuthorizeResponse{}, fiber.NewError(fiber.StatusBadRequest, "response_type must be 'code'")
	}

	// check client_id
	clientId, err := uuid.Parse(request.ClientId)
	if err != nil {
		return OAuthAuthorizeResponse{}, ErrClientNotFound
	}

	client, err := service.repository.GetOAuthClientById(c.Context(), clientId)
	if err != nil {
		return OAuthAuthorizeResponse{}, ErrClientNotFound
	}

	// check redirect_uri
	if !slices.Contains(client.RedirectUris, request.RedirectUri) {
		return OAuthAuthorizeResponse{}, ErrWrongRedirectUri
	}

	// check user exists
	claims := c.Locals(middleware.CtxClaims).(rest.JWTClaims)
	userId, err := uuid.Parse(claims.User.Id)
	if err != nil {
		return OAuthAuthorizeResponse{}, ErrUserNotFound
	}
	if _, err := service.repository.GetUserById(c.Context(), userId); err != nil {
		return OAuthAuthorizeResponse{}, ErrUserNotFound
	}

	authCode := domain.OAuthAuthCode{
		Id:        uuid.New(),
		UserId:    userId,
		AuthCode:  uuid.NewString(),
		CreatedAt: time.Now(),
	}

	if err := service.repository.InsertAuthCode(c.Context(), authCode); err != nil {
		return OAuthAuthorizeResponse{}, err
	}

	// generate full redirect_uri

	response := OAuthAuthorizeResponse{
		RedirectUri: service.generateRedirectUris(request.RedirectUri, authCode.AuthCode, request.State),
	}
	return response, nil
}

func (service service) Token(c *fiber.Ctx, request OAuthTokenRequest) (OAuthTokenResponse, error) {
	// check client_id
	clientId, err := uuid.Parse(request.ClientId)
	if err != nil {
		return OAuthTokenResponse{}, ErrInvalidClientCredentials
	}

	client, err := service.repository.GetOAuthClientById(c.Context(), clientId)
	if err != nil {
		return OAuthTokenResponse{}, ErrInvalidClientCredentials
	}

	// check redirect_uri
	if !slices.Contains(client.RedirectUris, request.RedirectUri) {
		return OAuthTokenResponse{}, ErrWrongRedirectUri
	}

	// check client_secret
	if client.Secret != request.ClientSecret {
		return OAuthTokenResponse{}, ErrInvalidClientCredentials
	}

	switch request.GrantType {
	case "authorization_code":

		authCode, err := service.repository.GetAuthCodeByCode(c.Context(), request.Code)
		if err != nil {
			log.Println(request.Code, err)
			return OAuthTokenResponse{}, err
		}

		user, err := service.repository.GetUserById(c.Context(), authCode.UserId)
		if err != nil {
			return OAuthTokenResponse{}, err
		}

		accessToken, err := helper.GenerateJwt(user)
		if err != nil {
			return OAuthTokenResponse{}, err
		}
		refreshToken, err := helper.GenerateRefreshJwt(user)
		if err != nil {
			return OAuthTokenResponse{}, err
		}

		response := OAuthTokenResponse{
			TokenType:    "Bearer",
			AccessToken:  accessToken,  // 1 hour
			RefreshToken: refreshToken, // doesn't expire
			ExpiresIn:    60 * 60,      // 1 hour
		}
		return response, nil

	case "refresh_token":
		refreshClaims, err := helper.ParseRefreshJwt(request.RefreshToken)
		if err != nil {
			return OAuthTokenResponse{}, err
		}

		userId, err := uuid.Parse(refreshClaims.User.Id)
		if err != nil {
			return OAuthTokenResponse{}, ErrClientNotFound
		}

		user, err := service.repository.GetUserById(c.Context(), userId)
		if err != nil {
			return OAuthTokenResponse{}, err
		}

		accessToken, err := helper.GenerateJwt(user)
		if err != nil {
			return OAuthTokenResponse{}, err
		}

		response := OAuthTokenResponse{
			TokenType:   "Bearer",
			AccessToken: accessToken, // 1 hour
			ExpiresIn:   60 * 60,     // 1 hour
		}
		return response, nil

	default:
		return OAuthTokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "'grant_type' must be 'authorization_code' or 'refresh_token'")
	}
}

// helper funcs
func (service service) generateRedirectUris(baseUrl, code, state string) string {
	queryParams := url.Values{}
	queryParams.Set("code", code)
	queryParams.Set("state", state)

	redirectUri := fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())
	return redirectUri
}
