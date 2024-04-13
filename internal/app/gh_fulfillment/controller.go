package gh_fulfillment

import (
	"log"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Fulfillment(c *fiber.Ctx) error
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (controller controller) Fulfillment(c *fiber.Ctx) error {
	var request Request
	parseOption := helper.ParseOptions{ParseBody: true}
	err := helper.ParseAndValidateRequest[Request](c, &request, parseOption)
	helper.PanicIfErr(err)

	log.Printf("\ngoogle fulfillment: %+v\n", request)

	switch request.Inputs[0].Intent {
	case constant.GhActionSync:
		response := controller.service.Sync(c, request)
		return c.JSON(response)
	case constant.GhActionQuery:
		response := controller.service.Query(c, request)
		return c.JSON(response)
	case constant.GhActionExecute:
		response := controller.service.Execute(c, request)
		return c.JSON(response)
	case constant.GhActionDisconnect:
		response := controller.service.Disconnect(c, request)
		return c.JSON(response)

	default:
		return c.JSON(map[string]any{})
	}
}
