package permission

import (
	"skripsi-be/internal/constant"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/permissions")

	api.Post(
		"/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionCreatePermission),
		controller.CreatePermission,
	)

	api.Get(
		"/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadPermission),
		controller.GetPermissions,
	)

	api.Get(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadPermission),
		controller.GetPermission,
	)

	api.Put(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdatePermission),
		controller.UpdatePermission,
	)

	api.Delete(
		"/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionDeletePermission),
		controller.DeletePermission,
	)
}
