package schema

import (
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

func DefaultID() string {
	return uniuri.NewLen(uniuri.UUIDLen)
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().DefaultFunc(DefaultID),
	}
}

// Indexes of the BaseMixin
func (BaseMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
