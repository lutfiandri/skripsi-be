package oauth_client

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/oauth2/clients")

	api.Post("/", middleware.NewAuthenticator(), controller.CreateOAuthClient)
	api.Get("/", middleware.NewAuthenticator(), controller.GetOAuthClients)
	api.Get("/:id", middleware.NewAuthenticator(), controller.GetOAuthClient)
	api.Put("/:id", middleware.NewAuthenticator(), controller.UpdateOAuthClient)
	api.Delete("/:id", middleware.NewAuthenticator(), controller.DeleteOAuthClient)

	api.Get("/:id/public", controller.GetOAuthClientPublic)
}
