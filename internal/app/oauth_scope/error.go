package oauth_scope

import "github.com/gofiber/fiber/v2"

var ErrNotFound = fiber.NewError(fiber.StatusNotFound, "OAuth Scope not found")
