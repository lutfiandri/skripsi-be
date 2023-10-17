package controller

import (
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type DeviceController interface {
	InitHttpRoute()
	GetDevices(c *fiber.Ctx) error
	GetDeviceById(c *fiber.Ctx) error
	CreateDevice(c *fiber.Ctx) error
	UpdateDevice(c *fiber.Ctx) error
	UpdateDeviceVersion(c *fiber.Ctx) error
	DeleteDevice(c *fiber.Ctx) error

	AcquireDevice(c *fiber.Ctx) error
}

type deviceController struct {
	app     *fiber.App
	service service.DeviceService
}

func NewDeviceController(app *fiber.App, service service.DeviceService) DeviceController {
	return &deviceController{
		app:     app,
		service: service,
	}
}

func (controller *deviceController) InitHttpRoute() {
	api := controller.app.Group("/devices")
	api.Post("/", controller.CreateDevice)
	api.Get("/", controller.GetDevices)
	api.Get("/:id", controller.GetDeviceById)
	api.Put("/:id", controller.UpdateDevice)
	api.Delete("/:id", controller.DeleteDevice)

	api.Post("/:id/acquire", middleware.NewAuthenticator(), controller.AcquireDevice)
}

func (controller *deviceController) GetDevices(c *fiber.Ctx) error {
	devices, err := controller.service.GetDevices(c.Context())
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(devices)

	return c.JSON(response)
}

func (controller *deviceController) GetDeviceById(c *fiber.Ctx) error {
	var request rest.GetDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.GetDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetDevice(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *deviceController) CreateDevice(c *fiber.Ctx) error {
	var request rest.CreateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.CreateDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateDevice(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller *deviceController) UpdateDevice(c *fiber.Ctx) error {
	var request rest.UpdateDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateDevice(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *deviceController) UpdateDeviceVersion(c *fiber.Ctx) error {
	var request rest.UpdateDeviceVersionRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateDeviceVersionRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateDeviceVersion(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *deviceController) DeleteDevice(c *fiber.Ctx) error {
	var request rest.DeleteDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.DeleteDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	if err := controller.service.DeleteDevice(c.Context(), request); err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller *deviceController) AcquireDevice(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.AcquireDeviceRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.AcquireDeviceRequest](c, &request, parseOption); err != nil {
		return err
	}

	if err := controller.service.AcquireDevice(c.Context(), claims, request); err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
