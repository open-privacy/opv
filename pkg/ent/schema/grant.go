package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Grant holds the schema definition for the Grant entity.
type Grant struct {
	ent.Schema
}

// Fields of the Grant.
func (Grant) Fields() []ent.Field {
	return []ent.Field{
		ID("grant"),
		field.String("hashed_grant_token").Sensitive(),
		field.String("domain"),
		field.String("version"),
		field.String("allowed_http_methods"),
		field.Strings("paths"),
	}
}

// Mixin of the Grant
func (Grant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Indexes of the Grant
func (Grant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hashed_grant_token"),
		index.Fields("domain"),
	}
}
