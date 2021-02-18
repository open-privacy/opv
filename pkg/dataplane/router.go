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
)

type DataPlane struct{}

func NewDataPlane() *DataPlane {
	return &DataPlane{}
}

func StartServer() {
	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	if config.ENV.DataPlaneCORSEnabled {
		e.Use(middleware.CORS())
	}

	dp := NewDataPlane()

	group := e.Group("/api/v1")
	group.POST("/scopes", dp.CreateScope)
	group.GET("/scopes/:id", dp.ShowScope)
	group.POST("/facts", dp.CreateFact)
	group.GET("/facts/:id", dp.ShowFact)
	group.POST("/fact_types", dp.CreateFactType)
	group.GET("/fact_types/:id", dp.ShowFactType)

	hostport := fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	dataplanedocs.SwaggerInfo.Host = hostport
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(hostport))
}

// @Tags Scope
// @Summary Show a scope
// @Description Show scope by ID
// @ID show-scope-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Scope ID"
// @Success 200 {object} apimodel.Scope
// @Failure 400 {object} apimodel.HTTPError
// @Failure 404 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /scopes/{id} [get]
func (dp *DataPlane) ShowScope(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Scope{})
}

// @Tags Scope
// @Summary Create a scope
// @Description Create a scope
// @ID create-scope
// @Accept  json
// @Produce  json
// @Param createScope body apimodel.CreateScope	true "Create Scope parameters"
// @Success 200 {object} apimodel.Scope
// @Failure 400 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /scopes [post]
func (dp *DataPlane) CreateScope(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Scope{})
}

// @Tags Fact
// @Summary Show a fact
// @Description Show a fact by ID
// @ID show-fact-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Fact ID"
// @Success 200 {object} apimodel.Fact
// @Failure 400 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /facts/{id} [get]
func (dp *DataPlane) ShowFact(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Fact{})
}

// @Tags Fact
// @Summary Create a fact
// @Description create a fact
// @ID create-fact
// @Accept  json
// @Produce  json
// @Param createFact body apimodel.CreateFact true "Create Fact Parameters"
// @Success 200 {object} apimodel.Fact
// @Failure 400 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /facts [post]
func (dp *DataPlane) CreateFact(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Fact{})
}

// @Tags Fact
// @Summary Show a fact Type
// @Description Show a fact type by ID
// @ID show-fact-type-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Fact Type ID"
// @Success 200 {object} apimodel.FactType
// @Failure 400 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /fact_types/{id} [get]
func (dp *DataPlane) ShowFactType(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.FactType{})
}

// @Tags Fact
// @Summary Create a fact type
// @Description create a fact type
// @ID create-fact-type
// @Accept  json
// @Produce  json
// @Param createFact body apimodel.CreateFactType true "Create Fact Type Parameters"
// @Success 200 {object} apimodel.FactType
// @Failure 400 {object} apimodel.HTTPError
// @Failure 500 {object} apimodel.HTTPError
// @Router /fact_types [post]
func (dp *DataPlane) CreateFactType(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.FactType{})
}
