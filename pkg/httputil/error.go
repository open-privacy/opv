package httputil

import "github.com/labstack/echo/v4"

// NewError example
func NewError(c echo.Context, status int, err error) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return c.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
