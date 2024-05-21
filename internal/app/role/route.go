package role

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/roles")

	api.Post("/", middleware.NewAuthenticator(), controller.CreateRole)
	api.Get("/", middleware.NewAuthenticator(), controller.GetRoles)
	api.Get("/:id", middleware.NewAuthenticator(), controller.GetRole)
	api.Put("/:id", middleware.NewAuthenticator(), controller.UpdateRole)
	api.Delete("/:id", middleware.NewAuthenticator(), controller.DeleteRole)
}
