package controller

import (
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/service"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type OAuthScopeController interface {
	InitHttpRoute()
	CreateOAuthScope(c *fiber.Ctx) error
	GetOAuthScopes(c *fiber.Ctx) error
	GetOAuthScope(c *fiber.Ctx) error
	UpdateOAuthScope(c *fiber.Ctx) error
	DeleteOAuthScope(c *fiber.Ctx) error
}

type oauthScopeController struct {
	app     *fiber.App
	service service.OAuthScopeService
}

func NewOAuthScopeController(app *fiber.App, service service.OAuthScopeService) OAuthScopeController {
	return &oauthScopeController{
		app:     app,
		service: service,
	}
}

func (controller *oauthScopeController) InitHttpRoute() {
	api := controller.app.Group("/oauth2/scopes")

	api.Post("/", middleware.NewAuthenticator(), controller.CreateOAuthScope)
	api.Get("/", middleware.NewAuthenticator(), controller.GetOAuthScopes)
	api.Get("/:id", middleware.NewAuthenticator(), controller.GetOAuthScope)
	api.Put("/:id", middleware.NewAuthenticator(), controller.UpdateOAuthScope)
	api.Delete("/:id", middleware.NewAuthenticator(), controller.DeleteOAuthScope)
}

func (controller *oauthScopeController) CreateOAuthScope(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.CreateOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[rest.CreateOAuthScopeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.CreateOAuthScope(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthScopeController) GetOAuthScopes(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	result, err := controller.service.GetOAuthScopes(c.Context(), claims)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthScopeController) GetOAuthScope(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.GetOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.GetOAuthScopeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.GetOAuthScope(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthScopeController) UpdateOAuthScope(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.UpdateOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.UpdateOAuthScopeRequest](c, &request, parseOption); err != nil {
		return err
	}

	result, err := controller.service.UpdateOAuthScope(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(result)

	return c.JSON(response)
}

func (controller *oauthScopeController) DeleteOAuthScope(c *fiber.Ctx) error {
	claims := c.Locals("claims").(rest.JWTClaims)

	var request rest.DeleteOAuthScopeRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[rest.DeleteOAuthScopeRequest](c, &request, parseOption); err != nil {
		return err
	}

	err := controller.service.DeleteOAuthScope(c.Context(), claims, request)
	if err != nil {
		return err
	}

	response := rest.NewSuccessResponse(nil)

	return c.JSON(response)
}
