package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/repo"
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
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /grants [post]
func (cp *ControlPlane) CreateGrant(c echo.Context) error {
	ctx := c.Request().Context()
	cg := &apimodel.CreateGrant{}
	err := c.Bind(cg)
	if err != nil {
		return apimodel.FormatHTTPError(c, apimodel.ErrJSONMalformatted)
	}

	if err := cp.Validator.Struct(cg); err != nil {
		return cp.Repo.HandleError(ctx, err)
	}

	token, err := apimodel.NewToken("v1", cg.Domain)
	if err != nil {
		return cp.Repo.HandleError(ctx, err)
	}

	g, err := cp.Repo.CreateGrant(ctx, &repo.CreateGrantOption{
		HashedGrantToken:   token.Hash(cp.Hasher),
		Domain:             cg.Domain,
		Version:            token.Version,
		AllowedHTTPMethods: cg.AllowedHTTPMethods,
	})
	if err != nil {
		return cp.Repo.HandleError(ctx, err)
	}

	return c.JSON(http.StatusOK, &apimodel.Grant{
		Token:              token.String(),
		Domain:             g.Domain,
		AllowedHTTPMethods: cg.AllowedHTTPMethods,
	})
}
