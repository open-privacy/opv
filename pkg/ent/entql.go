// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/predicate"
	"github.com/open-privacy/opv/pkg/ent/scope"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   fact.Table,
			Columns: fact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fact.FieldID,
			},
		},
		Type: "Fact",
		Fields: map[string]*sqlgraph.FieldSpec{
			fact.FieldCreateTime:     {Type: field.TypeTime, Column: fact.FieldCreateTime},
			fact.FieldUpdateTime:     {Type: field.TypeTime, Column: fact.FieldUpdateTime},
			fact.FieldEncryptedValue: {Type: field.TypeString, Column: fact.FieldEncryptedValue},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   facttype.Table,
			Columns: facttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: facttype.FieldID,
			},
		},
		Type: "FactType",
		Fields: map[string]*sqlgraph.FieldSpec{
			facttype.FieldCreateTime: {Type: field.TypeTime, Column: facttype.FieldCreateTime},
			facttype.FieldUpdateTime: {Type: field.TypeTime, Column: facttype.FieldUpdateTime},
			facttype.FieldSlug:       {Type: field.TypeString, Column: facttype.FieldSlug},
			facttype.FieldBuiltin:    {Type: field.TypeBool, Column: facttype.FieldBuiltin},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   scope.Table,
			Columns: scope.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scope.FieldID,
			},
		},
		Type: "Scope",
		Fields: map[string]*sqlgraph.FieldSpec{
			scope.FieldCreateTime: {Type: field.TypeTime, Column: scope.FieldCreateTime},
			scope.FieldUpdateTime: {Type: field.TypeTime, Column: scope.FieldUpdateTime},
			scope.FieldNonce:      {Type: field.TypeUUID, Column: scope.FieldNonce},
			scope.FieldType:       {Type: field.TypeString, Column: scope.FieldType},
			scope.FieldExpiresAt:  {Type: field.TypeTime, Column: scope.FieldExpiresAt},
		},
	}
	graph.MustAddE(
		"scope",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.ScopeTable,
			Columns: []string{fact.ScopeColumn},
			Bidi:    false,
		},
		"Fact",
		"Scope",
	)
	graph.MustAddE(
		"fact_type",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.FactTypeTable,
			Columns: []string{fact.FactTypeColumn},
			Bidi:    false,
		},
		"Fact",
		"FactType",
	)
	graph.MustAddE(
		"facts",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
			Bidi:    false,
		},
		"FactType",
		"Fact",
	)
	graph.MustAddE(
		"facts",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
		},
		"Scope",
		"Fact",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (fq *FactQuery) addPredicate(pred func(s *sql.Selector)) {
	fq.predicates = append(fq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the FactQuery builder.
func (fq *FactQuery) Filter() *FactFilter {
	return &FactFilter{fq}
}

// addPredicate implements the predicateAdder interface.
func (m *FactMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the FactMutation builder.
func (m *FactMutation) Filter() *FactFilter {
	return &FactFilter{m}
}

// FactFilter provides a generic filtering capability at runtime for FactQuery.
type FactFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *FactFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *FactFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(fact.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *FactFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(fact.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *FactFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(fact.FieldUpdateTime))
}

// WhereEncryptedValue applies the entql string predicate on the encrypted_value field.
func (f *FactFilter) WhereEncryptedValue(p entql.StringP) {
	f.Where(p.Field(fact.FieldEncryptedValue))
}

// WhereHasScope applies a predicate to check if query has an edge scope.
func (f *FactFilter) WhereHasScope() {
	f.Where(entql.HasEdge("scope"))
}

// WhereHasScopeWith applies a predicate to check if query has an edge scope with a given conditions (other predicates).
func (f *FactFilter) WhereHasScopeWith(preds ...predicate.Scope) {
	f.Where(entql.HasEdgeWith("scope", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasFactType applies a predicate to check if query has an edge fact_type.
func (f *FactFilter) WhereHasFactType() {
	f.Where(entql.HasEdge("fact_type"))
}

// WhereHasFactTypeWith applies a predicate to check if query has an edge fact_type with a given conditions (other predicates).
func (f *FactFilter) WhereHasFactTypeWith(preds ...predicate.FactType) {
	f.Where(entql.HasEdgeWith("fact_type", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (ftq *FactTypeQuery) addPredicate(pred func(s *sql.Selector)) {
	ftq.predicates = append(ftq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the FactTypeQuery builder.
func (ftq *FactTypeQuery) Filter() *FactTypeFilter {
	return &FactTypeFilter{ftq}
}

// addPredicate implements the predicateAdder interface.
func (m *FactTypeMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the FactTypeMutation builder.
func (m *FactTypeMutation) Filter() *FactTypeFilter {
	return &FactTypeFilter{m}
}

// FactTypeFilter provides a generic filtering capability at runtime for FactTypeQuery.
type FactTypeFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *FactTypeFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *FactTypeFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(facttype.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *FactTypeFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(facttype.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *FactTypeFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(facttype.FieldUpdateTime))
}

// WhereSlug applies the entql string predicate on the slug field.
func (f *FactTypeFilter) WhereSlug(p entql.StringP) {
	f.Where(p.Field(facttype.FieldSlug))
}

// WhereBuiltin applies the entql bool predicate on the builtin field.
func (f *FactTypeFilter) WhereBuiltin(p entql.BoolP) {
	f.Where(p.Field(facttype.FieldBuiltin))
}

// WhereHasFacts applies a predicate to check if query has an edge facts.
func (f *FactTypeFilter) WhereHasFacts() {
	f.Where(entql.HasEdge("facts"))
}

// WhereHasFactsWith applies a predicate to check if query has an edge facts with a given conditions (other predicates).
func (f *FactTypeFilter) WhereHasFactsWith(preds ...predicate.Fact) {
	f.Where(entql.HasEdgeWith("facts", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *ScopeQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ScopeQuery builder.
func (sq *ScopeQuery) Filter() *ScopeFilter {
	return &ScopeFilter{sq}
}

// addPredicate implements the predicateAdder interface.
func (m *ScopeMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ScopeMutation builder.
func (m *ScopeMutation) Filter() *ScopeFilter {
	return &ScopeFilter{m}
}

// ScopeFilter provides a generic filtering capability at runtime for ScopeQuery.
type ScopeFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ScopeFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ScopeFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(scope.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *ScopeFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(scope.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *ScopeFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(scope.FieldUpdateTime))
}

// WhereNonce applies the entql [16]byte predicate on the nonce field.
func (f *ScopeFilter) WhereNonce(p entql.ValueP) {
	f.Where(p.Field(scope.FieldNonce))
}

// WhereType applies the entql string predicate on the type field.
func (f *ScopeFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(scope.FieldType))
}

// WhereExpiresAt applies the entql time.Time predicate on the expires_at field.
func (f *ScopeFilter) WhereExpiresAt(p entql.TimeP) {
	f.Where(p.Field(scope.FieldExpiresAt))
}

// WhereHasFacts applies a predicate to check if query has an edge facts.
func (f *ScopeFilter) WhereHasFacts() {
	f.Where(entql.HasEdge("facts"))
}

// WhereHasFactsWith applies a predicate to check if query has an edge facts with a given conditions (other predicates).
func (f *ScopeFilter) WhereHasFactsWith(preds ...predicate.Fact) {
	f.Where(entql.HasEdgeWith("facts", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
