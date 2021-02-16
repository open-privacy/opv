// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/open-privacy-vault/opv/ent/fact"
	"github.com/open-privacy-vault/opv/ent/facttype"
	"github.com/open-privacy-vault/opv/ent/predicate"
)

// FactTypeUpdate is the builder for updating FactType entities.
type FactTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FactTypeMutation
}

// Where adds a new predicate for the FactTypeUpdate builder.
func (ftu *FactTypeUpdate) Where(ps ...predicate.FactType) *FactTypeUpdate {
	ftu.mutation.predicates = append(ftu.mutation.predicates, ps...)
	return ftu
}

// AddFactIDs adds the "facts" edge to the Fact entity by IDs.
func (ftu *FactTypeUpdate) AddFactIDs(ids ...uuid.UUID) *FactTypeUpdate {
	ftu.mutation.AddFactIDs(ids...)
	return ftu
}

// AddFacts adds the "facts" edges to the Fact entity.
func (ftu *FactTypeUpdate) AddFacts(f ...*Fact) *FactTypeUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFactIDs(ids...)
}

// Mutation returns the FactTypeMutation object of the builder.
func (ftu *FactTypeUpdate) Mutation() *FactTypeMutation {
	return ftu.mutation
}

// ClearFacts clears all "facts" edges to the Fact entity.
func (ftu *FactTypeUpdate) ClearFacts() *FactTypeUpdate {
	ftu.mutation.ClearFacts()
	return ftu
}

// RemoveFactIDs removes the "facts" edge to Fact entities by IDs.
func (ftu *FactTypeUpdate) RemoveFactIDs(ids ...uuid.UUID) *FactTypeUpdate {
	ftu.mutation.RemoveFactIDs(ids...)
	return ftu
}

// RemoveFacts removes "facts" edges to Fact entities.
func (ftu *FactTypeUpdate) RemoveFacts(f ...*Fact) *FactTypeUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFactIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FactTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ftu.defaults()
	if len(ftu.hooks) == 0 {
		affected, err = ftu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftu.mutation = mutation
			affected, err = ftu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftu.hooks) - 1; i >= 0; i-- {
			mut = ftu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FactTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FactTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FactTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FactTypeUpdate) defaults() {
	if _, ok := ftu.mutation.UpdateTime(); !ok {
		v := facttype.UpdateDefaultUpdateTime()
		ftu.mutation.SetUpdateTime(v)
	}
}

func (ftu *FactTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facttype.Table,
			Columns: facttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: facttype.FieldID,
			},
		},
	}
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facttype.FieldUpdateTime,
		})
	}
	if ftu.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	if nodes := ftu.mutation.RemovedFactsIDs(); len(nodes) > 0 && !ftu.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	if nodes := ftu.mutation.FactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FactTypeUpdateOne is the builder for updating a single FactType entity.
type FactTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *FactTypeMutation
}

// AddFactIDs adds the "facts" edge to the Fact entity by IDs.
func (ftuo *FactTypeUpdateOne) AddFactIDs(ids ...uuid.UUID) *FactTypeUpdateOne {
	ftuo.mutation.AddFactIDs(ids...)
	return ftuo
}

// AddFacts adds the "facts" edges to the Fact entity.
func (ftuo *FactTypeUpdateOne) AddFacts(f ...*Fact) *FactTypeUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFactIDs(ids...)
}

// Mutation returns the FactTypeMutation object of the builder.
func (ftuo *FactTypeUpdateOne) Mutation() *FactTypeMutation {
	return ftuo.mutation
}

// ClearFacts clears all "facts" edges to the Fact entity.
func (ftuo *FactTypeUpdateOne) ClearFacts() *FactTypeUpdateOne {
	ftuo.mutation.ClearFacts()
	return ftuo
}

// RemoveFactIDs removes the "facts" edge to Fact entities by IDs.
func (ftuo *FactTypeUpdateOne) RemoveFactIDs(ids ...uuid.UUID) *FactTypeUpdateOne {
	ftuo.mutation.RemoveFactIDs(ids...)
	return ftuo
}

// RemoveFacts removes "facts" edges to Fact entities.
func (ftuo *FactTypeUpdateOne) RemoveFacts(f ...*Fact) *FactTypeUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFactIDs(ids...)
}

// Save executes the query and returns the updated FactType entity.
func (ftuo *FactTypeUpdateOne) Save(ctx context.Context) (*FactType, error) {
	var (
		err  error
		node *FactType
	)
	ftuo.defaults()
	if len(ftuo.hooks) == 0 {
		node, err = ftuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FactTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftuo.mutation = mutation
			node, err = ftuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftuo.hooks) - 1; i >= 0; i-- {
			mut = ftuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FactTypeUpdateOne) SaveX(ctx context.Context) *FactType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FactTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FactTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FactTypeUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdateTime(); !ok {
		v := facttype.UpdateDefaultUpdateTime()
		ftuo.mutation.SetUpdateTime(v)
	}
}

func (ftuo *FactTypeUpdateOne) sqlSave(ctx context.Context) (_node *FactType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facttype.Table,
			Columns: facttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: facttype.FieldID,
			},
		},
	}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FactType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facttype.FieldUpdateTime,
		})
	}
	if ftuo.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	if nodes := ftuo.mutation.RemovedFactsIDs(); len(nodes) > 0 && !ftuo.mutation.FactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	if nodes := ftuo.mutation.FactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facttype.FactsTable,
			Columns: []string{facttype.FactsColumn},
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
	_node = &FactType{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
