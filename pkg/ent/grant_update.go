// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-privacy/opv/pkg/ent/grant"
	"github.com/open-privacy/opv/pkg/ent/predicate"
)

// GrantUpdate is the builder for updating Grant entities.
type GrantUpdate struct {
	config
	hooks    []Hook
	mutation *GrantMutation
}

// Where adds a new predicate for the GrantUpdate builder.
func (gu *GrantUpdate) Where(ps ...predicate.Grant) *GrantUpdate {
	gu.mutation.predicates = append(gu.mutation.predicates, ps...)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GrantUpdate) SetDeletedAt(t time.Time) *GrantUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GrantUpdate) SetNillableDeletedAt(t *time.Time) *GrantUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GrantUpdate) ClearDeletedAt() *GrantUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// SetHashedGrantToken sets the "hashed_grant_token" field.
func (gu *GrantUpdate) SetHashedGrantToken(s string) *GrantUpdate {
	gu.mutation.SetHashedGrantToken(s)
	return gu
}

// SetDomain sets the "domain" field.
func (gu *GrantUpdate) SetDomain(s string) *GrantUpdate {
	gu.mutation.SetDomain(s)
	return gu
}

// SetVersion sets the "version" field.
func (gu *GrantUpdate) SetVersion(s string) *GrantUpdate {
	gu.mutation.SetVersion(s)
	return gu
}

// SetAllowedHTTPMethods sets the "allowed_http_methods" field.
func (gu *GrantUpdate) SetAllowedHTTPMethods(s string) *GrantUpdate {
	gu.mutation.SetAllowedHTTPMethods(s)
	return gu
}

// SetPaths sets the "paths" field.
func (gu *GrantUpdate) SetPaths(s []string) *GrantUpdate {
	gu.mutation.SetPaths(s)
	return gu
}

// Mutation returns the GrantMutation object of the builder.
func (gu *GrantUpdate) Mutation() *GrantMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GrantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GrantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GrantUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GrantUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GrantUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GrantUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := grant.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

func (gu *GrantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grant.Table,
			Columns: grant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: grant.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grant.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grant.FieldDeletedAt,
		})
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: grant.FieldDeletedAt,
		})
	}
	if value, ok := gu.mutation.HashedGrantToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldHashedGrantToken,
		})
	}
	if value, ok := gu.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldDomain,
		})
	}
	if value, ok := gu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldVersion,
		})
	}
	if value, ok := gu.mutation.AllowedHTTPMethods(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldAllowedHTTPMethods,
		})
	}
	if value, ok := gu.mutation.Paths(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: grant.FieldPaths,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GrantUpdateOne is the builder for updating a single Grant entity.
type GrantUpdateOne struct {
	config
	hooks    []Hook
	mutation *GrantMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GrantUpdateOne) SetDeletedAt(t time.Time) *GrantUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GrantUpdateOne) SetNillableDeletedAt(t *time.Time) *GrantUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GrantUpdateOne) ClearDeletedAt() *GrantUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// SetHashedGrantToken sets the "hashed_grant_token" field.
func (guo *GrantUpdateOne) SetHashedGrantToken(s string) *GrantUpdateOne {
	guo.mutation.SetHashedGrantToken(s)
	return guo
}

// SetDomain sets the "domain" field.
func (guo *GrantUpdateOne) SetDomain(s string) *GrantUpdateOne {
	guo.mutation.SetDomain(s)
	return guo
}

// SetVersion sets the "version" field.
func (guo *GrantUpdateOne) SetVersion(s string) *GrantUpdateOne {
	guo.mutation.SetVersion(s)
	return guo
}

// SetAllowedHTTPMethods sets the "allowed_http_methods" field.
func (guo *GrantUpdateOne) SetAllowedHTTPMethods(s string) *GrantUpdateOne {
	guo.mutation.SetAllowedHTTPMethods(s)
	return guo
}

// SetPaths sets the "paths" field.
func (guo *GrantUpdateOne) SetPaths(s []string) *GrantUpdateOne {
	guo.mutation.SetPaths(s)
	return guo
}

// Mutation returns the GrantMutation object of the builder.
func (guo *GrantUpdateOne) Mutation() *GrantMutation {
	return guo.mutation
}

// Save executes the query and returns the updated Grant entity.
func (guo *GrantUpdateOne) Save(ctx context.Context) (*Grant, error) {
	var (
		err  error
		node *Grant
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GrantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GrantUpdateOne) SaveX(ctx context.Context) *Grant {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GrantUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GrantUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GrantUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := grant.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

func (guo *GrantUpdateOne) sqlSave(ctx context.Context) (_node *Grant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grant.Table,
			Columns: grant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: grant.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Grant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grant.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grant.FieldDeletedAt,
		})
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: grant.FieldDeletedAt,
		})
	}
	if value, ok := guo.mutation.HashedGrantToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldHashedGrantToken,
		})
	}
	if value, ok := guo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldDomain,
		})
	}
	if value, ok := guo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldVersion,
		})
	}
	if value, ok := guo.mutation.AllowedHTTPMethods(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grant.FieldAllowedHTTPMethods,
		})
	}
	if value, ok := guo.mutation.Paths(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: grant.FieldPaths,
		})
	}
	_node = &Grant{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
