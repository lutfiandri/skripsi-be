package oauth_client

import "github.com/gofiber/fiber/v2"

var ErrNotFound = fiber.NewError(fiber.StatusNotFound, "OAuth Client not found")
