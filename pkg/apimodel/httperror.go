package apimodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/repo"
)

var (
	ErrInternalServerError = HTTPError{Message: "Internal server error", Code: http.StatusInternalServerError}
	ErrNotFound            = HTTPError{Message: "Resource not found", Code: http.StatusInternalServerError}
	ErrUnauthorized        = HTTPError{Message: "Unnauthorized", Code: http.StatusUnauthorized}
	ErrBadRequest          = func(message string) HTTPError { return HTTPError{Message: message, Code: http.StatusBadRequest} }
	ErrJSONMalformatted    = ErrBadRequest("JSON Malformatted")
)

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return he.Message
}

type HTTPErrorResponse struct {
	Error    HTTPError    `json:"error"`
}

func HTTPErrorHandler(err error, c echo.Context) {
	FormatHTTPError(c, HTTPErrorFactory(err))
}

func FormatHTTPError(c echo.Context, err HTTPError) error {
	er := HTTPErrorResponse{Error: err}
	return c.JSON(err.Code, er)
}

func HTTPErrorFactory(err error) HTTPError {
	switch err.(type) {
	default:
		return ErrInternalServerError
	case *echo.HTTPError:
		if (err == echo.ErrUnauthorized) {
			return ErrUnauthorized
		} else {
			return ErrInternalServerError
		}
	case *repo.NotFoundError:
		return ErrNotFound
	case *repo.ValidationError:
		return ErrBadRequest(err.Error())
	}
}
