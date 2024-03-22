package oauth

import "github.com/gofiber/fiber/v2"

var (
	ErrClientNotFound           = fiber.NewError(fiber.StatusNotFound, "Client not found")
	ErrUserNotFound             = fiber.NewError(fiber.StatusNotFound, "User not found")
	ErrAuthCodeNotFound         = fiber.NewError(fiber.StatusNotFound, "Auth code not found")
	ErrInvalidClientCredentials = fiber.NewError(fiber.StatusBadRequest, "Invalid client credentials")
	ErrResponseType             = fiber.NewError(fiber.StatusBadRequest, "'response_type' must be 'code'")
	ErrGrantType                = fiber.NewError(fiber.StatusBadRequest, "'grant_type' must be 'authorization_code' or 'refresh_token'")
	ErrWrongRedirectUri         = fiber.NewError(fiber.StatusBadRequest, "Wrong redirect_uri")
	ErrInvalidGrant             = fiber.NewError(fiber.StatusBadRequest, "invalid_grant")
)
