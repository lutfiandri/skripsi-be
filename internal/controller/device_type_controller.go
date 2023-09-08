package controller

import (
	"skripsi-be/internal/service"

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
	api.Get("/", controller.GetDeviceTypes)
	api.Get("/:id", controller.GetDeviceTypeById)
}

func (controller *deviceTypeController) GetDeviceTypes(c *fiber.Ctx) error {
	deviceTypes, err := controller.service.GetDeviceTypes(c.Context())
	if err != nil {
		return err
	}

	c.JSON(deviceTypes)

	return nil
}

func (controller *deviceTypeController) GetDeviceTypeById(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (controller *deviceTypeController) CreateDeviceType(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (controller *deviceTypeController) UpdateDeviceType(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (controller *deviceTypeController) DeleteDeviceType(c *fiber.Ctx) error {
	panic("unimplemented")
}
