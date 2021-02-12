package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Scope holds the schema definition for the Scope entity.
type Scope struct {
	ent.Schema
}

// Fields of the Scope.
func (Scope) Fields() []ent.Field {
	return []ent.Field{
		// default fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),

		field.UUID("nonce", uuid.UUID{}).Default(uuid.New),
		field.Time("expires_at").Optional().Nillable(),
	}
}

// Edges of the Scope.
func (Scope) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("facts", Fact.Type),
	}
}
