package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// ShowScope godoc
// @tags Scope
// @summary Show a scope
// @description Show scope by ID
// @id show-scope-by-id
// @accept  json
// @produce  json
// @security ApiKeyAuth
// @param id path string true "Scope ID"
// @success 200 {object} apimodel.Scope
// @failure 400 {object} apimodel.HTTPError
// @failure 404 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /scopes/{id} [get]
func (dp *DataPlane) ShowScope(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Scope{})
}

// CreateScope godoc
// @tags Scope
// @summary Create a scope
// @description Create a scope
// @id create-scope
// @accept  json
// @produce  json
// @security ApiKeyAuth
// @param createScope body apimodel.CreateScope	true "Create Scope parameters"
// @success 200 {object} apimodel.Scope
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /scopes [post]
func (dp *DataPlane) CreateScope(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Scope{})
}
