// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/roney492/opv/pkg/ent/apiaudit"
	"github.com/roney492/opv/pkg/ent/predicate"
)

// APIAuditUpdate is the builder for updating APIAudit entities.
type APIAuditUpdate struct {
	config
	hooks    []Hook
	mutation *APIAuditMutation
}

// Where adds a new predicate for the APIAuditUpdate builder.
func (aau *APIAuditUpdate) Where(ps ...predicate.APIAudit) *APIAuditUpdate {
	aau.mutation.predicates = append(aau.mutation.predicates, ps...)
	return aau
}

// SetDeletedAt sets the "deleted_at" field.
func (aau *APIAuditUpdate) SetDeletedAt(t time.Time) *APIAuditUpdate {
	aau.mutation.SetDeletedAt(t)
	return aau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableDeletedAt(t *time.Time) *APIAuditUpdate {
	if t != nil {
		aau.SetDeletedAt(*t)
	}
	return aau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (aau *APIAuditUpdate) ClearDeletedAt() *APIAuditUpdate {
	aau.mutation.ClearDeletedAt()
	return aau
}

// SetPlane sets the "plane" field.
func (aau *APIAuditUpdate) SetPlane(s string) *APIAuditUpdate {
	aau.mutation.SetPlane(s)
	return aau
}

// SetHashedGrantToken sets the "hashed_grant_token" field.
func (aau *APIAuditUpdate) SetHashedGrantToken(s string) *APIAuditUpdate {
	aau.mutation.SetHashedGrantToken(s)
	return aau
}

// SetNillableHashedGrantToken sets the "hashed_grant_token" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableHashedGrantToken(s *string) *APIAuditUpdate {
	if s != nil {
		aau.SetHashedGrantToken(*s)
	}
	return aau
}

// ClearHashedGrantToken clears the value of the "hashed_grant_token" field.
func (aau *APIAuditUpdate) ClearHashedGrantToken() *APIAuditUpdate {
	aau.mutation.ClearHashedGrantToken()
	return aau
}

// SetDomain sets the "domain" field.
func (aau *APIAuditUpdate) SetDomain(s string) *APIAuditUpdate {
	aau.mutation.SetDomain(s)
	return aau
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableDomain(s *string) *APIAuditUpdate {
	if s != nil {
		aau.SetDomain(*s)
	}
	return aau
}

// ClearDomain clears the value of the "domain" field.
func (aau *APIAuditUpdate) ClearDomain() *APIAuditUpdate {
	aau.mutation.ClearDomain()
	return aau
}

// SetHTTPPath sets the "http_path" field.
func (aau *APIAuditUpdate) SetHTTPPath(s string) *APIAuditUpdate {
	aau.mutation.SetHTTPPath(s)
	return aau
}

// SetNillableHTTPPath sets the "http_path" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableHTTPPath(s *string) *APIAuditUpdate {
	if s != nil {
		aau.SetHTTPPath(*s)
	}
	return aau
}

// ClearHTTPPath clears the value of the "http_path" field.
func (aau *APIAuditUpdate) ClearHTTPPath() *APIAuditUpdate {
	aau.mutation.ClearHTTPPath()
	return aau
}

// SetHTTPMethod sets the "http_method" field.
func (aau *APIAuditUpdate) SetHTTPMethod(s string) *APIAuditUpdate {
	aau.mutation.SetHTTPMethod(s)
	return aau
}

// SetNillableHTTPMethod sets the "http_method" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableHTTPMethod(s *string) *APIAuditUpdate {
	if s != nil {
		aau.SetHTTPMethod(*s)
	}
	return aau
}

// ClearHTTPMethod clears the value of the "http_method" field.
func (aau *APIAuditUpdate) ClearHTTPMethod() *APIAuditUpdate {
	aau.mutation.ClearHTTPMethod()
	return aau
}

// SetSentHTTPStatus sets the "sent_http_status" field.
func (aau *APIAuditUpdate) SetSentHTTPStatus(i int) *APIAuditUpdate {
	aau.mutation.ResetSentHTTPStatus()
	aau.mutation.SetSentHTTPStatus(i)
	return aau
}

// SetNillableSentHTTPStatus sets the "sent_http_status" field if the given value is not nil.
func (aau *APIAuditUpdate) SetNillableSentHTTPStatus(i *int) *APIAuditUpdate {
	if i != nil {
		aau.SetSentHTTPStatus(*i)
	}
	return aau
}

// AddSentHTTPStatus adds i to the "sent_http_status" field.
func (aau *APIAuditUpdate) AddSentHTTPStatus(i int) *APIAuditUpdate {
	aau.mutation.AddSentHTTPStatus(i)
	return aau
}

// ClearSentHTTPStatus clears the value of the "sent_http_status" field.
func (aau *APIAuditUpdate) ClearSentHTTPStatus() *APIAuditUpdate {
	aau.mutation.ClearSentHTTPStatus()
	return aau
}

