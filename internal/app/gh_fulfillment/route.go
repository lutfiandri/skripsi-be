package gh_fulfillment

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/google-home")

	api.Post("/fulfillment", middleware.NewAuthenticator(), controller.Fulfillment)
}
