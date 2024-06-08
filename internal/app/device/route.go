package device

import (
	"skripsi-be/internal/constant"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/devices")

	api.Post(
		"/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionCreateDevice),
		controller.CreateDevice,
	)

	api.Get(
		"/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadDevice),
		controller.GetDevices,
	)

	api.Get(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadDevice),
		controller.GetDeviceById,
	)

	api.Put(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdateDevice),
		controller.UpdateDevice,
	)

	api.Put("/:id/version",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdateVersionDevice),
		controller.UpdateDeviceVersion,
	)

	api.Delete(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionDeleteDevice),
		controller.DeleteDevice,
	)

	api.Post(
		"/:id/acquire",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionAcquireDevice),
		controller.AcquireDevice,
	)

	api.Post(
		"/:id/unacquire",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionAcquireDevice),
		controller.UnacquireDevice,
	)

	api.Post(
		"/:id/command",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdateDevice),
		controller.CommandDevice,
	)
}
