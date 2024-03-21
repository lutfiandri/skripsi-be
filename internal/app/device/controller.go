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
	devices, err := controller.service.GetDevices(c)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(devices)

	return c.JSON(response)
}

func (controller controller) GetDeviceById(c *fiber.Ctx) error {
	var request GetDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[GetDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetDevice(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) CreateDevice(c *fiber.Ctx) error {
	var request CreateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[CreateDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateDevice(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller controller) UpdateDevice(c *fiber.Ctx) error {
	var request UpdateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[UpdateDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateDevice(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) UpdateDeviceVersion(c *fiber.Ctx) error {
	var request UpdateDeviceVersionRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[UpdateDeviceVersionRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateDeviceVersion(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) DeleteDevice(c *fiber.Ctx) error {
	var request DeleteDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[DeleteDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	if err := controller.service.DeleteDevice(c, request); err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller controller) AcquireDevice(c *fiber.Ctx) error {
	var request AcquireDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[AcquireDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.AcquireDevice(c, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
