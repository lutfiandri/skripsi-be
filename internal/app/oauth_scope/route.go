package oauth_scope

import (
	"skripsi-be/internal/constant"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller Controller) {
	api := app.Group("/oauth2/scopes")

	api.Post(
		"/",
		middleware.NewAuthenticator(),
		middleware.NewAuthorizer(constant.PermissionCreateOAuthScope),
		controller.CreateOAuthScope,
	)

	api.Get(
		"/",
		middleware.NewAuthenticator(),
		controller.GetOAuthScopes,
	)

	api.Get(
		"/:id",
		middleware.NewAuthenticator(),
		controller.GetOAuthScope,
	)

	api.Get("/:id/public", controller.GetOAuthScopePublic)

	api.Put(
		"/:id",
		middleware.NewAuthenticator(),
		controller.UpdateOAuthScope,
	)

	api.Delete(
		"/:id",
		middleware.NewAuthenticator(),
		controller.DeleteOAuthScope,
	)
}
