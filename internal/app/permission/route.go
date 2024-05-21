package permission

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/permissions")

	api.Post("/", middleware.NewAuthenticator(), controller.CreatePermission)
	api.Get("/", middleware.NewAuthenticator(), controller.GetPermissions)
	api.Get("/:id", middleware.NewAuthenticator(), controller.GetPermission)
	api.Put("/:id", middleware.NewAuthenticator(), controller.UpdatePermission)
	api.Delete("/:id", middleware.NewAuthenticator(), controller.DeletePermission)
}
