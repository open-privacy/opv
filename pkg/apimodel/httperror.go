package apimodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/ent"
)

const (
	MessageNotFound = "Resource not found"
	MessageJSONMalformated = "JSON Malformated"
	MessageInternalServerError = "Internal server error."
)

// NewHTTPError creates a new HTTPError
func NewHTTPError(c echo.Context, message string, status int) error {
	er := HTTPErrorResponse{
		Error: HTTPError{
			Code:    status,
			Message: message,
	}}
	return c.JSON(status, er)
}

// NewEntError creates an error directly from entgo framework
// The status code is based on the ent error type
func NewEntError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}
	if ent.IsNotFound(err) {
		return NewHTTPError(c, MessageNotFound, http.StatusNotFound)
	}
	if ent.IsValidationError(err) {
		return NewHTTPError(c, err.Error(), http.StatusBadRequest)
	}
	if ent.IsConstraintError(err) {
		return NewHTTPError(c, err.Error(), http.StatusBadRequest)
	}

	return NewHTTPError(c, MessageInternalServerError, http.StatusInternalServerError)
}

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPErrorResponse struct {
	Error    HTTPError    `json:"error"`
}
