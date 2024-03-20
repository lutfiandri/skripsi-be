package device

import (
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/devices")

	api.Post("/", controller.CreateDevice)
	api.Get("/", controller.GetDevices)
	api.Get("/:id", controller.GetDeviceById)
	api.Put("/:id", controller.UpdateDevice)
	api.Put("/:id/version", controller.UpdateDeviceVersion)
	api.Delete("/:id", controller.DeleteDevice)

	api.Post("/:id/acquire", middleware.NewAuthenticator(), controller.AcquireDevice)
}
