package apimodel

// CreateGrant represents the request object of creating a grant
type CreateGrant struct {
	Domain             string   `json:"domain" validate:"fqdn"`
	AllowedHTTPMethods []string `json:"allowed_http_methods" validate:"gt=0,dive,oneof=* GET POST PUT DELETE"`

	// TODO we can add more predicates here to limit how we want to expose
	// objects that the grant can operate on
}

// Grant represents the grant object
type Grant struct {
	Token              string   `json:"token"`
	Domain             string   `json:"domain"`
	AllowedHTTPMethods []string `json:"allowed_http_methods"`
}
