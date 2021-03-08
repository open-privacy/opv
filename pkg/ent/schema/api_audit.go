package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// APIAudit holds the schema definition for the APIAudit entity.
type APIAudit struct {
	ent.Schema
}

// Fields of the APIAudit.
func (APIAudit) Fields() []ent.Field {
	return []ent.Field{
		ID("api_audit"),
		field.String("plane"),
		field.String("hashed_grant_token").Sensitive().Optional(),
		field.String("domain").Optional(),
		field.String("http_path").Optional(),
		field.String("http_method").Optional(),
		field.Int("sent_http_status").Optional(),
	}
}

// Edges of the APIAudit.
func (APIAudit) Edges() []ent.Edge {
	return nil
}

// Mixin of the APIAudit
func (APIAudit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Indexes of the APIAudit
func (APIAudit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("plane"),
		index.Fields("hashed_grant_token"),
		index.Fields("domain"),
		index.Fields("http_path"),
		index.Fields("http_method"),
		index.Fields("sent_http_status"),
	}
}
