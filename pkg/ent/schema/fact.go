package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Fact holds the schema definition for the Fact entity.
type Fact struct {
	ent.Schema
}

// Fields of the Fact.
func (Fact) Fields() []ent.Field {
	return []ent.Field{
		field.String("encrypted_value"),
	}
}

// Edges of the Fact.
func (Fact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scope", Scope.Type).Ref("facts").Unique(),
		edge.From("fact_type", FactType.Type).Ref("facts").Unique(),
	}
}

// Mixin of the Fact
func (Fact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}
