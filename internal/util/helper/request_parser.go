package helper

import (
	"reflect"

	"skripsi-be/internal/model/rest"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ParseOptions struct {
	ParseBody   bool
	ParseQuery  bool
	ParseParams bool
}

func ParseAndValidateRequest[T any](c *fiber.Ctx, request *T, options ParseOptions) error {
	// reject if request is not pointer
	value := reflect.ValueOf(request)
	if value.Kind() != reflect.Pointer {
		return fiber.NewError(fiber.StatusInternalServerError, "request must be a pointer")
	}

	// parse body
	if options.ParseBody {
		if err := c.BodyParser(request); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}
	}

	// parse query
	if options.ParseQuery {
		if err := c.QueryParser(request); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "can not parse query")
		}
	}

	// parse params
	if options.ParseParams {
		if err := c.ParamsParser(request); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "can not parse params")
		}
	}

	// validate
	if err := validate.Struct(request); err != nil {
		errs := err.(validator.ValidationErrors)
		validationErrorsResponse := parseValidationError[T](c, errs, *request)
		c.Locals("validation_errors_response", validationErrorsResponse)

		return err
	}

	return nil
}

func parseValidationError[T any](c *fiber.Ctx, errs validator.ValidationErrors, request T) []rest.ValidationErrorResponse {
	var validationErrorsResponse []rest.ValidationErrorResponse

	// get json tags
	jsonTags := make(map[string]string)
	t := reflect.TypeOf(request)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTags[field.Name] = field.Tag.Get("json")
	}

	// create error list
	for _, err := range errs {
		var element rest.ValidationErrorResponse

		element.FailedField = jsonTags[err.Field()]
		element.Tag = err.Tag()
		element.Value = err.Param()

		validationErrorsResponse = append(validationErrorsResponse, element)
	}

	return validationErrorsResponse
}
