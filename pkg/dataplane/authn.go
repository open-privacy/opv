package dataplane

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var grantValidationMiddleware = middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	KeyLookup: "header:x-opv-grant-key",
	Validator: func(key string, c echo.Context) (bool, error) {
		// TODO grant key lookup and auth policy
		return true, nil
	},
})
