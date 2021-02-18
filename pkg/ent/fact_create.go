// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/scope"
)

// FactCreate is the builder for creating a Fact entity.
type FactCreate struct {
	config
	mutation *FactMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (fc *FactCreate) SetCreateTime(t time.Time) *FactCreate {
	fc.mutation.SetCreateTime(t)
	return fc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fc *FactCreate) SetNillableCreateTime(t *time.Time) *FactCreate {
	if t != nil {
		fc.SetCreateTime(*t)
	}
	return fc
}

// SetUpdateTime sets the "update_time" field.
func (fc *FactCreate) SetUpdateTime(t time.Time) *FactCreate {
	fc.mutation.SetUpdateTime(t)
	return fc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (fc *FactCreate) SetNillableUpdateTime(t *time.Time) *FactCreate {
	if t != nil {
		fc.SetUpdateTime(*t)
	}
	return fc
}

// SetEncryptedValue sets the "encrypted_value" field.
func (fc *FactCreate) SetEncryptedValue(s string) *FactCreate {
	fc.mutation.SetEncryptedValue(s)
	return fc
}

// SetID sets the "id" field.
func (fc *FactCreate) SetID(u uuid.UUID) *FactCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetScopeID sets the "scope" edge to the Scope entity by ID.
func (fc *FactCreate) SetScopeID(id uuid.UUID) *FactCreate {
	fc.mutation.SetScopeID(id)
	return fc
}

// SetNillableScopeID sets the "scope" edge to the Scope entity by ID if the given value is not nil.
func (fc *FactCreate) SetNillableScopeID(id *uuid.UUID) *FactCreate {
	if id != nil {
		fc = fc.SetScopeID(*id)
	}
	return fc
}

// SetScope sets the "scope" edge to the Scope entity.
func (fc *FactCreate) SetScope(s *Scope) *FactCreate {
	return fc.SetScopeID(s.ID)
}

// SetFactTypeID sets the "fact_type" edge to the FactType entity by ID.
func (fc *FactCreate) SetFactTypeID(id uuid.UUID) *FactCreate {
	fc.mutation.SetFactTypeID(id)
	return fc
}

// SetNillableFactTypeID sets the "fact_type" edge to the FactType entity by ID if the given value is not nil.
func (fc *FactCreate) SetNillableFactTypeID(id *uuid.UUID) *FactCreate {
	if id != nil {
		fc = fc.SetFactTypeID(*id)
	}
	return fc
}

// SetFactType sets the "fact_type" edge to the FactType entity.
func (fc *FactCreate) SetFactType(f *FactType) *FactCreate {
	return fc.SetFactTypeID(f.ID)
}

// Mutation returns the FactMutation object of the builder.
func (fc *FactCreate) Mutation() *FactMutation {
	return fc.mutation
}

// Save creates the Fact in the database.
func (fc *FactCreate) Save(ctx context.Context) (*Fact, error) {
	var (
		err  error
		node *Fact
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FactCreate) SaveX(ctx context.Context) *Fact {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FactCreate) defaults() {
	if _, ok := fc.mutation.CreateTime(); !ok {
		v := fact.DefaultCreateTime()
		fc.mutation.SetCreateTime(v)
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		v := fact.DefaultUpdateTime()
		fc.mutation.SetUpdateTime(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := fact.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FactCreate) check() error {
	if _, ok := fc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := fc.mutation.EncryptedValue(); !ok {
		return &ValidationError{Name: "encrypted_value", err: errors.New("ent: missing required field \"encrypted_value\"")}
	}
	return nil
}

func (fc *FactCreate) sqlSave(ctx context.Context) (*Fact, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (fc *FactCreate) createSpec() (*Fact, *sqlgraph.CreateSpec) {
	var (
		_node = &Fact{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fact.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := fc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fact.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := fc.mutation.EncryptedValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fact.FieldEncryptedValue,
		})
		_node.EncryptedValue = value
	}
	if nodes := fc.mutation.ScopeIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FactTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fact.FactTypeTable,
			Columns: []string{fact.FactTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: facttype.FieldID,
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

// FactCreateBulk is the builder for creating many Fact entities in bulk.
type FactCreateBulk struct {
	config
	builders []*FactCreate
}

// Save creates the Fact entities in the database.
func (fcb *FactCreateBulk) Save(ctx context.Context) ([]*Fact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fact, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FactMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FactCreateBulk) SaveX(ctx context.Context) []*Fact {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
