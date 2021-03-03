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
		field.String("hashed_token"),
		field.String("domain"),
		field.String("version"),
		field.String("allowed_http_methods"),
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
		index.Fields("hashed_token"),
		index.Fields("domain"),
	}
}
