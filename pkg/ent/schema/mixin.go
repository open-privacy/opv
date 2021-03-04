package schema

import (
	"fmt"
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
func DefaultID(prefix string) func() string {
	return func() string {
		return fmt.Sprintf("%s_%s", prefix, uniuri.NewLen(uniuri.UUIDLen))
	}
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Immutable(),
		field.Time("deleted_at").Optional(),
	}
}

// Indexes of the BaseMixin
func (BaseMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}
