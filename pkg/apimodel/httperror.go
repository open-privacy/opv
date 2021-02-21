package apimodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/ent"
)

// NewHTTPError creates a new HTTPError
func NewHTTPError(c echo.Context, err error, status int) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return c.JSON(status, er)
}

func NewEntError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}
	if ent.IsNotFound(err) {
		return NewHTTPError(c, err, http.StatusNotFound)
	}
	if ent.IsValidationError(err) {
		return NewHTTPError(c, err, http.StatusBadRequest)
	}
	if ent.IsConstraintError(err) {
		return NewHTTPError(c, err, http.StatusBadRequest)
	}

	return NewHTTPError(c, err, http.StatusInternalServerError)
}

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
