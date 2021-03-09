package dataplane

import (
	"context"

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
		HTTPMethod:       &a.Request().Method,
		HTTPPath:         &a.Request().URL.Path,
		SentHTTPStatus:   &a.Response().Status,
	}
}

func (dp *DataPlane) auditLogMiddleware() echo.MiddlewareFunc {
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
