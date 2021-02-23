package apimodel

// CreateGrant represents the request object of creating a grant
type CreateGrant struct {
	Plane string `json:"plane"`
}

// Grant represents the grant object
type Grant struct {
	ID    string `json:"id"`
	Plane string `json:"plane"`
	Token string `json:"token"`
}
