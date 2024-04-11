package oauth

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Authorize(c *fiber.Ctx) error
	Token(c *fiber.Ctx) error
}

type controller struct {
	app     *fiber.App
	service Service
}

func NewController(app *fiber.App, service Service) Controller {
	return &controller{
		app:     app,
		service: service,
	}
}

func (controller controller) Authorize(c *fiber.Ctx) error {
	var request OAuthAuthorizeRequest
	parseOption := helper.ParseOptions{ParseQuery: true}
	err := helper.ParseAndValidateRequest[OAuthAuthorizeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.Authorize(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) Token(c *fiber.Ctx) error {
	errorResponse := OAuthTokenErrorResponse{
		Error: "invalid_grant",
	}

	var request OAuthTokenRequest
	parseOption := helper.ParseOptions{ParseQuery: true, ParseParams: true, ParseBody: true}

	err := helper.ParseAndValidateRequest[OAuthTokenRequest](c, &request, parseOption)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
	}

	response, err := controller.service.Token(c, request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
	}

	return c.JSON(response)
}
