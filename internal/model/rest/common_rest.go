package rest

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success          bool                       `json:"success"`
	Message          string                     `json:"message"`
	ValidationErrors *[]ValidationErrorResponse `json:"validation_errors"`
}

type ValidationErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(message string, validationErrors *[]ValidationErrorResponse) ErrorResponse {
	return ErrorResponse{
		Success:          false,
		Message:          message,
		ValidationErrors: validationErrors,
	}
}
