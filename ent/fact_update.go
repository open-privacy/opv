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

// FactUpdate is the builder for updating Fact entities.
type FactUpdate struct {
	config
	hooks    []Hook
	mutation *FactMutation
}

// Where adds a new predicate for the FactUpdate builder.
func (fu *FactUpdate) Where(ps ...predicate.Fact) *FactUpdate {
	fu.mutation.predicates = append(fu.mutation.predicates, ps...)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FactUpdate) SetCreatedAt(t time.Time) *FactUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FactUpdate) SetNillableCreatedAt(t *time.Time) *FactUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FactUpdate) SetUpdatedAt(t time.Time) *FactUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FactUpdate) SetNillableUpdatedAt(t *time.Time) *FactUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// SetEncryptedValue sets the "encrypted_value" field.
func (fu *FactUpdate) SetEncryptedValue(b []byte) *FactUpdate {
	fu.mutation.SetEncryptedValue(b)
	return fu
}

// SetScopeID sets the "scope" edge to the Scope entity by ID.
func (fu *FactUpdate) SetScopeID(id uuid.UUID) *FactUpdate {
	fu.mutation.SetScopeID(id)
	return fu
}

// SetNillableScopeID sets the "scope" edge to the Scope entity by ID if the given value is not nil.
func (fu *FactUpdate) SetNillableScopeID(id *uuid.UUID) *FactUpdate {
	if id != nil {
		fu = fu.SetScopeID(*id)
	}
	return fu
}

// SetScope sets the "scope" edge to the Scope entity.
func (fu *FactUpdate) SetScope(s *Scope) *FactUpdate {
	return fu.SetScopeID(s.ID)
}

// Mutation returns the FactMutation object of the builder.
func (fu *FactUpdate) Mutation() *FactMutation {
	return fu.mutation
}

// ClearScope clears the "scope" edge to the Scope entity.
func (fu *FactUpdate) ClearScope() *FactUpdate {
	fu.mutation.ClearScope()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FactUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FactUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FactUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FactUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fact.Table,
			Columns: fact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fact.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.EncryptedValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: fact.FieldEncryptedValue,
		})
	}
	if fu.mutation.ScopeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.ScopeTable,
			Columns: []string{fact.ScopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ScopeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.ScopeTable,
			Columns: []string{fact.ScopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fact.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FactUpdateOne is the builder for updating a single Fact entity.
type FactUpdateOne struct {
	config
	hooks    []Hook
	mutation *FactMutation
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FactUpdateOne) SetCreatedAt(t time.Time) *FactUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FactUpdateOne) SetNillableCreatedAt(t *time.Time) *FactUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FactUpdateOne) SetUpdatedAt(t time.Time) *FactUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FactUpdateOne) SetNillableUpdatedAt(t *time.Time) *FactUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// SetEncryptedValue sets the "encrypted_value" field.
func (fuo *FactUpdateOne) SetEncryptedValue(b []byte) *FactUpdateOne {
	fuo.mutation.SetEncryptedValue(b)
	return fuo
}

// SetScopeID sets the "scope" edge to the Scope entity by ID.
func (fuo *FactUpdateOne) SetScopeID(id uuid.UUID) *FactUpdateOne {
	fuo.mutation.SetScopeID(id)
	return fuo
}

// SetNillableScopeID sets the "scope" edge to the Scope entity by ID if the given value is not nil.
func (fuo *FactUpdateOne) SetNillableScopeID(id *uuid.UUID) *FactUpdateOne {
	if id != nil {
		fuo = fuo.SetScopeID(*id)
	}
	return fuo
}

// SetScope sets the "scope" edge to the Scope entity.
func (fuo *FactUpdateOne) SetScope(s *Scope) *FactUpdateOne {
	return fuo.SetScopeID(s.ID)
}

// Mutation returns the FactMutation object of the builder.
func (fuo *FactUpdateOne) Mutation() *FactMutation {
	return fuo.mutation
}

// ClearScope clears the "scope" edge to the Scope entity.
func (fuo *FactUpdateOne) ClearScope() *FactUpdateOne {
	fuo.mutation.ClearScope()
	return fuo
}

// Save executes the query and returns the updated Fact entity.
func (fuo *FactUpdateOne) Save(ctx context.Context) (*Fact, error) {
	var (
		err  error
		node *Fact
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FactUpdateOne) SaveX(ctx context.Context) *Fact {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FactUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FactUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FactUpdateOne) sqlSave(ctx context.Context) (_node *Fact, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fact.Table,
			Columns: fact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fact.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Fact.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.EncryptedValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: fact.FieldEncryptedValue,
		})
	}
	if fuo.mutation.ScopeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.ScopeTable,
			Columns: []string{fact.ScopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ScopeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.ScopeTable,
			Columns: []string{fact.ScopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Fact{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fact.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
