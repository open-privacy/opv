package apimodel

import "time"

// APIAudit ...
type APIAudit struct {
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
	Plane            string    `json:"plane,omitempty"`
	HashedGrantToken string    `json:"-"`
	Domain           string    `json:"domain,omitempty"`
	HTTPPath         string    `json:"http_path,omitempty"`
	HTTPMethod       string    `json:"http_method,omitempty"`
	SentHTTPStatus   int       `json:"sent_http_status,omitempty"`
}
