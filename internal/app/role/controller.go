package role

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreateRole(c *fiber.Ctx) error
	GetRoles(c *fiber.Ctx) error
	GetRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
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

func (controller controller) CreateRole(c *fiber.Ctx) error {
	var request CreateRoleRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[CreateRoleRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreateRole(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetRoles(c *fiber.Ctx) error {
	result := controller.service.GetRoles(c)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) GetRole(c *fiber.Ctx) error {
	var request GetRoleRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetRoleRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetRole(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateRole(c *fiber.Ctx) error {
	var request UpdateRoleRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	err := helper.ParseAndValidateRequest[UpdateRoleRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateRole(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteRole(c *fiber.Ctx) error {
	var request DeleteRoleRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeleteRoleRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeleteRole(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
