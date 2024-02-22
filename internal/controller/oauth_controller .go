package controller

import (
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type OAuthController interface {
	InitHttpRoute()
	Authorize(c *fiber.Ctx) error
	Token(c *fiber.Ctx) error
}

type oauthController struct {
	app     *fiber.App
	service service.OAuthService
}

func NewOAuthController(app *fiber.App, service service.OAuthService) OAuthController {
	return &oauthController{
		app:     app,
		service: service,
	}
}

func (controller *oauthController) InitHttpRoute() {
	api := controller.app.Group("/oauth2/auth")

	api.Post("/authorize", controller.Authorize)
	api.Post("/token", controller.Token)
}

func (controller *oauthController) Authorize(c *fiber.Ctx) error {
	var request rest.OAuthAuthorizeRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseQuery: true}
	if err := helper.ParseAndValidateRequest[rest.OAuthAuthorizeRequest](c, &request, parseOption); err != nil {
		return err
	}

	err := controller.service.Authorize(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller *oauthController) Token(c *fiber.Ctx) error {
	errorResponse := rest.OAuthTokenErrorResponse{
		Error: "invalid_grant",
	}

	var request rest.OAuthTokenRequest
	parseOption := helper.ParseOptions{ParseQuery: true}

	if err := helper.ParseAndValidateRequest[rest.OAuthTokenRequest](c, &request, parseOption); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
	}

	result, err := controller.service.Token(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
