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
	if err := helper.ParseAndValidateRequest[CreateOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateOAuthClient(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthClients(c *fiber.Ctx) error {
	result, err := controller.service.GetOAuthClients(c)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthClient(c *fiber.Ctx) error {
	var request GetOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[GetOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetOAuthClient(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateOAuthClient(c *fiber.Ctx) error {
	var request UpdateOAuthClientRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	if err := helper.ParseAndValidateRequest[UpdateOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateOAuthClient(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteOAuthClient(c *fiber.Ctx) error {
	var request DeleteOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[DeleteOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	err := controller.service.DeleteOAuthClient(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller controller) GetOAuthClientPublic(c *fiber.Ctx) error {
	var request GetOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[GetOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetOAuthClientPublic(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
