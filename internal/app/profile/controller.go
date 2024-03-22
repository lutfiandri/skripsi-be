package profile

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
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

func (controller *controller) GetProfile(c *fiber.Ctx) error {
	result := controller.service.GetProfile(c)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *controller) UpdateProfile(c *fiber.Ctx) error {
	var request UpdateProfileRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[UpdateProfileRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.UpdateProfile(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
