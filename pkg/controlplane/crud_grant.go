package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roney492/opv/pkg/apimodel"
	"github.com/roney492/opv/pkg/repo"
)

// CreateGrant is the endpoint for creating a new grant
// @tags Grant
// @summary Create a grant
// @description Create a grant
// @id create-grant
// @accept json
// @produce json
// @param createGrant body apimodel.CreateGrant true "Create Grant parameters"
// @success 200 {object} apimodel.Grant
// @failure 400 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @router /grants [post]
func (cp *ControlPlane) CreateGrant(c echo.Context) error {
	ctx := c.Request().Context()
	cg := &apimodel.CreateGrant{}
	err := c.Bind(cg)
	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	if err := cp.Validator.Struct(cg); err != nil {
		return apimodel.NewHTTPError(err)
	}

	token, err := apimodel.NewToken("v1", cg.Domain)
	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	g, err := cp.Repo.CreateGrant(ctx, &repo.CreateGrantOption{
		HashedGrantToken:   token.Hash(cp.Hasher),
		Domain:             cg.Domain,
		Version:            token.Version,
		AllowedHTTPMethods: cg.AllowedHTTPMethods,
		Paths:              cg.Paths,
	})
	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	return c.JSON(http.StatusOK, &apimodel.Grant{
		Token:              token.String(),
		Domain:             g.Domain,
		AllowedHTTPMethods: cg.AllowedHTTPMethods,
		Paths:              g.Paths,
	})
}
