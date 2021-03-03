package apimodel

import (
	"github.com/labstack/echo/v4"
)

// NewHTTPError creates a new HTTPError
func NewHTTPError(c echo.Context, err error, status int) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return c.JSON(status, er)
}

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
