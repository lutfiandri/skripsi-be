package controller

import (
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	InitHttpRoute()
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authController struct {
	app     *fiber.App
	service service.AuthService
}

func NewAuthController(app *fiber.App, service service.AuthService) AuthController {
	return &authController{
		app:     app,
		service: service,
	}
}

func (controller *authController) InitHttpRoute() {
	api := controller.app.Group("/auth")
	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
}

func (controller *authController) Register(c *fiber.Ctx) error {
	var request rest.RegisterRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.RegisterRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.Register(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	var request rest.LoginRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.LoginRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.Login(c.Context(), request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
