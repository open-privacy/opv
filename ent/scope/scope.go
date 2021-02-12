// Code generated by entc, DO NOT EDIT.

package scope

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the scope type in the database.
	Label = "scope"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"

	// EdgeFacts holds the string denoting the facts edge name in mutations.
	EdgeFacts = "facts"

	// Table holds the table name of the scope in the database.
	Table = "scopes"
	// FactsTable is the table the holds the facts relation/edge.
	FactsTable = "facts"
	// FactsInverseTable is the table name for the Fact entity.
	// It exists in this package in order to avoid circular dependency with the "fact" package.
	FactsInverseTable = "facts"
	// FactsColumn is the table column denoting the facts relation/edge.
	FactsColumn = "scope_facts"
)

// Columns holds all SQL columns for scope fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNonce,
	FieldExpiresAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultNonce holds the default value on creation for the "nonce" field.
	DefaultNonce func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
