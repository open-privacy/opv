package dataplane

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/open-privacy/opv/pkg/apimodel"
)

const (
	headerOPVGrantToken = "x-opv-grant-token"

	contextAuthzSub = "authz.sub"
	contextAuthzAct = "authz.act"
	contextAuthzObj = "authz.obj"
	contextAuthzDom = "authz.dom"
)

func (dp *DataPlane) grantValidationMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: fmt.Sprintf("header:%s", headerOPVGrantToken),
		Validator: func(key string, c echo.Context) (bool, error) {
			token := &apimodel.GrantToken{}
			if err := token.ParseFromString(key); err != nil {
				return false, err
			}

			sub := token.Hash(dp.Hasher)
			act := c.Request().Method
			obj := c.Request().URL.Path
			dom := token.Domain

			c.Set(contextAuthzSub, sub)
			c.Set(contextAuthzAct, act)
			c.Set(contextAuthzObj, obj)
			c.Set(contextAuthzDom, dom)
			return dp.Enforcer.Enforce(sub, dom, obj, act)
		},
	})
}

func currentDomain(c echo.Context) string {
	return c.Get(contextAuthzDom).(string)
}
