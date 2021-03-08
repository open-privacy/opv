package dataplane

import (
	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/repo"
)

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
		HTTPMethod:       a.getContextStringPtr(contextAuthzAct),
		HTTPPath:         a.getContextStringPtr(contextAuthzObj),
		SentHTTPStatus:   &a.Response().Status,
	}
}

func (dp *DataPlane) auditLogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := next(c)

			go func() {
				ac := &auditLogContext{c}
				dp.Repo.CreateAPIAudit(ac.Request().Context(), ac.mapCreateAPIOption())
			}()

			return res
		}
	}
}
