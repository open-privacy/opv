package dataplane

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/repo"
)

const (
	headerOPVGrantToken = "x-opv-grant-token"

	contextAuthzSub = "authz.sub"
	contextAuthzAct = "authz.act"
	contextAuthzObj = "authz.obj"
	contextAuthzDom = "authz.dom"
)

func (dp *DataPlane) middlewareGrantValidation() echo.MiddlewareFunc {
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

type auditLogContext struct {
	echo.Context
}

func (a *auditLogContext) getContextStringPtr(name string) *string {
	val := a.Get(name)
	if val == nil {
		return nil
	}
	s, ok := val.(string)
	if !ok {
		return nil
	}
	return &s
}

func (a *auditLogContext) mapCreateAPIOption() *repo.CreateAPIAuditOption {
	return &repo.CreateAPIAuditOption{
		Plane:            repo.DataplaneName,
		HashedGrantToken: a.getContextStringPtr(contextAuthzSub),
		Domain:           a.getContextStringPtr(contextAuthzDom),
		HTTPMethod:       &a.Request().Method,
		HTTPPath:         &a.Request().URL.Path,
		SentHTTPStatus:   &a.Response().Status,
	}
}

func (dp *DataPlane) middlewareAPIAudit() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := next(c)

			go func() {
				ac := &auditLogContext{c}
				_, err := dp.Repo.CreateAPIAudit(context.Background(), ac.mapCreateAPIOption())
				if err != nil {
					dp.Logger.Error(err)
				}
			}()

			return res
		}
	}
}
