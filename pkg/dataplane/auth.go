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
			grant := &apimodel.Grant{Token: key}

			sub := grant.Hash(dp.Hasher)
			act := c.Request().Method
			obj := c.Request().URL.Path
			dom, err := grant.Domain()
			if err != nil {
				return false, err
			}

			c.Set(contextAuthzSub, sub)
			c.Set(contextAuthzAct, act)
			c.Set(contextAuthzObj, obj)
			c.Set(contextAuthzDom, dom)
			return dp.CasbinEnforcer.Enforce(sub, dom, obj, act)
		},
	})
}

func currentDomain(c echo.Context) string {
	return c.Get(contextAuthzDom).(string)
}
