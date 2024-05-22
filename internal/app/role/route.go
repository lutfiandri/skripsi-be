package role

import (
	"skripsi-be/internal/constant"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/roles")

	api.Post("/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionCreateRole),
		controller.CreateRole,
	)

	api.Get("/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadRole),
		controller.GetRoles,
	)

	api.Get("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionReadRole),
		controller.GetRole,
	)

	api.Put("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionUpdateRole),
		controller.UpdateRole,
	)

	api.Delete("/:id",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionDeleteRole),
		controller.DeleteRole,
	)
}
