package permission

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreatePermission(c *fiber.Ctx) error
	GetPermissions(c *fiber.Ctx) error
	GetPermission(c *fiber.Ctx) error
	UpdatePermission(c *fiber.Ctx) error
	DeletePermission(c *fiber.Ctx) error
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

func (controller controller) CreatePermission(c *fiber.Ctx) error {
	var request CreatePermissionRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[CreatePermissionRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreatePermission(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetPermissions(c *fiber.Ctx) error {
	result := controller.service.GetPermissions(c)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetPermission(c *fiber.Ctx) error {
	var request GetPermissionRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetPermissionRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetPermission(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdatePermission(c *fiber.Ctx) error {
	var request UpdatePermissionRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	err := helper.ParseAndValidateRequest[UpdatePermissionRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdatePermission(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeletePermission(c *fiber.Ctx) error {
	var request DeletePermissionRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeletePermissionRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeletePermission(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
