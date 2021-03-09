package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/repo"
)

// QueryAPIAudits is the endpoint for querying a list of api audit logs
// @tags Audit
// @summary Query API Audits
// @description Query API Audits
// @id query-api-audits
// @accept json
// @produce json
// @param domain query string false "Domain"
// @param plane query string false "Plane"
// @success 200 {object} []apimodel.APIAudit
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /api_audits [get]
func (cp *ControlPlane) QueryAPIAudits(c echo.Context) error {
	var (
		domain *string
		plane  *string
	)
	if d := c.QueryParam("domain"); d != "" {
		domain = &d
	}
	if p := c.QueryParam("plane"); p != "" {
		plane = &p
	}

	apiAudits, err := cp.Repo.QueryAPIAudits(c.Request().Context(), &repo.QueryAPIAuditOption{
		Domain: domain,
		Plane:  plane,
	})
	if err != nil {
		return cp.Repo.HandleError(c.Request().Context(), err)
	}

	return c.JSON(http.StatusOK, mapAPIAudits(apiAudits))
}

func mapAPIAudits(r []*ent.APIAudit) []apimodel.APIAudit {
	ret := []apimodel.APIAudit{}
	if len(r) == 0 {
		return ret
	}

	for _, item := range r {
		ret = append(ret, apimodel.APIAudit{
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			DeletedAt:        item.DeletedAt,
			Plane:            item.Plane,
			HashedGrantToken: item.HashedGrantToken,
			Domain:           item.Domain,
			HTTPPath:         item.HTTPPath,
			HTTPMethod:       item.HTTPMethod,
			SentHTTPStatus:   item.SentHTTPStatus,
		})
	}
	return ret
}
