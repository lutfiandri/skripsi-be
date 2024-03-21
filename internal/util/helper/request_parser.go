package helper

import (
	"reflect"

	"skripsi-be/internal/interface/rest"

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
		validationErrorsResponse := parseValidationError[T](errs, *request)
		c.Locals("validation_errors_response", validationErrorsResponse)

		return err
	}

	return nil
}

func parseValidationError[T any](errs validator.ValidationErrors, request T) []rest.ValidationErrorResponse {
	var validationErrorsResponse []rest.ValidationErrorResponse

	// get json tags
	fieldNameMap := make(map[string]string)
	fieldTypeMap := make(map[string]string)
	t := reflect.TypeOf(request)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fieldName, fieldType := getFieldNameAndType(field)

		fieldNameMap[field.Name] = fieldName
		fieldTypeMap[field.Name] = fieldType

	}

	// create error list
	for _, err := range errs {
		var element rest.ValidationErrorResponse

		element.Field = fieldNameMap[err.Field()]
		element.Type = fieldTypeMap[err.Field()]
		element.Tag = err.Tag()
		element.Value = err.Param()

		validationErrorsResponse = append(validationErrorsResponse, element)
	}

	return validationErrorsResponse
}

func getFieldNameAndType(field reflect.StructField) (string, string) {
	fieldName := field.Tag.Get("json")
	fieldType := "json"

	if fieldName == "" {
		fieldName = field.Tag.Get("query")
		fieldType = "query"
	}

	if fieldName == "" {
		fieldName = field.Tag.Get("params")
		fieldType = "params"
	}

	return fieldName, fieldType
}
