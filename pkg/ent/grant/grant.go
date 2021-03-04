// Code generated by entc, DO NOT EDIT.

package grant

import (
	"time"
)

const (
	// Label holds the string label denoting the grant type in the database.
	Label = "grant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldHashedToken holds the string denoting the hashed_token field in the database.
	FieldHashedToken = "hashed_token"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldAllowedHTTPMethods holds the string denoting the allowed_http_methods field in the database.
	FieldAllowedHTTPMethods = "allowed_http_methods"

	// Table holds the table name of the grant in the database.
	Table = "grants"
)

// Columns holds all SQL columns for grant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHashedToken,
	FieldDomain,
	FieldVersion,
	FieldAllowedHTTPMethods,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