// Mutation returns the APIAuditMutation object of the builder.
func (aau *APIAuditUpdate) Mutation() *APIAuditMutation {
	return aau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aau *APIAuditUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aau.defaults()
	if len(aau.hooks) == 0 {
		affected, err = aau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*APIAuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aau.mutation = mutation
			affected, err = aau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aau.hooks) - 1; i >= 0; i-- {
			mut = aau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aau *APIAuditUpdate) SaveX(ctx context.Context) int {
	affected, err := aau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aau *APIAuditUpdate) Exec(ctx context.Context) error {
	_, err := aau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aau *APIAuditUpdate) ExecX(ctx context.Context) {
	if err := aau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aau *APIAuditUpdate) defaults() {
	if _, ok := aau.mutation.UpdatedAt(); !ok {
		v := apiaudit.UpdateDefaultUpdatedAt()
		aau.mutation.SetUpdatedAt(v)
	}
}

func (aau *APIAuditUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apiaudit.Table,
			Columns: apiaudit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: apiaudit.FieldID,
			},
		},
	}
	if ps := aau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldUpdatedAt,
		})
	}
	if value, ok := aau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldDeletedAt,
		})
	}
	if aau.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: apiaudit.FieldDeletedAt,
		})
	}
	if value, ok := aau.mutation.Plane(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldPlane,
		})
	}
	if value, ok := aau.mutation.HashedGrantToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHashedGrantToken,
		})
	}
	if aau.mutation.HashedGrantTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHashedGrantToken,
		})
	}
	if value, ok := aau.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldDomain,
		})
	}
	if aau.mutation.DomainCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldDomain,
		})
	}
	if value, ok := aau.mutation.HTTPPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPPath,
		})
	}
	if aau.mutation.HTTPPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHTTPPath,
		})
	}
	if value, ok := aau.mutation.HTTPMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPMethod,
		})
	}
	if aau.mutation.HTTPMethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHTTPMethod,
		})
	}
	if value, ok := aau.mutation.SentHTTPStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	if value, ok := aau.mutation.AddedSentHTTPStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	if aau.mutation.SentHTTPStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apiaudit.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// APIAuditUpdateOne is the builder for updating a single APIAudit entity.
