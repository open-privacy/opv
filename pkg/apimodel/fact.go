package apimodel

// CreateFact represents the request to create a fact
type CreateFact struct {
	ScopeCustomID string `json:"scope_custom_id" validate:"required"`
	FactTypeSlug  string `json:"fact_type_slug" validate:"required"`
	Value         string `json:"value" validate:"required"`
}

// Fact represents a fact
type Fact struct {
	ID            string `json:"id"`
	ScopeCustomID string `json:"scope_custom_id"`
	FactTypeSlug  string `json:"fact_type_slug"`
	Domain        string `json:"domain"`
	Value         string `json:"value,omitempty"`
}
