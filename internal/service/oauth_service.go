package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"slices"
	"time"

	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OAuthService interface {
	Authorize(ctx context.Context, claims rest.JWTClaims, request rest.OAuthAuthorizeRequest) (rest.OAuthAuthorizeResponse, error)
	Token(ctx context.Context, request rest.OAuthTokenRequest) (rest.OAuthTokenResponse, error)
}

type oauthService struct {
	oauthClientRepository   repository.OAuthClientRepository
	oauthAuthCodeRepository repository.OAuthAuthCodeRepository
	userRepository          repository.UserRepository
}

func NewOAuthService(oauthClientRepository repository.OAuthClientRepository, oauthAuthCodeRepository repository.OAuthAuthCodeRepository, userRepository repository.UserRepository) OAuthService {
	return &oauthService{
		oauthClientRepository:   oauthClientRepository,
		oauthAuthCodeRepository: oauthAuthCodeRepository,
		userRepository:          userRepository,
	}
}

func (service *oauthService) Authorize(ctx context.Context, claims rest.JWTClaims, request rest.OAuthAuthorizeRequest) (rest.OAuthAuthorizeResponse, error) {
	// check response_type must be code
	if request.ResponseType != "code" {
		return rest.OAuthAuthorizeResponse{}, fiber.NewError(fiber.StatusBadRequest, "response_type must be 'code'")
	}

	// check client_id
	client, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.ClientId)
	if err != nil {
		return rest.OAuthAuthorizeResponse{}, err
	}

	// check redirect_uri
	if !slices.Contains(client.RedirectUris, request.RedirectUri) {
		return rest.OAuthAuthorizeResponse{}, fiber.NewError(fiber.StatusBadRequest, "wrong redirect_uri")
	}

	// check user exists
	if _, err := service.userRepository.GetUserById(ctx, claims.User.Id); err != nil {
		return rest.OAuthAuthorizeResponse{}, err
	}

	authCode := domain.OAuthAuthCode{
		Id:        uuid.NewString(),
		AuthCode:  uuid.NewString(),
		UserId:    claims.User.Id,
		CreatedAt: time.Now(),
	}

	if err := service.oauthAuthCodeRepository.InsertAuthCode(ctx, authCode); err != nil {
		return rest.OAuthAuthorizeResponse{}, err
	}

	// generate full redirect_uri

	response := rest.OAuthAuthorizeResponse{
		RedirectUri: service.generateRedirectUris(request.RedirectUri, authCode.AuthCode, request.State),
	}
	return response, nil
}

func (service *oauthService) Token(ctx context.Context, request rest.OAuthTokenRequest) (rest.OAuthTokenResponse, error) {
	// check client_id
	client, err := service.oauthClientRepository.GetOAuthClientById(ctx, request.ClientId)
	if err != nil {
		return rest.OAuthTokenResponse{}, err
	}

	// check redirect_uri
	if !slices.Contains(client.RedirectUris, request.RedirectUri) {
		return rest.OAuthTokenResponse{}, fiber.NewError(fiber.StatusBadRequest, "wrong redirect_uri")
	}

	// check client_secret
	if client.Secret != request.ClientSecret {
		return rest.OAuthTokenResponse{}, fiber.NewError(fiber.StatusBadRequest, "wrong client_secret")
	}

	switch request.GrantType {
	case "authorization_code":
	// authCode, err := service.oauthAuthCodeRepository.GetAuthCodeByCode(ctx, request.Code)
	// if err != nil {
	// 	return rest.OAuthTokenResponse{}, err
	// }

	// user, err := service.userRepository.GetUserById(ctx, authCode.UserId)
	// if err != nil {
	// 	return rest.OAuthTokenResponse{}, err
	// }

	// case "refresh_token":

	default:
		return rest.OAuthTokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "'grant_type' must be 'authorization_code' or 'refresh_token'")
	}

	return rest.OAuthTokenResponse{}, fiber.ErrInternalServerError
}

// helper funcs
func (service *oauthService) generateRedirectUris(baseUrl, code, state string) string {
	queryParams := url.Values{}
	queryParams.Set("code", code)
	queryParams.Set("state", state)

	redirectUri := fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())
	log.Println("redirectUri", redirectUri)
	return redirectUri
}