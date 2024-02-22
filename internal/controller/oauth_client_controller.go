package controller

import (
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type OAuthClientController interface {
	InitHttpRoute()
	CreateOAuthClient(c *fiber.Ctx) error
	GetOAuthClients(c *fiber.Ctx) error
	GetOAuthClient(c *fiber.Ctx) error
	UpdateOAuthClient(c *fiber.Ctx) error
	DeleteOAuthClient(c *fiber.Ctx) error
}

type oauthClientController struct {
	app     *fiber.App
	service service.OAuthClientService
}

func NewOAuthClientController(app *fiber.App, service service.OAuthClientService) OAuthClientController {
	return &oauthClientController{
		app:     app,
		service: service,
	}
}

func (controller *oauthClientController) InitHttpRoute() {
	api := controller.app.Group("/oauth2/clients")

	api.Post("/", middleware.NewAuthenticator(), controller.CreateOAuthClient)
	api.Get("/", middleware.NewAuthenticator(), controller.GetOAuthClients)
	api.Get("/:id", middleware.NewAuthenticator(), controller.GetOAuthClient)
	api.Put("/:id", middleware.NewAuthenticator(), controller.UpdateOAuthClient)
	api.Delete("/:id", middleware.NewAuthenticator(), controller.DeleteOAuthClient)
}

func (controller *oauthClientController) CreateOAuthClient(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.CreateOAuthClientRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.CreateOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateOAuthClient(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthClientController) GetOAuthClients(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	result, err := controller.service.GetOAuthClients(c.Context(), claims)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthClientController) GetOAuthClient(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.GetOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.GetOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetOAuthClient(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthClientController) UpdateOAuthClient(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.UpdateOAuthClientRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateOAuthClient(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthClientController) DeleteOAuthClient(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.DeleteOAuthClientRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.DeleteOAuthClientRequest](c, &request, parseOption); err != nil {
		return err
	}

	err := controller.service.DeleteOAuthClient(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
