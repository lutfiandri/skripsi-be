package controller

import (
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type DeviceTypeController interface {
	InitHttpRoute()
	GetDeviceTypes(c *fiber.Ctx) error
	GetDeviceTypeById(c *fiber.Ctx) error
	CreateDeviceType(c *fiber.Ctx) error
	UpdateDeviceType(c *fiber.Ctx) error
	DeleteDeviceType(c *fiber.Ctx) error
}

type deviceTypeController struct {
	app     *fiber.App
	service service.DeviceTypeService
}

func NewDeviceTypeController(app *fiber.App, service service.DeviceTypeService) DeviceTypeController {
	return &deviceTypeController{
		app:     app,
		service: service,
	}
}

func (controller *deviceTypeController) InitHttpRoute() {
	api := controller.app.Group("/device-types")
	api.Post("/", controller.CreateDeviceType)
	api.Get("/", controller.GetDeviceTypes)
	api.Get("/:id", controller.GetDeviceTypeById)
	api.Put("/:id", controller.UpdateDeviceType)
	api.Delete("/:id", controller.DeleteDeviceType)
}

func (controller *deviceTypeController) GetDeviceTypes(c *fiber.Ctx) error {
	deviceTypes, err := controller.service.GetDeviceTypes(c.Context())
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(deviceTypes)

	return c.JSON(response)
}

func (controller *deviceTypeController) GetDeviceTypeById(c *fiber.Ctx) error {
	var request rest.GetDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.GetDeviceTypeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetDeviceType(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *deviceTypeController) CreateDeviceType(c *fiber.Ctx) error {
	var request rest.CreateDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.CreateDeviceTypeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateDeviceType(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (controller *deviceTypeController) UpdateDeviceType(c *fiber.Ctx) error {
	var request rest.UpdateDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true, ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateDeviceTypeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateDeviceType(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *deviceTypeController) DeleteDeviceType(c *fiber.Ctx) error {
	var request rest.DeleteDeviceTypeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.DeleteDeviceTypeRequest](c, &request, parseOption); err != nil {
		return err
	}

	if err := controller.service.DeleteDeviceType(c.Context(), request); err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
