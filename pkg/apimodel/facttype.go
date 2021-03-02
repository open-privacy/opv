package apimodel

// CreateFactType represents the request to create a fact_type
type CreateFactType struct {
	Slug string `json:"slug"`
}

// FactType represents the fact_type api model struct
type FactType struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}
