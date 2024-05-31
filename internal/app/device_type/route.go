package device_type

import (
	"skripsi-be/internal/constant"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/device-types")

	api.Post("/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionCreateDeviceType),
		controller.CreateDeviceType,
	)

	api.Get("/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadDeviceType),
		controller.GetDeviceTypes,
	)

	api.Get("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadDeviceType),
		controller.GetDeviceTypeById,
	)

	api.Put("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdateDeviceType),
		controller.UpdateDeviceType,
	)

	api.Delete("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionDeleteDeviceType),
		controller.DeleteDeviceType,
	)
}
