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
	"github.com/kurojs/ipinfo/ent/history"
	"github.com/kurojs/ipinfo/ent/predicate"
)

// HistoryQuery is the builder for querying History entities.
type HistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.History
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HistoryQuery builder.
func (hq *HistoryQuery) Where(ps ...predicate.History) *HistoryQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HistoryQuery) Limit(limit int) *HistoryQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HistoryQuery) Offset(offset int) *HistoryQuery {
	hq.offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HistoryQuery) Unique(unique bool) *HistoryQuery {
	hq.unique = &unique
	return hq
}

// Order adds an order step to the query.
func (hq *HistoryQuery) Order(o ...OrderFunc) *HistoryQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// First returns the first History entity from the query.
// Returns a *NotFoundError when no History was found.
func (hq *HistoryQuery) First(ctx context.Context) (*History, error) {
	nodes, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{history.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HistoryQuery) FirstX(ctx context.Context) *History {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first History ID from the query.
// Returns a *NotFoundError when no History ID was found.
func (hq *HistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{history.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single History entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one History entity is not found.
// Returns a *NotFoundError when no History entities are found.
func (hq *HistoryQuery) Only(ctx context.Context) (*History, error) {
	nodes, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{history.Label}
	default:
		return nil, &NotSingularError{history.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HistoryQuery) OnlyX(ctx context.Context) *History {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only History ID in the query.
// Returns a *NotSingularError when exactly one History ID is not found.
// Returns a *NotFoundError when no entities are found.
func (hq *HistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = &NotSingularError{history.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Histories.
func (hq *HistoryQuery) All(ctx context.Context) ([]*History, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HistoryQuery) AllX(ctx context.Context) []*History {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of History IDs.
func (hq *HistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hq.Select(history.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HistoryQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HistoryQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HistoryQuery) Clone() *HistoryQuery {
	if hq == nil {
		return nil
	}
	return &HistoryQuery{
		config:     hq.config,
		limit:      hq.limit,
		offset:     hq.offset,
		order:      append([]OrderFunc{}, hq.order...),
		predicates: append([]predicate.History{}, hq.predicates...),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IP string `json:"ip,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.History.Query().
//		GroupBy(history.FieldIP).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hq *HistoryQuery) GroupBy(field string, fields ...string) *HistoryGroupBy {
	group := &HistoryGroupBy{config: hq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IP string `json:"ip,omitempty"`
//	}
//
//	client.History.Query().
//		Select(history.FieldIP).
//		Scan(ctx, &v)
//
func (hq *HistoryQuery) Select(fields ...string) *HistorySelect {
	hq.fields = append(hq.fields, fields...)
	return &HistorySelect{HistoryQuery: hq}
}

func (hq *HistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hq.fields {
		if !history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HistoryQuery) sqlAll(ctx context.Context) ([]*History, error) {
	var (
		nodes = []*History{}
		_spec = hq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &History{config: hq.config}
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
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hq *HistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (hq *HistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   history.Table,
			Columns: history.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: history.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if unique := hq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for i := range fields {
			if fields[i] != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(history.Table)
	columns := hq.fields
	if len(columns) == 0 {
		columns = history.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HistoryGroupBy is the group-by builder for History entities.
type HistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HistoryGroupBy) Aggregate(fns ...AggregateFunc) *HistoryGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hgb *HistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hgb *HistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hgb *HistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := hgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hgb *HistoryGroupBy) StringX(ctx context.Context) string {
	v, err := hgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hgb *HistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := hgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hgb *HistoryGroupBy) IntX(ctx context.Context) int {
	v, err := hgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hgb *HistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hgb *HistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hgb *HistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hgb *HistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := hgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hgb *HistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hgb.fields {
		if !history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HistoryGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql.Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
		for _, f := range hgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(hgb.fields...)...)
}

// HistorySelect is the builder for selecting fields of History entities.
type HistorySelect struct {
	*HistoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	hs.sql = hs.HistoryQuery.sqlQuery(ctx)
	return hs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hs *HistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := hs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hs *HistorySelect) StringsX(ctx context.Context) []string {
	v, err := hs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hs *HistorySelect) StringX(ctx context.Context) string {
	v, err := hs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hs *HistorySelect) IntsX(ctx context.Context) []int {
	v, err := hs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hs *HistorySelect) IntX(ctx context.Context) int {
	v, err := hs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hs *HistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := hs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hs *HistorySelect) Float64X(ctx context.Context) float64 {
	v, err := hs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hs *HistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := hs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = fmt.Errorf("ent: HistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hs *HistorySelect) BoolX(ctx context.Context) bool {
	v, err := hs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hs *HistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sql.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
