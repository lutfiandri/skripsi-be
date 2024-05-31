package auth

import (
	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
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

func (controller controller) Register(c *fiber.Ctx) error {
	var request RegisterRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[RegisterRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.Register(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) Login(c *fiber.Ctx) error {
	var request LoginRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[LoginRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	result := controller.service.Login(c, request)
	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller controller) ForgotPassword(c *fiber.Ctx) error {
	var request ForgotPasswordRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[ForgotPasswordRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.ForgotPassword(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}

func (controller controller) ResetPassword(c *fiber.Ctx) error {
	var request ResetPasswordRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[ResetPasswordRequest](c, &request, parseOption)
	helper.PanicIfErr(err)

	controller.service.ResetPassword(c, request)
	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
