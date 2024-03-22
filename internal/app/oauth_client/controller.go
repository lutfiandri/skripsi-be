package oauth_client

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreateOAuthClient(c *fiber.Ctx) error
	GetOAuthClients(c *fiber.Ctx) error
	GetOAuthClient(c *fiber.Ctx) error
	UpdateOAuthClient(c *fiber.Ctx) error
	DeleteOAuthClient(c *fiber.Ctx) error

	GetOAuthClientPublic(c *fiber.Ctx) error
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

func (controller controller) CreateOAuthClient(c *fiber.Ctx) error {
	var request CreateOAuthClientRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[CreateOAuthClientRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreateOAuthClient(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthClients(c *fiber.Ctx) error {
	result := controller.service.GetOAuthClients(c)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthClient(c *fiber.Ctx) error {
	var request GetOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetOAuthClientRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetOAuthClient(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateOAuthClient(c *fiber.Ctx) error {
	var request UpdateOAuthClientRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	err := helper.ParseAndValidateRequest[UpdateOAuthClientRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateOAuthClient(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteOAuthClient(c *fiber.Ctx) error {
	var request DeleteOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeleteOAuthClientRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeleteOAuthClient(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller controller) GetOAuthClientPublic(c *fiber.Ctx) error {
	var request GetOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetOAuthClientRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetOAuthClientPublic(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
