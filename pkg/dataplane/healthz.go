package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// Healthz godoc
// @tags Healthz
// @summary Show the health of the dataplane
// @description Show the health of the dataplane
// @id healthz
// @produce json
// @success 200 {object} apimodel.Healthz
// @failure 500 {object} apimodel.HTTPError
// @router /healthz [get]
func (dp *DataPlane) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Healthz{Status: "OK"})
}
