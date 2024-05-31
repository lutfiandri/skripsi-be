package auth

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/auth")

	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)

	api.Put(
		"/password",
		middleware.NewAuthenticator(),
		controller.UpdatePassword,
	)
}
