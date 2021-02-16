// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/open-privacy-vault/opv/ent/fact"
	"github.com/open-privacy-vault/opv/ent/predicate"
	"github.com/open-privacy-vault/opv/ent/scope"
)

// ScopeUpdate is the builder for updating Scope entities.
type ScopeUpdate struct {
	config
	hooks    []Hook
	mutation *ScopeMutation
}

// Where adds a new predicate for the ScopeUpdate builder.
func (su *ScopeUpdate) Where(ps ...predicate.Scope) *ScopeUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetNonce sets the "nonce" field.
func (su *ScopeUpdate) SetNonce(u uuid.UUID) *ScopeUpdate {
	su.mutation.SetNonce(u)
	return su
}

// SetType sets the "type" field.
func (su *ScopeUpdate) SetType(s string) *ScopeUpdate {
	su.mutation.SetType(s)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *ScopeUpdate) SetNillableType(s *string) *ScopeUpdate {
	if s != nil {
		su.SetType(*s)
	}
	return su
}

// ClearType clears the value of the "type" field.
func (su *ScopeUpdate) ClearType() *ScopeUpdate {
	su.mutation.ClearType()
	return su
}

// SetExpiresAt sets the "expires_at" field.
func (su *ScopeUpdate) SetExpiresAt(t time.Time) *ScopeUpdate {
	su.mutation.SetExpiresAt(t)
	return su
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (su *ScopeUpdate) SetNillableExpiresAt(t *time.Time) *ScopeUpdate {
	if t != nil {
		su.SetExpiresAt(*t)
	}
	return su
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (su *ScopeUpdate) ClearExpiresAt() *ScopeUpdate {
	su.mutation.ClearExpiresAt()
	return su
}

// AddFactIDs adds the "facts" edge to the Fact entity by IDs.
func (su *ScopeUpdate) AddFactIDs(ids ...uuid.UUID) *ScopeUpdate {
	su.mutation.AddFactIDs(ids...)
	return su
}

// AddFacts adds the "facts" edges to the Fact entity.
func (su *ScopeUpdate) AddFacts(f ...*Fact) *ScopeUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddFactIDs(ids...)
}

// Mutation returns the ScopeMutation object of the builder.
func (su *ScopeUpdate) Mutation() *ScopeMutation {
	return su.mutation
}

// ClearFacts clears all "facts" edges to the Fact entity.
func (su *ScopeUpdate) ClearFacts() *ScopeUpdate {
	su.mutation.ClearFacts()
	return su
}

// RemoveFactIDs removes the "facts" edge to Fact entities by IDs.
func (su *ScopeUpdate) RemoveFactIDs(ids ...uuid.UUID) *ScopeUpdate {
	su.mutation.RemoveFactIDs(ids...)
	return su
}

// RemoveFacts removes "facts" edges to Fact entities.
func (su *ScopeUpdate) RemoveFacts(f ...*Fact) *ScopeUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveFactIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScopeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScopeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScopeUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScopeUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScopeUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ScopeUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := scope.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

func (su *ScopeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scope.Table,
			Columns: scope.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scope.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: scope.FieldNonce,
		})
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scope.FieldType,
		})
	}
	if su.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: scope.FieldType,
		})
	}
	if value, ok := su.mutation.ExpiresAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldExpiresAt,
		})
	}
	if su.mutation.ExpiresAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: scope.FieldExpiresAt,
		})
	}
	if su.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedFactsIDs(); len(nodes) > 0 && !su.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scope.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ScopeUpdateOne is the builder for updating a single Scope entity.
type ScopeUpdateOne struct {
	config
	hooks    []Hook
	mutation *ScopeMutation
}

// SetNonce sets the "nonce" field.
func (suo *ScopeUpdateOne) SetNonce(u uuid.UUID) *ScopeUpdateOne {
	suo.mutation.SetNonce(u)
	return suo
}

// SetType sets the "type" field.
func (suo *ScopeUpdateOne) SetType(s string) *ScopeUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *ScopeUpdateOne) SetNillableType(s *string) *ScopeUpdateOne {
	if s != nil {
		suo.SetType(*s)
	}
	return suo
}

// ClearType clears the value of the "type" field.
func (suo *ScopeUpdateOne) ClearType() *ScopeUpdateOne {
	suo.mutation.ClearType()
	return suo
}

// SetExpiresAt sets the "expires_at" field.
func (suo *ScopeUpdateOne) SetExpiresAt(t time.Time) *ScopeUpdateOne {
	suo.mutation.SetExpiresAt(t)
	return suo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (suo *ScopeUpdateOne) SetNillableExpiresAt(t *time.Time) *ScopeUpdateOne {
	if t != nil {
		suo.SetExpiresAt(*t)
	}
	return suo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (suo *ScopeUpdateOne) ClearExpiresAt() *ScopeUpdateOne {
	suo.mutation.ClearExpiresAt()
	return suo
}

// AddFactIDs adds the "facts" edge to the Fact entity by IDs.
func (suo *ScopeUpdateOne) AddFactIDs(ids ...uuid.UUID) *ScopeUpdateOne {
	suo.mutation.AddFactIDs(ids...)
	return suo
}

// AddFacts adds the "facts" edges to the Fact entity.
func (suo *ScopeUpdateOne) AddFacts(f ...*Fact) *ScopeUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddFactIDs(ids...)
}

// Mutation returns the ScopeMutation object of the builder.
func (suo *ScopeUpdateOne) Mutation() *ScopeMutation {
	return suo.mutation
}

// ClearFacts clears all "facts" edges to the Fact entity.
func (suo *ScopeUpdateOne) ClearFacts() *ScopeUpdateOne {
	suo.mutation.ClearFacts()
	return suo
}

// RemoveFactIDs removes the "facts" edge to Fact entities by IDs.
func (suo *ScopeUpdateOne) RemoveFactIDs(ids ...uuid.UUID) *ScopeUpdateOne {
	suo.mutation.RemoveFactIDs(ids...)
	return suo
}

// RemoveFacts removes "facts" edges to Fact entities.
func (suo *ScopeUpdateOne) RemoveFacts(f ...*Fact) *ScopeUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveFactIDs(ids...)
}

// Save executes the query and returns the updated Scope entity.
func (suo *ScopeUpdateOne) Save(ctx context.Context) (*Scope, error) {
	var (
		err  error
		node *Scope
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScopeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScopeUpdateOne) SaveX(ctx context.Context) *Scope {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScopeUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScopeUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ScopeUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := scope.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

func (suo *ScopeUpdateOne) sqlSave(ctx context.Context) (_node *Scope, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scope.Table,
			Columns: scope.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scope.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Scope.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: scope.FieldNonce,
		})
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scope.FieldType,
		})
	}
	if suo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: scope.FieldType,
		})
	}
	if value, ok := suo.mutation.ExpiresAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldExpiresAt,
		})
	}
	if suo.mutation.ExpiresAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: scope.FieldExpiresAt,
		})
	}
	if suo.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedFactsIDs(); len(nodes) > 0 && !suo.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Scope{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scope.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
