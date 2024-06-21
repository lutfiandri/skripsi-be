package auth

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrDuplicateEmail     = fiber.NewError(fiber.StatusConflict, "email already registered")
	ErrInvalidCredentials = fiber.NewError(fiber.StatusForbidden, "invalid credentials")
	ErrInvalidGrant       = fiber.NewError(fiber.StatusForbidden, "invalid grant")
)
