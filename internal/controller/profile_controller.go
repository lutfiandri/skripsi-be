package controller

import (
	"log"

	"skripsi-be/internal/middleware"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	InitHttpRoute()
	GetProfile(c *fiber.Ctx) error
	EditProfile(c *fiber.Ctx) error
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
	api.Put("/", controller.EditProfile)
}

func (controller *profileController) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)
	log.Println(claims)

	result, err := controller.service.GetProfile(c.Context(), claims)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *profileController) EditProfile(c *fiber.Ctx) error {
	panic("unimplemented")
}
