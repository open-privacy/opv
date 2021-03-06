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
	"github.com/open-privacy/opv/pkg/ent/apiaudit"
	"github.com/open-privacy/opv/pkg/ent/predicate"
)

// APIAuditQuery is the builder for querying APIAudit entities.
type APIAuditQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.APIAudit
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the APIAuditQuery builder.
func (aaq *APIAuditQuery) Where(ps ...predicate.APIAudit) *APIAuditQuery {
	aaq.predicates = append(aaq.predicates, ps...)
	return aaq
}

// Limit adds a limit step to the query.
func (aaq *APIAuditQuery) Limit(limit int) *APIAuditQuery {
	aaq.limit = &limit
	return aaq
}

// Offset adds an offset step to the query.
func (aaq *APIAuditQuery) Offset(offset int) *APIAuditQuery {
	aaq.offset = &offset
	return aaq
}

// Order adds an order step to the query.
func (aaq *APIAuditQuery) Order(o ...OrderFunc) *APIAuditQuery {
	aaq.order = append(aaq.order, o...)
	return aaq
}

// First returns the first APIAudit entity from the query.
// Returns a *NotFoundError when no APIAudit was found.
func (aaq *APIAuditQuery) First(ctx context.Context) (*APIAudit, error) {
	nodes, err := aaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apiaudit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aaq *APIAuditQuery) FirstX(ctx context.Context) *APIAudit {
	node, err := aaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first APIAudit ID from the query.
// Returns a *NotFoundError when no APIAudit ID was found.
func (aaq *APIAuditQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apiaudit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aaq *APIAuditQuery) FirstIDX(ctx context.Context) string {
	id, err := aaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single APIAudit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one APIAudit entity is not found.
// Returns a *NotFoundError when no APIAudit entities are found.
func (aaq *APIAuditQuery) Only(ctx context.Context) (*APIAudit, error) {
	nodes, err := aaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apiaudit.Label}
	default:
		return nil, &NotSingularError{apiaudit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aaq *APIAuditQuery) OnlyX(ctx context.Context) *APIAudit {
	node, err := aaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only APIAudit ID in the query.
// Returns a *NotSingularError when exactly one APIAudit ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aaq *APIAuditQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = &NotSingularError{apiaudit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aaq *APIAuditQuery) OnlyIDX(ctx context.Context) string {
	id, err := aaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of APIAudits.
func (aaq *APIAuditQuery) All(ctx context.Context) ([]*APIAudit, error) {
	if err := aaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aaq *APIAuditQuery) AllX(ctx context.Context) []*APIAudit {
	nodes, err := aaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of APIAudit IDs.
func (aaq *APIAuditQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := aaq.Select(apiaudit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aaq *APIAuditQuery) IDsX(ctx context.Context) []string {
	ids, err := aaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aaq *APIAuditQuery) Count(ctx context.Context) (int, error) {
	if err := aaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aaq *APIAuditQuery) CountX(ctx context.Context) int {
	count, err := aaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aaq *APIAuditQuery) Exist(ctx context.Context) (bool, error) {
	if err := aaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aaq *APIAuditQuery) ExistX(ctx context.Context) bool {
	exist, err := aaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the APIAuditQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aaq *APIAuditQuery) Clone() *APIAuditQuery {
	if aaq == nil {
		return nil
	}
	return &APIAuditQuery{
		config:     aaq.config,
		limit:      aaq.limit,
		offset:     aaq.offset,
		order:      append([]OrderFunc{}, aaq.order...),
		predicates: append([]predicate.APIAudit{}, aaq.predicates...),
		// clone intermediate query.
		sql:  aaq.sql.Clone(),
		path: aaq.path,
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
//	client.APIAudit.Query().
//		GroupBy(apiaudit.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aaq *APIAuditQuery) GroupBy(field string, fields ...string) *APIAuditGroupBy {
	group := &APIAuditGroupBy{config: aaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aaq.sqlQuery(ctx), nil
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
//	client.APIAudit.Query().
//		Select(apiaudit.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (aaq *APIAuditQuery) Select(field string, fields ...string) *APIAuditSelect {
	aaq.fields = append([]string{field}, fields...)
	return &APIAuditSelect{APIAuditQuery: aaq}
}

func (aaq *APIAuditQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aaq.fields {
		if !apiaudit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aaq.path != nil {
		prev, err := aaq.path(ctx)
		if err != nil {
			return err
		}
		aaq.sql = prev
	}
	return nil
}

func (aaq *APIAuditQuery) sqlAll(ctx context.Context) ([]*APIAudit, error) {
	var (
		nodes = []*APIAudit{}
		_spec = aaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &APIAudit{config: aaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aaq *APIAuditQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aaq.querySpec()
	return sqlgraph.CountNodes(ctx, aaq.driver, _spec)
}

func (aaq *APIAuditQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aaq *APIAuditQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apiaudit.Table,
			Columns: apiaudit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: apiaudit.FieldID,
			},
		},
		From:   aaq.sql,
		Unique: true,
	}
	if fields := aaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apiaudit.FieldID)
		for i := range fields {
			if fields[i] != apiaudit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, apiaudit.ValidColumn)
			}
		}
	}
	return _spec
}

func (aaq *APIAuditQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aaq.driver.Dialect())
	t1 := builder.Table(apiaudit.Table)
	selector := builder.Select(t1.Columns(apiaudit.Columns...)...).From(t1)
	if aaq.sql != nil {
		selector = aaq.sql
		selector.Select(selector.Columns(apiaudit.Columns...)...)
	}
	for _, p := range aaq.predicates {
		p(selector)
	}
	for _, p := range aaq.order {
		p(selector, apiaudit.ValidColumn)
	}
	if offset := aaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// APIAuditGroupBy is the group-by builder for APIAudit entities.
type APIAuditGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aagb *APIAuditGroupBy) Aggregate(fns ...AggregateFunc) *APIAuditGroupBy {
	aagb.fns = append(aagb.fns, fns...)
	return aagb
}

// Scan applies the group-by query and scans the result into the given value.
func (aagb *APIAuditGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aagb.path(ctx)
	if err != nil {
		return err
	}
	aagb.sql = query
	return aagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aagb *APIAuditGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aagb.fields) > 1 {
		return nil, errors.New("ent: APIAuditGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aagb *APIAuditGroupBy) StringsX(ctx context.Context) []string {
	v, err := aagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aagb *APIAuditGroupBy) StringX(ctx context.Context) string {
	v, err := aagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aagb.fields) > 1 {
		return nil, errors.New("ent: APIAuditGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aagb *APIAuditGroupBy) IntsX(ctx context.Context) []int {
	v, err := aagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aagb *APIAuditGroupBy) IntX(ctx context.Context) int {
	v, err := aagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aagb.fields) > 1 {
		return nil, errors.New("ent: APIAuditGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aagb *APIAuditGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aagb *APIAuditGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aagb.fields) > 1 {
		return nil, errors.New("ent: APIAuditGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aagb *APIAuditGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aagb *APIAuditGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aagb *APIAuditGroupBy) BoolX(ctx context.Context) bool {
	v, err := aagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aagb *APIAuditGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aagb.fields {
		if !apiaudit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aagb *APIAuditGroupBy) sqlQuery() *sql.Selector {
	selector := aagb.sql
	columns := make([]string, 0, len(aagb.fields)+len(aagb.fns))
	columns = append(columns, aagb.fields...)
	for _, fn := range aagb.fns {
		columns = append(columns, fn(selector, apiaudit.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(aagb.fields...)
}

// APIAuditSelect is the builder for selecting fields of APIAudit entities.
type APIAuditSelect struct {
	*APIAuditQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aas *APIAuditSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aas.prepareQuery(ctx); err != nil {
		return err
	}
	aas.sql = aas.APIAuditQuery.sqlQuery(ctx)
	return aas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aas *APIAuditSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aas.fields) > 1 {
		return nil, errors.New("ent: APIAuditSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aas *APIAuditSelect) StringsX(ctx context.Context) []string {
	v, err := aas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aas *APIAuditSelect) StringX(ctx context.Context) string {
	v, err := aas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aas.fields) > 1 {
		return nil, errors.New("ent: APIAuditSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aas *APIAuditSelect) IntsX(ctx context.Context) []int {
	v, err := aas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aas *APIAuditSelect) IntX(ctx context.Context) int {
	v, err := aas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aas.fields) > 1 {
		return nil, errors.New("ent: APIAuditSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aas *APIAuditSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aas *APIAuditSelect) Float64X(ctx context.Context) float64 {
	v, err := aas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aas.fields) > 1 {
		return nil, errors.New("ent: APIAuditSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aas *APIAuditSelect) BoolsX(ctx context.Context) []bool {
	v, err := aas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aas *APIAuditSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{apiaudit.Label}
	default:
		err = fmt.Errorf("ent: APIAuditSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aas *APIAuditSelect) BoolX(ctx context.Context) bool {
	v, err := aas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aas *APIAuditSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aas.sqlQuery().Query()
	if err := aas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aas *APIAuditSelect) sqlQuery() sql.Querier {
	selector := aas.sql
	selector.Select(selector.Columns(aas.fields...)...)
	return selector
}
