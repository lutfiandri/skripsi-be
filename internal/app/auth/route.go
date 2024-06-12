package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/auth")

	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Post("/token", controller.Token)

	api.Post("/forgot-password", controller.ForgotPassword)
	api.Post("/reset-password", controller.ResetPassword)
}
