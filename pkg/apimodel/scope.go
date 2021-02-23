package apimodel

// CreateScope represents the request of creating a scope
type CreateScope struct {
	CustomID string `json:"type"`
}

// Scope represents the struct of a scope
type Scope struct {
	ID       string `json:"id"`
	CustomID string `json:"custom_id"`
}
