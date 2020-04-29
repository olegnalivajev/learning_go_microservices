package errors

import "net/http"

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
