package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Scope holds the schema definition for the Scope entity.
type Scope struct {
	ent.Schema
}

// DefaultNonce generates a secure random nonce
func DefaultNonce() string {
	return uuid.New().String()
}

// Fields of the Scope.
func (Scope) Fields() []ent.Field {
	return []ent.Field{
		field.String("custom_id"),
		field.String("nonce").Sensitive().DefaultFunc(DefaultNonce),
	}
}

// Edges of the Scope.
func (Scope) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("facts", Fact.Type),
	}
}

// Mixin of the Scope
func (Scope) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

// Indexes of the Scope
func (Scope) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("custom_id").Unique(),
	}
}
