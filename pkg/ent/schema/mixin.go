package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/dchest/uniuri"
)

// BaseMixin is the base entity mixin
type BaseMixin struct {
	mixin.Schema
}

// DefaultID generates BaseMixin's default ID
func DefaultID() string {
	return uniuri.NewLen(uniuri.UUIDLen)
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().DefaultFunc(DefaultID),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Immutable(),
	}
}

// Indexes of the BaseMixin
func (BaseMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
