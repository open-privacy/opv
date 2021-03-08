package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// FactType holds the schema definition for the FactType entity.
type FactType struct {
	ent.Schema
}

// Fields of the FactType.
func (FactType) Fields() []ent.Field {
	return []ent.Field{
		ID("fact_type"),
		field.String("slug"),
		field.Bool("built_in").Default(false),
		field.String("validation").Optional(),
	}
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
	}
}

// Indexes of the FactType
func (FactType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}
