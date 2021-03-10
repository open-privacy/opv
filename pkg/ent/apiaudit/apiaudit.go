// Code generated by entc, DO NOT EDIT.

package apiaudit

import (
	"time"
)

const (
	// Label holds the string label denoting the apiaudit type in the database.
	Label = "api_audit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPlane holds the string denoting the plane field in the database.
	FieldPlane = "plane"
	// FieldHashedGrantToken holds the string denoting the hashed_grant_token field in the database.
	FieldHashedGrantToken = "hashed_grant_token"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldHTTPPath holds the string denoting the http_path field in the database.
	FieldHTTPPath = "http_path"
	// FieldHTTPMethod holds the string denoting the http_method field in the database.
	FieldHTTPMethod = "http_method"
	// FieldSentHTTPStatus holds the string denoting the sent_http_status field in the database.
	FieldSentHTTPStatus = "sent_http_status"

	// Table holds the table name of the apiaudit in the database.
	Table = "api_audits"
)

// Columns holds all SQL columns for apiaudit fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldPlane,
	FieldHashedGrantToken,
	FieldDomain,
	FieldHTTPPath,
	FieldHTTPMethod,
	FieldSentHTTPStatus,
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)