package apimodel

// CreateGrant represents the request object of creating a grant
type CreateGrant struct {
	Domain             string   `json:"domain" validate:"fqdn"`
	AllowedHTTPMethods []string `json:"allowed_http_methods" validate:"gt=0,dive,oneof=* GET POST PUT DELETE"`

	// Paths represent the path the grant token can access
	// We leverage KeyMatch2 to define paths https://github.com/casbin/casbin/blob/v2.25.5/util/builtin_operators_test.go#L88-L117
	// By default if "paths" is not set, by default it's "*". For example,
	//
	//     *
	//     /api/v1/facts
	//     /api/v1/facts/:id
	//     /api/v1/*
	//     /js/v1/facts
	//
	Paths []string `json:"paths"`

	// TODO we can add more predicates here to limit how we want to expose
	// objects that the grant can operate on
}

// Grant represents the grant object
type Grant struct {
	Token              string   `json:"token"`
	Domain             string   `json:"domain"`
	AllowedHTTPMethods []string `json:"allowed_http_methods"`
	Paths              []string `json:"paths,omitempty"`
}
