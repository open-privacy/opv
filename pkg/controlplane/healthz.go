package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roney492/opv/pkg/apimodel"
)

// Healthz godoc
// @tags Healthz
// @summary Show the health of the controlplane
// @description Show the health of the controlplane
// @id healthz
// @produce json
// @success 200 {object} apimodel.Healthz
// @failure 500 {object} echo.HTTPError
// @router /healthz [get]
func (cp *ControlPlane) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Healthz{Status: "OK"})
}
