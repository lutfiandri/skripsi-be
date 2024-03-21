package profile

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/profile")

	api.Get("/", middleware.NewAuthenticator(), controller.GetProfile)
	api.Put("/", middleware.NewAuthenticator(), controller.UpdateProfile)
}
