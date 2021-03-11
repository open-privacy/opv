// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-privacy/opv/pkg/ent/apiaudit"
)

// APIAuditCreate is the builder for creating a APIAudit entity.
type APIAuditCreate struct {
	config
	mutation *APIAuditMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (aac *APIAuditCreate) SetCreatedAt(t time.Time) *APIAuditCreate {
	aac.mutation.SetCreatedAt(t)
	return aac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableCreatedAt(t *time.Time) *APIAuditCreate {
	if t != nil {
		aac.SetCreatedAt(*t)
	}
	return aac
}

// SetUpdatedAt sets the "updated_at" field.
func (aac *APIAuditCreate) SetUpdatedAt(t time.Time) *APIAuditCreate {
	aac.mutation.SetUpdatedAt(t)
	return aac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableUpdatedAt(t *time.Time) *APIAuditCreate {
	if t != nil {
		aac.SetUpdatedAt(*t)
	}
	return aac
}

// SetDeletedAt sets the "deleted_at" field.
func (aac *APIAuditCreate) SetDeletedAt(t time.Time) *APIAuditCreate {
	aac.mutation.SetDeletedAt(t)
	return aac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableDeletedAt(t *time.Time) *APIAuditCreate {
	if t != nil {
		aac.SetDeletedAt(*t)
	}
	return aac
}

// SetPlane sets the "plane" field.
func (aac *APIAuditCreate) SetPlane(s string) *APIAuditCreate {
	aac.mutation.SetPlane(s)
	return aac
}

// SetHashedGrantToken sets the "hashed_grant_token" field.
func (aac *APIAuditCreate) SetHashedGrantToken(s string) *APIAuditCreate {
	aac.mutation.SetHashedGrantToken(s)
	return aac
}

// SetNillableHashedGrantToken sets the "hashed_grant_token" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableHashedGrantToken(s *string) *APIAuditCreate {
	if s != nil {
		aac.SetHashedGrantToken(*s)
	}
	return aac
}

// SetDomain sets the "domain" field.
func (aac *APIAuditCreate) SetDomain(s string) *APIAuditCreate {
	aac.mutation.SetDomain(s)
	return aac
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableDomain(s *string) *APIAuditCreate {
	if s != nil {
		aac.SetDomain(*s)
	}
	return aac
}

// SetHTTPPath sets the "http_path" field.
func (aac *APIAuditCreate) SetHTTPPath(s string) *APIAuditCreate {
	aac.mutation.SetHTTPPath(s)
	return aac
}

// SetNillableHTTPPath sets the "http_path" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableHTTPPath(s *string) *APIAuditCreate {
	if s != nil {
		aac.SetHTTPPath(*s)
	}
	return aac
}

// SetHTTPMethod sets the "http_method" field.
func (aac *APIAuditCreate) SetHTTPMethod(s string) *APIAuditCreate {
	aac.mutation.SetHTTPMethod(s)
	return aac
}

// SetNillableHTTPMethod sets the "http_method" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableHTTPMethod(s *string) *APIAuditCreate {
	if s != nil {
		aac.SetHTTPMethod(*s)
	}
	return aac
}

// SetSentHTTPStatus sets the "sent_http_status" field.
func (aac *APIAuditCreate) SetSentHTTPStatus(i int) *APIAuditCreate {
	aac.mutation.SetSentHTTPStatus(i)
	return aac
}

// SetNillableSentHTTPStatus sets the "sent_http_status" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableSentHTTPStatus(i *int) *APIAuditCreate {
	if i != nil {
		aac.SetSentHTTPStatus(*i)
	}
	return aac
}

// SetID sets the "id" field.
func (aac *APIAuditCreate) SetID(s string) *APIAuditCreate {
	aac.mutation.SetID(s)
	return aac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aac *APIAuditCreate) SetNillableID(s *string) *APIAuditCreate {
	if s != nil {
		aac.SetID(*s)
	}
	return aac
}

// Mutation returns the APIAuditMutation object of the builder.
func (aac *APIAuditCreate) Mutation() *APIAuditMutation {
	return aac.mutation
}

// Save creates the APIAudit in the database.
func (aac *APIAuditCreate) Save(ctx context.Context) (*APIAudit, error) {
	var (
		err  error
		node *APIAudit
	)
	aac.defaults()
	if len(aac.hooks) == 0 {
		if err = aac.check(); err != nil {
			return nil, err
		}
		node, err = aac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*APIAuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aac.check(); err != nil {
				return nil, err
			}
			aac.mutation = mutation
			node, err = aac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aac.hooks) - 1; i >= 0; i-- {
			mut = aac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aac *APIAuditCreate) SaveX(ctx context.Context) *APIAudit {
	v, err := aac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (aac *APIAuditCreate) defaults() {
	if _, ok := aac.mutation.CreatedAt(); !ok {
		v := apiaudit.DefaultCreatedAt()
		aac.mutation.SetCreatedAt(v)
	}
	if _, ok := aac.mutation.UpdatedAt(); !ok {
		v := apiaudit.DefaultUpdatedAt()
		aac.mutation.SetUpdatedAt(v)
	}
	if _, ok := aac.mutation.ID(); !ok {
		v := apiaudit.DefaultID()
		aac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aac *APIAuditCreate) check() error {
	if _, ok := aac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := aac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := aac.mutation.Plane(); !ok {
		return &ValidationError{Name: "plane", err: errors.New("ent: missing required field \"plane\"")}
	}
	if v, ok := aac.mutation.ID(); ok {
		if err := apiaudit.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (aac *APIAuditCreate) sqlSave(ctx context.Context) (*APIAudit, error) {
	_node, _spec := aac.createSpec()
	if err := sqlgraph.CreateNode(ctx, aac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (aac *APIAuditCreate) createSpec() (*APIAudit, *sqlgraph.CreateSpec) {
	var (
		_node = &APIAudit{config: aac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: apiaudit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: apiaudit.FieldID,
			},
		}
	)
	if id, ok := aac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := aac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := aac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := aac.mutation.Plane(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldPlane,
		})
		_node.Plane = value
	}
	if value, ok := aac.mutation.HashedGrantToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHashedGrantToken,
		})
		_node.HashedGrantToken = value
	}
	if value, ok := aac.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldDomain,
		})
		_node.Domain = value
	}
	if value, ok := aac.mutation.HTTPPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPPath,
		})
		_node.HTTPPath = value
	}
	if value, ok := aac.mutation.HTTPMethod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPMethod,
		})
		_node.HTTPMethod = value
	}
	if value, ok := aac.mutation.SentHTTPStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apiaudit.FieldSentHTTPStatus,
		})
		_node.SentHTTPStatus = value
	}
	return _node, _spec
}

// APIAuditCreateBulk is the builder for creating many APIAudit entities in bulk.
type APIAuditCreateBulk struct {
	config
	builders []*APIAuditCreate
}

// Save creates the APIAudit entities in the database.
func (aacb *APIAuditCreateBulk) Save(ctx context.Context) ([]*APIAudit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aacb.builders))
	nodes := make([]*APIAudit, len(aacb.builders))
	mutators := make([]Mutator, len(aacb.builders))
	for i := range aacb.builders {
		func(i int, root context.Context) {
			builder := aacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*APIAuditMutation)
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
					_, err = mutators[i+1].Mutate(root, aacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aacb *APIAuditCreateBulk) SaveX(ctx context.Context) []*APIAudit {
	v, err := aacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
