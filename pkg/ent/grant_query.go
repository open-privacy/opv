// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/roney492/opv/pkg/ent/grant"
	"github.com/roney492/opv/pkg/ent/predicate"
)

// GrantQuery is the builder for querying Grant entities.
type GrantQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Grant
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GrantQuery builder.
func (gq *GrantQuery) Where(ps ...predicate.Grant) *GrantQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GrantQuery) Limit(limit int) *GrantQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GrantQuery) Offset(offset int) *GrantQuery {
	gq.offset = &offset
	return gq
}

// Order adds an order step to the query.
func (gq *GrantQuery) Order(o ...OrderFunc) *GrantQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Grant entity from the query.
// Returns a *NotFoundError when no Grant was found.
func (gq *GrantQuery) First(ctx context.Context) (*Grant, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GrantQuery) FirstX(ctx context.Context) *Grant {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Grant ID from the query.
// Returns a *NotFoundError when no Grant ID was found.
func (gq *GrantQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GrantQuery) FirstIDX(ctx context.Context) string {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Grant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Grant entity is not found.
// Returns a *NotFoundError when no Grant entities are found.
func (gq *GrantQuery) Only(ctx context.Context) (*Grant, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grant.Label}
	default:
		return nil, &NotSingularError{grant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GrantQuery) OnlyX(ctx context.Context) *Grant {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Grant ID in the query.
// Returns a *NotSingularError when exactly one Grant ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gq *GrantQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = &NotSingularError{grant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GrantQuery) OnlyIDX(ctx context.Context) string {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Grants.
func (gq *GrantQuery) All(ctx context.Context) ([]*Grant, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GrantQuery) AllX(ctx context.Context) []*Grant {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Grant IDs.
func (gq *GrantQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := gq.Select(grant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GrantQuery) IDsX(ctx context.Context) []string {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GrantQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GrantQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GrantQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GrantQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GrantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GrantQuery) Clone() *GrantQuery {
	if gq == nil {
		return nil
	}
	return &GrantQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		order:      append([]OrderFunc{}, gq.order...),
		predicates: append([]predicate.Grant{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Grant.Query().
//		GroupBy(grant.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GrantQuery) GroupBy(field string, fields ...string) *GrantGroupBy {
	group := &GrantGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Grant.Query().
//		Select(grant.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (gq *GrantQuery) Select(field string, fields ...string) *GrantSelect {
	gq.fields = append([]string{field}, fields...)
	return &GrantSelect{GrantQuery: gq}
}

func (gq *GrantQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !grant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GrantQuery) sqlAll(ctx context.Context) ([]*Grant, error) {
	var (
		nodes = []*Grant{}
		_spec = gq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Grant{config: gq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gq *GrantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GrantQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gq *GrantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grant.Table,
			Columns: grant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: grant.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grant.FieldID)
		for i := range fields {
			if fields[i] != grant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, grant.ValidColumn)
			}
		}
	}
	return _spec
}

func (gq *GrantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(grant.Table)
	selector := builder.Select(t1.Columns(grant.Columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(grant.Columns...)...)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector, grant.ValidColumn)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GrantGroupBy is the group-by builder for Grant entities.
type GrantGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GrantGroupBy) Aggregate(fns ...AggregateFunc) *GrantGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GrantGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GrantGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GrantGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GrantGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GrantGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GrantGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GrantGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GrantGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GrantGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GrantGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GrantGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GrantGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GrantGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GrantGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GrantGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GrantGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ggb.fields {
		if !grant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GrantGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql
	columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
	columns = append(columns, ggb.fields...)
	for _, fn := range ggb.fns {
		columns = append(columns, fn(selector, grant.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ggb.fields...)
}

// GrantSelect is the builder for selecting fields of Grant entities.
type GrantSelect struct {
	*GrantQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GrantSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GrantQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GrantSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GrantSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GrantSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GrantSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GrantSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GrantSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GrantSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GrantSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GrantSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GrantSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GrantSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GrantSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gs *GrantSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grant.Label}
	default:
		err = fmt.Errorf("ent: GrantSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GrantSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GrantSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sqlQuery().Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gs *GrantSelect) sqlQuery() sql.Querier {
	selector := gs.sql
	selector.Select(selector.Columns(gs.fields...)...)
	return selector
}
