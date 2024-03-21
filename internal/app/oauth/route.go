package oauth

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/oauth2/auth")

	api.Post("/authorize", middleware.NewAuthenticator(), controller.Authorize)
	api.Post("/token", controller.Token)
}
