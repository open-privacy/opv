package apimodel

// CreateFactType represents the request to create a fact_type
type CreateFactType struct {
	Slug       string `json:"slug"`
	Validation string `json:"validation"`
}

// FactType represents the fact_type api model struct
type FactType struct {
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	BuiltIn    bool   `json:"built_in"`
	Validation string `json:"validation"`
}
