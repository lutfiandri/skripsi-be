package device_type

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/device-types")

	api.Post("/", controller.CreateDeviceType)
	api.Get("/", controller.GetDeviceTypes)
	api.Get("/:id", controller.GetDeviceTypeById)
	api.Put("/:id", controller.UpdateDeviceType)
	api.Delete("/:id", controller.DeleteDeviceType)
}
