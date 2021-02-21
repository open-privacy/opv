package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// @Tags Fact
// @Summary Show a fact Type
// @Description Show a fact type by ID
// @ID show-fact-type-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Fact Type ID"
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
