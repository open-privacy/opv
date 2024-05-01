package apimodel

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/roney492/opv/pkg/repo"
)

// NewHTTPError creates a new echo.HTTPError from the given error
// We make sure the messages in the response body is sanitized without leaking sensitive information
func NewHTTPError(err error) *echo.HTTPError {
	switch e := err.(type) {
	case *echo.HTTPError:
		return e
		// TODO or we can sanitize the error message with default status text
		// return echo.NewHTTPError(e.Code)
	case repo.NotFoundError:
		return echo.ErrNotFound
	case repo.ValidationError:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	case repo.UnauthorizedError:
		return echo.ErrUnauthorized
	case govalidator.Error, govalidator.Errors, validator.ValidationErrors:
		return echo.NewHTTPError(http.StatusBadRequest, "Validation Error")

	default:
		return echo.ErrInternalServerError
	}
}
