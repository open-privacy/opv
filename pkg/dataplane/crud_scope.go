package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// @Tags Scope
// @Summary Show a scope
// @Description Show scope by ID
// @ID show-scope-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Scope ID"
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
