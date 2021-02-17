package apimodel

import (
	"time"

	"github.com/google/uuid"
)

// CreateFact represents the request to create a fact
type CreateFact struct {
	ScopeID      uuid.UUID `json:"scope_id"`
	FactTypeSlug string    `json:"fact_type_slug"`
	Value        string    `json:"value"`
}

// Fact represents a fact
type Fact struct {
	ID             uuid.UUID `json:"id"`
	ScopeID        uuid.UUID `json:"scope_id"`
	FactType       string    `json:"fact_type"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
	EncryptedValue string    `json:"encrypted_value"`
}

// CreateFactType represents the request to create a fact_type
type CreateFactType struct {
	Slug string `json:"slug"`
}

type FactType struct {
	ID         uuid.UUID `json:"id"`
	Slug       string    `json:"slug"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
