package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// ShowFactType godoc
// @tags Fact
// @summary Show a fact Type
// @description Show a fact type by ID
// @id show-fact-type-by-id
// @accept  json
// @produce  json
// @param id path string true "Fact Type ID"
// @success 200 {object} apimodel.FactType
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /fact_types/{id} [get]
func (dp *DataPlane) ShowFactType(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.FactType{})
}

// CreateFactType godoc
// @tags Fact
// @summary Create a fact type
// @description create a fact type
// @id create-fact-type
// @accept  json
// @produce  json
// @param createFact body apimodel.CreateFactType true "Create Fact Type Parameters"
// @success 200 {object} apimodel.FactType
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /fact_types [post]
func (dp *DataPlane) CreateFactType(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.FactType{})
}
