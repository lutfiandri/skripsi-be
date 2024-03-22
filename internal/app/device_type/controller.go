package device_type

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetDeviceTypes(c *fiber.Ctx) error
	GetDeviceTypeById(c *fiber.Ctx) error
	CreateDeviceType(c *fiber.Ctx) error
	UpdateDeviceType(c *fiber.Ctx) error
	DeleteDeviceType(c *fiber.Ctx) error
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

func (controller controller) GetDeviceTypes(c *fiber.Ctx) error {
	deviceTypes := controller.service.GetDeviceTypes(c)
	response := rest.NewSuccessResponse(deviceTypes)

	return c.JSON(response)
}

func (controller controller) GetDeviceTypeById(c *fiber.Ctx) error {
	var request GetDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetDeviceTypeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetDeviceType(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) CreateDeviceType(c *fiber.Ctx) error {
	var request CreateDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[CreateDeviceTypeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreateDeviceType(c, request)
	response := rest.NewSuccessResponse(result)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller controller) UpdateDeviceType(c *fiber.Ctx) error {
	var request UpdateDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[UpdateDeviceTypeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateDeviceType(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteDeviceType(c *fiber.Ctx) error {
	var request DeleteDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeleteDeviceTypeRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeleteDeviceType(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
