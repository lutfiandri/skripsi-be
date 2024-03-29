package middleware

import (
	"errors"

	"skripsi-be/internal/interface/rest"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	var validationErrorsResponse *[]rest.ValidationErrorResponse
	if valErrsResponse, ok := c.Locals("validation_errors_response").([]rest.ValidationErrorResponse); ok {
		validationErrorsResponse = &valErrsResponse
	}

	switch err {
	case mongo.ErrNoDocuments:
		code = fiber.StatusNotFound
	}

	response := rest.NewErrorResponse(err.Error(), validationErrorsResponse)
	return c.Status(code).JSON(response)
}
