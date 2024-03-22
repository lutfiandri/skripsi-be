package auth

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrDuplicateEmail     = fiber.NewError(fiber.StatusConflict, "email already registered")
	ErrInvalidCredentials = fiber.NewError(fiber.StatusForbidden, "invalid credentials")
)

// var (
// 	ErrDuplicateEmail     = helper.BuildPanicMessage(fiber.StatusConflict, "email already registered")
// 	ErrInvalidCredentials = helper.BuildPanicMessage(fiber.StatusForbidden, "invalid credentials")
// )
