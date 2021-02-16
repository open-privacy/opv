package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/mixin"
)

// FactType holds the schema definition for the FactType entity.
type FactType struct {
	ent.Schema
}

// Fields of the FactType.
func (FactType) Fields() []ent.Field {
	return nil
}

// Edges of the FactType.
func (FactType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("facts", Fact.Type),
	}
}

// Mixin of the FactType
func (FactType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}
