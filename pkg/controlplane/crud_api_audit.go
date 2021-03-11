package controlplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/repo"
	"github.com/zhouzhuojie/iso8601ms"
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
// @param http_path query string false "HTTP Path"
// @param http_method query string false "HTTP Method"
// @param sent_http_status query int false "Sent HTTP Status"
// @param limit query int false "Limit"
// @param offset query int false "Offset"
// @param order_by query string false "Order By"
// @param order_desc query bool false "Order Desc"
// @success 200 {object} []apimodel.APIAudit
// @failure 400 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @router /api_audits [get]
func (cp *ControlPlane) QueryAPIAudits(c echo.Context) error {
	q := &apimodel.QueryAPIAudit{}
	err := c.Bind(q)
	if err != nil {
		return apimodel.NewHTTPError(err)
	}

	apiAudits, err := cp.Repo.QueryAPIAudits(c.Request().Context(), &repo.QueryAPIAuditOption{
		Domain:         q.Domain,
		Plane:          q.Plane,
		HTTPPath:       q.HTTPPath,
		HTTPMethod:     q.HTTPMethod,
		SentHTTPStatus: q.SentHTTPStatus,

		Limit:     q.Limit,
		Offset:    q.Offset,
		OrderBy:   q.OrderBy,
		OrderDesc: q.OrderDesc,
	})

	if err != nil {
		return apimodel.NewHTTPError(err)
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
			// Standardize the time to UTC before JSON serialization
			CreatedAt: iso8601ms.Time(item.CreatedAt.UTC()),
			UpdatedAt: iso8601ms.Time(item.UpdatedAt.UTC()),

			Plane:          item.Plane,
			Domain:         item.Domain,
			HTTPPath:       item.HTTPPath,
			HTTPMethod:     item.HTTPMethod,
			SentHTTPStatus: item.SentHTTPStatus,
		})
	}
	return ret
}
