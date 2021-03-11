package apimodel

import (
	"github.com/zhouzhuojie/iso8601ms"
)

// QueryAPIAudit is the query struct
type QueryAPIAudit struct {
	Domain         *string `query:"domain"`
	Plane          *string `query:"plane"`
	HTTPPath       *string `query:"http_path"`
	HTTPMethod     *string `query:"http_method"`
	SentHTTPStatus *int    `query:"sent_http_status"`

	Limit     *int    `query:"limit"`
	Offset    *int    `query:"offset"`
	OrderBy   *string `query:"order_by"`
	OrderDesc bool    `query:"order_desc"`
}

// APIAudit is the response apimodel
type APIAudit struct {
	CreatedAt iso8601ms.Time `json:"created_at,omitempty"`
	UpdatedAt iso8601ms.Time `json:"updated_at,omitempty"`
	Plane     string         `json:"plane,omitempty"`
	Domain    string         `json:"domain,omitempty"`

	HTTPPath       string `json:"http_path,omitempty"`
	HTTPMethod     string `json:"http_method,omitempty"`
	SentHTTPStatus int    `json:"sent_http_status,omitempty"`
}
