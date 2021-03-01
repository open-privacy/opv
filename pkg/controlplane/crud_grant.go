package controlplane

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
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
	cg := &apimodel.CreateGrant{}
	err := c.Bind(cg)
	if err != nil {
		return apimodel.NewHTTPError(c, err, http.StatusBadRequest)
	}

	token, err := apimodel.NewToken("v1", cg.Domain)
	if err != nil {
		return apimodel.NewHTTPError(c, err, http.StatusInternalServerError)
	}

	grant := &apimodel.Grant{
		AllowedActions: cg.AllowedActions,
		Domain:         cg.Domain,
		Token:          token.String(),
	}

	// grouping policy for RBAC with domain pattern
	// example: https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_policy.csv
	//
	// default domain admin token =>  p, hash(token1234), domain, *, *, allow
	_, err = cp.CasbinEnforcer.AddPolicy(
		token.Hash(cp.Hasher),
		cg.Domain,
		"*",
		mergeAllowedActions(grant.AllowedActions),
		"allow",
	)
	if err != nil {
		return apimodel.NewHTTPError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, grant)
}

func mergeAllowedActions(actions []string) string {
	regexActions := []string{}
	for _, action := range actions {
		if action == "*" {
			return ".*"
		}
		regexActions = append(regexActions, fmt.Sprintf("(%s)", action))
	}
	return strings.Join(regexActions, "|")
}
