// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/predicate"
)

// FactDelete is the builder for deleting a Fact entity.
type FactDelete struct {
	config
	hooks    []Hook
	mutation *FactMutation
}

// Where adds a new predicate to the FactDelete builder.
func (fd *FactDelete) Where(ps ...predicate.Fact) *FactDelete {
	fd.mutation.predicates = append(fd.mutation.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FactDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FactDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FactDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fact.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FactDeleteOne is the builder for deleting a single Fact entity.
type FactDeleteOne struct {
	fd *FactDelete
}

// Exec executes the deletion query.
func (fdo *FactDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fact.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FactDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}