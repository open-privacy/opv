package apimodel

import (
	"time"
)

// CreateFact represents the request to create a fact
type CreateFact struct {
	ScopeCustomID string `json:"scope_custom_id"`
	FactTypeSlug  string `json:"fact_type_slug"`
	Value         string `json:"value"`
}

// Fact represents a fact
type Fact struct {
	ID            string    `json:"id"`
	ScopeCustomID string    `json:"scope_custom_id"`
	FactTypeSlug  string    `json:"fact_type_slug"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	Value         string    `json:"value"`
}

// CreateFactType represents the request to create a fact_type
type CreateFactType struct {
	Slug string `json:"slug"`
}

type FactType struct {
	ID         string    `json:"id"`
	Slug       string    `json:"slug"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
