package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Scope holds the schema definition for the Scope entity.
type Scope struct {
	ent.Schema
}

// Fields of the Scope.
func (Scope) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("nonce", uuid.UUID{}).Default(uuid.New),
		field.String("type").Optional(),
		field.Time("expires_at").Optional().Nillable(),
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
