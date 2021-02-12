package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Fact holds the schema definition for the Fact entity.
type Fact struct {
	ent.Schema
}

// Fields of the Fact.
func (Fact) Fields() []ent.Field {
	return []ent.Field{
		// default fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),

		field.Bytes("encrypted_value"),
	}
}

// Edges of the Fact.
func (Fact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scope", Scope.Type).Ref("facts").Unique(),
	}
}
