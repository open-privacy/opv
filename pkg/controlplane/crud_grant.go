package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// CreateGrants is the endpoint for creating a new grant
// @tags Grant
// @summary Create a grant
// @description Create a grant
// @id create-grant
// @accept  json
// @produce  json
// @param createGrant body apimodel.CreateGrant true "Create Grant parameters"
// @success 200 {object} apimodel.Grant
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /grants [post]
func (cp *ControlPlane) CreateGrants(c echo.Context) error {
	return c.JSON(http.StatusOK, apimodel.Grant{})
}
