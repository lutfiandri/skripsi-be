package device

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetDevices(c *fiber.Ctx) error
	GetDeviceById(c *fiber.Ctx) error
	CreateDevice(c *fiber.Ctx) error
	UpdateDevice(c *fiber.Ctx) error
	UpdateDeviceVersion(c *fiber.Ctx) error
	DeleteDevice(c *fiber.Ctx) error

	AcquireDevice(c *fiber.Ctx) error

	CommandDevice(c *fiber.Ctx) error
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

func (controller controller) GetDevices(c *fiber.Ctx) error {
	devices := controller.service.GetDevices(c)
	response := rest.NewSuccessResponse(devices)

	return c.JSON(response)
}

func (controller controller) GetDeviceById(c *fiber.Ctx) error {
	var request GetDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[GetDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.GetDevice(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) CreateDevice(c *fiber.Ctx) error {
	var request CreateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[CreateDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.CreateDevice(c, request)
	response := rest.NewSuccessResponse(result)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller controller) UpdateDevice(c *fiber.Ctx) error {
	var request UpdateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[UpdateDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateDevice(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateDeviceVersion(c *fiber.Ctx) error {
	var request UpdateDeviceVersionRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[UpdateDeviceVersionRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateDeviceVersion(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteDevice(c *fiber.Ctx) error {
	var request DeleteDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[DeleteDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.DeleteDevice(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller controller) AcquireDevice(c *fiber.Ctx) error {
	var request AcquireDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	err := helper.ParseAndValidateRequest[AcquireDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.AcquireDevice(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) CommandDevice(c *fiber.Ctx) error {
	var request CommandDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	err := helper.ParseAndValidateRequest[CommandDeviceRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.CommandDevice(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
