package schema

import (
	"crypto/rand"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

// BaseMixin is the base entity mixin
type BaseMixin struct {
	mixin.Schema
}

// DefaultULID generates a new ULID string
func DefaultULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().DefaultFunc(DefaultULID),
	}
}

// Indexes of the BaseMixin
func (BaseMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
