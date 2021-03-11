// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/scope"
)

// ScopeCreate is the builder for creating a Scope entity.
type ScopeCreate struct {
	config
	mutation *ScopeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ScopeCreate) SetCreatedAt(t time.Time) *ScopeCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ScopeCreate) SetNillableCreatedAt(t *time.Time) *ScopeCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ScopeCreate) SetUpdatedAt(t time.Time) *ScopeCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ScopeCreate) SetNillableUpdatedAt(t *time.Time) *ScopeCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *ScopeCreate) SetDeletedAt(t time.Time) *ScopeCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *ScopeCreate) SetNillableDeletedAt(t *time.Time) *ScopeCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetCustomID sets the "custom_id" field.
func (sc *ScopeCreate) SetCustomID(s string) *ScopeCreate {
	sc.mutation.SetCustomID(s)
	return sc
}

// SetNonce sets the "nonce" field.
func (sc *ScopeCreate) SetNonce(s string) *ScopeCreate {
	sc.mutation.SetNonce(s)
	return sc
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (sc *ScopeCreate) SetNillableNonce(s *string) *ScopeCreate {
	if s != nil {
		sc.SetNonce(*s)
	}
	return sc
}

// SetDomain sets the "domain" field.
func (sc *ScopeCreate) SetDomain(s string) *ScopeCreate {
	sc.mutation.SetDomain(s)
	return sc
}

// SetID sets the "id" field.
func (sc *ScopeCreate) SetID(s string) *ScopeCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ScopeCreate) SetNillableID(s *string) *ScopeCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// AddFactIDs adds the "facts" edge to the Fact entity by IDs.
func (sc *ScopeCreate) AddFactIDs(ids ...string) *ScopeCreate {
	sc.mutation.AddFactIDs(ids...)
	return sc
}

// AddFacts adds the "facts" edges to the Fact entity.
func (sc *ScopeCreate) AddFacts(f ...*Fact) *ScopeCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return sc.AddFactIDs(ids...)
}

// Mutation returns the ScopeMutation object of the builder.
func (sc *ScopeCreate) Mutation() *ScopeMutation {
	return sc.mutation
}

// Save creates the Scope in the database.
func (sc *ScopeCreate) Save(ctx context.Context) (*Scope, error) {
	var (
		err  error
		node *Scope
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScopeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScopeCreate) SaveX(ctx context.Context) *Scope {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sc *ScopeCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := scope.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := scope.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Nonce(); !ok {
		v := scope.DefaultNonce()
		sc.mutation.SetNonce(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := scope.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScopeCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := sc.mutation.CustomID(); !ok {
		return &ValidationError{Name: "custom_id", err: errors.New("ent: missing required field \"custom_id\"")}
	}
	if _, ok := sc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New("ent: missing required field \"nonce\"")}
	}
	if _, ok := sc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New("ent: missing required field \"domain\"")}
	}
	if v, ok := sc.mutation.ID(); ok {
		if err := scope.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (sc *ScopeCreate) sqlSave(ctx context.Context) (*Scope, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sc *ScopeCreate) createSpec() (*Scope, *sqlgraph.CreateSpec) {
	var (
		_node = &Scope{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: scope.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: scope.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: scope.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sc.mutation.CustomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scope.FieldCustomID,
		})
		_node.CustomID = value
	}
	if value, ok := sc.mutation.Nonce(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scope.FieldNonce,
		})
		_node.Nonce = value
	}
	if value, ok := sc.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scope.FieldDomain,
		})
		_node.Domain = value
	}
	if nodes := sc.mutation.FactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scope.FactsTable,
			Columns: []string{scope.FactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: fact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScopeCreateBulk is the builder for creating many Scope entities in bulk.
type ScopeCreateBulk struct {
	config
	builders []*ScopeCreate
}

// Save creates the Scope entities in the database.
func (scb *ScopeCreateBulk) Save(ctx context.Context) ([]*Scope, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Scope, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScopeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScopeCreateBulk) SaveX(ctx context.Context) []*Scope {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
