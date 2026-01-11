package handler

import "github.com/pixlcrashr/roomy/pkg/api/ogen/gen"

// NewErrorResponse creates a new ErrorResponse with the given code and message.
func NewErrorResponse(code, message string) gen.ErrorResponse {
	return gen.ErrorResponse{
		Error: gen.ErrorResponseError{
			Code:    code,
			Message: message,
		},
	}
}

// InternalError creates an internal server error response.
func InternalError(message string) gen.ErrorResponse {
	return NewErrorResponse("INTERNAL_ERROR", message)
}

// NotFoundError creates a not found error response.
func NotFoundError(message string) gen.ErrorResponse {
	return NewErrorResponse("NOT_FOUND", message)
}

// BadRequestError creates a bad request error response.
func BadRequestError(message string) gen.ErrorResponse {
	return NewErrorResponse("BAD_REQUEST", message)
}
