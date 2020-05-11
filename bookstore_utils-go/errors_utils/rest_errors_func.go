package errors_utils

import (
	"net/http"
)

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}

func NewInternalServerError(message string) *RestErr  {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}