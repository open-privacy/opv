package controlplane

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/repo"
)

type auditLogContext struct {
	echo.Context
}

func (a *auditLogContext) mapCreateAPIOption() *repo.CreateAPIAuditOption {
	return &repo.CreateAPIAuditOption{
		Plane:            repo.DataplaneName,
		HashedGrantToken: nil, // control plane doesn't have grant token validation yet
		Domain:           nil, // control plane doesn't have grant token validation yet
		HTTPMethod:       &a.Request().Method,
		HTTPPath:         &a.Request().URL.Path,
		SentHTTPStatus:   &a.Response().Status,
	}
}

func (cp *ControlPlane) middlewareAPIAudit() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := next(c)

			go func() {
				ac := &auditLogContext{c}
				_, err := cp.Repo.CreateAPIAudit(context.Background(), ac.mapCreateAPIOption())
				if err != nil {
					cp.Logger.Error(err)
				}
			}()

			return res
		}
	}
}
