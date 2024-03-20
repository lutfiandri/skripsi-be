package controller

import (
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	InitHttpRoute()
	GetProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
}

type profileController struct {
	app     *fiber.App
	service service.ProfileService
}

func NewProfileController(app *fiber.App, service service.ProfileService) ProfileController {
	return &profileController{
		app:     app,
		service: service,
	}
}

func (controller *profileController) InitHttpRoute() {
	api := controller.app.Group("/profile")
	api.Get("/", middleware.NewAuthenticator(), controller.GetProfile)
	api.Put("/", middleware.NewAuthenticator(), controller.UpdateProfile)
}

func (controller *profileController) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	result, err := controller.service.GetProfile(c.Context(), claims)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *profileController) UpdateProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.UpdateProfileRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateProfileRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateProfile(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}