type APIAuditUpdateOne struct {
	config
	hooks    []Hook
	mutation *APIAuditMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (aauo *APIAuditUpdateOne) SetDeletedAt(t time.Time) *APIAuditUpdateOne {
	aauo.mutation.SetDeletedAt(t)
	return aauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableDeletedAt(t *time.Time) *APIAuditUpdateOne {
	if t != nil {
		aauo.SetDeletedAt(*t)
	}
	return aauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (aauo *APIAuditUpdateOne) ClearDeletedAt() *APIAuditUpdateOne {
	aauo.mutation.ClearDeletedAt()
	return aauo
}

// SetPlane sets the "plane" field.
func (aauo *APIAuditUpdateOne) SetPlane(s string) *APIAuditUpdateOne {
	aauo.mutation.SetPlane(s)
	return aauo
}

// SetHashedGrantToken sets the "hashed_grant_token" field.
func (aauo *APIAuditUpdateOne) SetHashedGrantToken(s string) *APIAuditUpdateOne {
	aauo.mutation.SetHashedGrantToken(s)
	return aauo
}

// SetNillableHashedGrantToken sets the "hashed_grant_token" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableHashedGrantToken(s *string) *APIAuditUpdateOne {
	if s != nil {
		aauo.SetHashedGrantToken(*s)
	}
	return aauo
}

// ClearHashedGrantToken clears the value of the "hashed_grant_token" field.
func (aauo *APIAuditUpdateOne) ClearHashedGrantToken() *APIAuditUpdateOne {
	aauo.mutation.ClearHashedGrantToken()
	return aauo
}

// SetDomain sets the "domain" field.
func (aauo *APIAuditUpdateOne) SetDomain(s string) *APIAuditUpdateOne {
	aauo.mutation.SetDomain(s)
	return aauo
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableDomain(s *string) *APIAuditUpdateOne {
	if s != nil {
		aauo.SetDomain(*s)
	}
	return aauo
}

// ClearDomain clears the value of the "domain" field.
func (aauo *APIAuditUpdateOne) ClearDomain() *APIAuditUpdateOne {
	aauo.mutation.ClearDomain()
	return aauo
}

// SetHTTPPath sets the "http_path" field.
func (aauo *APIAuditUpdateOne) SetHTTPPath(s string) *APIAuditUpdateOne {
	aauo.mutation.SetHTTPPath(s)
	return aauo
}

// SetNillableHTTPPath sets the "http_path" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableHTTPPath(s *string) *APIAuditUpdateOne {
	if s != nil {
		aauo.SetHTTPPath(*s)
	}
	return aauo
}

// ClearHTTPPath clears the value of the "http_path" field.
func (aauo *APIAuditUpdateOne) ClearHTTPPath() *APIAuditUpdateOne {
	aauo.mutation.ClearHTTPPath()
	return aauo
}

// SetHTTPMethod sets the "http_method" field.
func (aauo *APIAuditUpdateOne) SetHTTPMethod(s string) *APIAuditUpdateOne {
	aauo.mutation.SetHTTPMethod(s)
	return aauo
}

// SetNillableHTTPMethod sets the "http_method" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableHTTPMethod(s *string) *APIAuditUpdateOne {
	if s != nil {
		aauo.SetHTTPMethod(*s)
	}
	return aauo
}

// ClearHTTPMethod clears the value of the "http_method" field.
func (aauo *APIAuditUpdateOne) ClearHTTPMethod() *APIAuditUpdateOne {
	aauo.mutation.ClearHTTPMethod()
	return aauo
}

// SetSentHTTPStatus sets the "sent_http_status" field.
func (aauo *APIAuditUpdateOne) SetSentHTTPStatus(i int) *APIAuditUpdateOne {
	aauo.mutation.ResetSentHTTPStatus()
	aauo.mutation.SetSentHTTPStatus(i)
	return aauo
}

// SetNillableSentHTTPStatus sets the "sent_http_status" field if the given value is not nil.
func (aauo *APIAuditUpdateOne) SetNillableSentHTTPStatus(i *int) *APIAuditUpdateOne {
	if i != nil {
		aauo.SetSentHTTPStatus(*i)
	}
	return aauo
}

// AddSentHTTPStatus adds i to the "sent_http_status" field.
func (aauo *APIAuditUpdateOne) AddSentHTTPStatus(i int) *APIAuditUpdateOne {
	aauo.mutation.AddSentHTTPStatus(i)
	return aauo
}

// ClearSentHTTPStatus clears the value of the "sent_http_status" field.
func (aauo *APIAuditUpdateOne) ClearSentHTTPStatus() *APIAuditUpdateOne {
	aauo.mutation.ClearSentHTTPStatus()
	return aauo
}

// Mutation returns the APIAuditMutation object of the builder.
func (aauo *APIAuditUpdateOne) Mutation() *APIAuditMutation {
	return aauo.mutation
}

// Save executes the query and returns the updated APIAudit entity.
func (aauo *APIAuditUpdateOne) Save(ctx context.Context) (*APIAudit, error) {
	var (
		err  error
		node *APIAudit
	)
	aauo.defaults()
	if len(aauo.hooks) == 0 {
		node, err = aauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*APIAuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aauo.mutation = mutation
			node, err = aauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aauo.hooks) - 1; i >= 0; i-- {
			mut = aauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aauo *APIAuditUpdateOne) SaveX(ctx context.Context) *APIAudit {
	node, err := aauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aauo *APIAuditUpdateOne) Exec(ctx context.Context) error {
	_, err := aauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aauo *APIAuditUpdateOne) ExecX(ctx context.Context) {
	if err := aauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aauo *APIAuditUpdateOne) defaults() {
	if _, ok := aauo.mutation.UpdatedAt(); !ok {
		v := apiaudit.UpdateDefaultUpdatedAt()
		aauo.mutation.SetUpdatedAt(v)
	}
}

func (aauo *APIAuditUpdateOne) sqlSave(ctx context.Context) (_node *APIAudit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apiaudit.Table,
			Columns: apiaudit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: apiaudit.FieldID,
			},
		},
	}
	id, ok := aauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing APIAudit.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := aauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldUpdatedAt,
		})
	}
	if value, ok := aauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: apiaudit.FieldDeletedAt,
		})
	}
	if aauo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: apiaudit.FieldDeletedAt,
		})
	}
	if value, ok := aauo.mutation.Plane(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldPlane,
		})
	}
	if value, ok := aauo.mutation.HashedGrantToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHashedGrantToken,
		})
	}
	if aauo.mutation.HashedGrantTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHashedGrantToken,
		})
	}
	if value, ok := aauo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldDomain,
		})
	}
	if aauo.mutation.DomainCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldDomain,
		})
	}
	if value, ok := aauo.mutation.HTTPPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPPath,
		})
	}
	if aauo.mutation.HTTPPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHTTPPath,
		})
	}
	if value, ok := aauo.mutation.HTTPMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apiaudit.FieldHTTPMethod,
		})
	}
	if aauo.mutation.HTTPMethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: apiaudit.FieldHTTPMethod,
		})
	}
	if value, ok := aauo.mutation.SentHTTPStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	if value, ok := aauo.mutation.AddedSentHTTPStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	if aauo.mutation.SentHTTPStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: apiaudit.FieldSentHTTPStatus,
		})
	}
	_node = &APIAudit{config: aauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apiaudit.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
