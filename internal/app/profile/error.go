package profile

import "github.com/gofiber/fiber/v2"

var ErrNotFound = fiber.NewError(fiber.StatusNotFound, "User not found")
