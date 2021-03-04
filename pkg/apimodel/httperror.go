package apimodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/repo"
)

const (
	MessageNotFound = "Resource not found"
	MessageJSONMalformated = "JSON Malformated"
	MessageInternalServerError = "Internal server error"
	MessageValidationError = "Validation error"
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

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPErrorResponse struct {
	Error    HTTPError    `json:"error"`
}

func HTTPErrorHandler(err error, c echo.Context) {
	switch err.(type) {
	default:
		NewHTTPError(c, MessageInternalServerError, http.StatusInternalServerError)
	case *repo.NotFoundError:
		NewHTTPError(c, MessageNotFound, http.StatusNotFound)
	case *repo.ValidationError:
		message := MessageValidationError + ": " + err.Error()
		NewHTTPError(c, message, http.StatusBadRequest)
	}
}
