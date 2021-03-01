package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Fact holds the schema definition for the Fact entity.
type Fact struct {
	ent.Schema
}

// Fields of the Fact.
func (Fact) Fields() []ent.Field {
	return []ent.Field{
		field.String("hashed_value").Sensitive(),
		field.String("encrypted_value").Sensitive(),
		field.String("domain"),
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

// Indexes of the Fact
func (Fact) Indexes() []ent.Index {
	return []ent.Index{
		// unique hashed_value constraint on same scope and fact_type
		index.Fields("hashed_value").Edges("scope", "fact_type").Unique(),
		index.Fields("domain"),
	}
}
