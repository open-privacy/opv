package dataplane

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	dataplanedocs "github.com/open-privacy/opv/cmd/dataplane/docs"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/httputil"
)

type DataPlane struct{}

func NewDataPlane() *DataPlane {
	return &DataPlane{}
}

func StartServer() {
	e := echo.New()

	if config.ENV.DataPlaneCORSEnabled {
		e.Use(middleware.CORS())
	}

	dp := NewDataPlane()

	group := e.Group("/api/v1")
	group.GET("/scopes/:id", dp.ShowScope)

	hostport := fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	dataplanedocs.SwaggerInfo.Host = hostport
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(hostport))
}

// ShowScope godoc
// @Summary Show a scope
// @Description get scope by ID
// @ID get-scope-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Scope ID"
// @Success 200 {object} apimodel.Scope
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /scopes/{id} [get]
func (dp *DataPlane) ShowScope(c echo.Context) error {
	if 1 > 2 {
		return httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("not possible"))
	}
	return c.JSON(http.StatusOK, apimodel.Scope{})
}
