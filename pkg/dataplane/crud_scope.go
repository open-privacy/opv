package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roney492/opv/pkg/apimodel"
	"github.com/roney492/opv/pkg/repo"
)

// QueryScopes godoc
// @tags Scope
// @summary Query scopes
// @description Query scopes
// @id query-scopes
// @produce json
// @security ApiKeyAuth
// @param custom_id query string false "get scopes by custom_id"
// @success 200 {object} []apimodel.Scope
// @failure 400 {object} echo.HTTPError
// @failure 404 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @router /scopes [get]
func (dp *DataPlane) QueryScopes(c echo.Context) error {
	s, err := dp.Repo.GetScope(c.Request().Context(), &repo.GetScopeOption{
		ScopeCustomID: c.QueryParam("custom_id"),
		Domain:        currentDomain(c),
	})
	if err != nil {
		return apimodel.NewHTTPError(err)
	}
	return c.JSON(http.StatusOK, apimodel.Scope{
		ID:       s.ID,
		CustomID: s.CustomID,
	})
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
// @failure 400 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @router /scopes [post]
func (dp *DataPlane) CreateScope(c echo.Context) error {
	cs := &apimodel.CreateScope{}
	err := c.Bind(cs)
	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	s, err := dp.Repo.CreateScope(c.Request().Context(), &repo.CreateScopeOption{
		ScopeCustomID: cs.CustomID,
		Domain:        currentDomain(c),
	})

	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	return c.JSON(http.StatusOK, apimodel.Scope{
		ID:       s.ID,
		CustomID: s.CustomID,
	})
}
