package oauth_scope

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreateOAuthScope(c *fiber.Ctx) error
	GetOAuthScopes(c *fiber.Ctx) error
	GetOAuthScope(c *fiber.Ctx) error
	UpdateOAuthScope(c *fiber.Ctx) error
	DeleteOAuthScope(c *fiber.Ctx) error
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

func (controller controller) CreateOAuthScope(c *fiber.Ctx) error {
	var request CreateOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[CreateOAuthScopeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreateOAuthScope(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthScopes(c *fiber.Ctx) error {
	result := controller.service.GetOAuthScopes(c)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetOAuthScope(c *fiber.Ctx) error {
	var request GetOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetOAuthScopeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetOAuthScope(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateOAuthScope(c *fiber.Ctx) error {
	var request UpdateOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	err := helper.ParseAndValidateRequest[UpdateOAuthScopeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateOAuthScope(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteOAuthScope(c *fiber.Ctx) error {
	var request DeleteOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeleteOAuthScopeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeleteOAuthScope(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
