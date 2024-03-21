package device

import "github.com/gofiber/fiber/v2"

var ErrNotFound = fiber.NewError(fiber.StatusNotFound, "Device not found")
