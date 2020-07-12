package errors_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("error message", errors.New("some server error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "some server error", err.Causes[0])

	errBytes, _ := json.Marshal(err)
	fmt.Println(string(errBytes))
}

func TestNewBadRequestErr(t *testing.T) {
	err := NewBadRequestErr("bad request error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "bad request error message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundErr(t *testing.T) {
	err := NewNotFoundErr("not found error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not found error message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewNotAuthorizedError(t *testing.T) {
	err := NewNotAuthorizedError("not authorized error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "not authorized error message", err.Message)
	assert.EqualValues(t, "not_authorized", err.Error)
}

func TestNewNotImplementedError(t *testing.T) {
	err := NewNotImplementedError()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Status)
	assert.EqualValues(t, "implement me!", err.Message)
	assert.EqualValues(t, "not_implemented", err.Error)
}