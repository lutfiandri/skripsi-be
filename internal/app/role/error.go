package role

import "github.com/gofiber/fiber/v2"

var (
	ErrNotFound           = fiber.NewError(fiber.StatusNotFound, "Role not found")
	ErrPermissionNotFound = fiber.NewError(fiber.StatusNotFound, "One or more permissions not found")
)
