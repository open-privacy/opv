// Code generated by entc, DO NOT EDIT.

package facttype

import (
	"time"
)

const (
	// Label holds the string label denoting the facttype type in the database.
	Label = "fact_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldBuiltin holds the string denoting the builtin field in the database.
	FieldBuiltin = "builtin"
	// FieldValidation holds the string denoting the validation field in the database.
	FieldValidation = "validation"

	// EdgeFacts holds the string denoting the facts edge name in mutations.
	EdgeFacts = "facts"

	// Table holds the table name of the facttype in the database.
	Table = "fact_types"
	// FactsTable is the table the holds the facts relation/edge.
	FactsTable = "facts"
	// FactsInverseTable is the table name for the Fact entity.
	// It exists in this package in order to avoid circular dependency with the "fact" package.
	FactsInverseTable = "facts"
	// FactsColumn is the table column denoting the facts relation/edge.
	FactsColumn = "fact_type_facts"
)

// Columns holds all SQL columns for facttype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSlug,
	FieldBuiltin,
	FieldValidation,
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
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultBuiltin holds the default value on creation for the "builtin" field.
	DefaultBuiltin bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
