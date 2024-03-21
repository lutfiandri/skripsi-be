package device_type

import "github.com/gofiber/fiber/v2"

var ErrNotFound = fiber.NewError(fiber.StatusNotFound, "Device type not found")
