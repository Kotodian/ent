// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/entc/integration/migrate/entv2/zoo"
	"entgo.io/ent/schema/field"
)

// ZooQuery is the builder for querying Zoo entities.
type ZooQuery struct {
	config
	ctx        *QueryContext
	order      []zoo.OrderOption
	inters     []Interceptor
	predicates []predicate.Zoo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ZooQuery builder.
func (zq *ZooQuery) Where(ps ...predicate.Zoo) *ZooQuery {
	zq.predicates = append(zq.predicates, ps...)
	return zq
}

// Limit the number of records to be returned by this query.
func (zq *ZooQuery) Limit(limit int) *ZooQuery {
	zq.ctx.Limit = &limit
	return zq
}

// Offset to start from.
func (zq *ZooQuery) Offset(offset int) *ZooQuery {
	zq.ctx.Offset = &offset
	return zq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (zq *ZooQuery) Unique(unique bool) *ZooQuery {
	zq.ctx.Unique = &unique
	return zq
}

// Order specifies how the records should be ordered.
func (zq *ZooQuery) Order(o ...zoo.OrderOption) *ZooQuery {
	zq.order = append(zq.order, o...)
	return zq
}

// First returns the first Zoo entity from the query.
// Returns a *NotFoundError when no Zoo was found.
func (zq *ZooQuery) First(ctx context.Context) (*Zoo, error) {
	nodes, err := zq.Limit(1).All(setContextOp(ctx, zq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{zoo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (zq *ZooQuery) FirstX(ctx context.Context) *Zoo {
	node, err := zq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Zoo ID from the query.
// Returns a *NotFoundError when no Zoo ID was found.
func (zq *ZooQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zq.Limit(1).IDs(setContextOp(ctx, zq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{zoo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (zq *ZooQuery) FirstIDX(ctx context.Context) int {
	id, err := zq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Zoo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Zoo entity is found.
// Returns a *NotFoundError when no Zoo entities are found.
func (zq *ZooQuery) Only(ctx context.Context) (*Zoo, error) {
	nodes, err := zq.Limit(2).All(setContextOp(ctx, zq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{zoo.Label}
	default:
		return nil, &NotSingularError{zoo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (zq *ZooQuery) OnlyX(ctx context.Context) *Zoo {
	node, err := zq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Zoo ID in the query.
// Returns a *NotSingularError when more than one Zoo ID is found.
// Returns a *NotFoundError when no entities are found.
func (zq *ZooQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zq.Limit(2).IDs(setContextOp(ctx, zq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{zoo.Label}
	default:
		err = &NotSingularError{zoo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (zq *ZooQuery) OnlyIDX(ctx context.Context) int {
	id, err := zq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Zoos.
func (zq *ZooQuery) All(ctx context.Context) ([]*Zoo, error) {
	ctx = setContextOp(ctx, zq.ctx, "All")
	if err := zq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Zoo, *ZooQuery]()
	return withInterceptors[[]*Zoo](ctx, zq, qr, zq.inters)
}

// AllX is like All, but panics if an error occurs.
func (zq *ZooQuery) AllX(ctx context.Context) []*Zoo {
	nodes, err := zq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Zoo IDs.
func (zq *ZooQuery) IDs(ctx context.Context) (ids []int, err error) {
	if zq.ctx.Unique == nil && zq.path != nil {
		zq.Unique(true)
	}
	ctx = setContextOp(ctx, zq.ctx, "IDs")
	if err = zq.Select(zoo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (zq *ZooQuery) IDsX(ctx context.Context) []int {
	ids, err := zq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (zq *ZooQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, zq.ctx, "Count")
	if err := zq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, zq, querierCount[*ZooQuery](), zq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (zq *ZooQuery) CountX(ctx context.Context) int {
	count, err := zq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (zq *ZooQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, zq.ctx, "Exist")
	switch _, err := zq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entv2: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (zq *ZooQuery) ExistX(ctx context.Context) bool {
	exist, err := zq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ZooQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (zq *ZooQuery) Clone() *ZooQuery {
	if zq == nil {
		return nil
	}
	return &ZooQuery{
		config:     zq.config,
		ctx:        zq.ctx.Clone(),
		order:      append([]zoo.OrderOption{}, zq.order...),
		inters:     append([]Interceptor{}, zq.inters...),
		predicates: append([]predicate.Zoo{}, zq.predicates...),
		// clone intermediate query.
		sql:  zq.sql.Clone(),
		path: zq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (zq *ZooQuery) GroupBy(field string, fields ...string) *ZooGroupBy {
	zq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ZooGroupBy{build: zq}
	grbuild.flds = &zq.ctx.Fields
	grbuild.label = zoo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (zq *ZooQuery) Select(fields ...string) *ZooSelect {
	zq.ctx.Fields = append(zq.ctx.Fields, fields...)
	sbuild := &ZooSelect{ZooQuery: zq}
	sbuild.label = zoo.Label
	sbuild.flds, sbuild.scan = &zq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ZooSelect configured with the given aggregations.
func (zq *ZooQuery) Aggregate(fns ...AggregateFunc) *ZooSelect {
	return zq.Select().Aggregate(fns...)
}

func (zq *ZooQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range zq.inters {
		if inter == nil {
			return fmt.Errorf("entv2: uninitialized interceptor (forgotten import entv2/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, zq); err != nil {
				return err
			}
		}
	}
	for _, f := range zq.ctx.Fields {
		if !zoo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
		}
	}
	if zq.path != nil {
		prev, err := zq.path(ctx)
		if err != nil {
			return err
		}
		zq.sql = prev
	}
	return nil
}

func (zq *ZooQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Zoo, error) {
	var (
		nodes = []*Zoo{}
		_spec = zq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Zoo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Zoo{config: zq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, zq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (zq *ZooQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := zq.querySpec()
	_spec.Node.Columns = zq.ctx.Fields
	if len(zq.ctx.Fields) > 0 {
		_spec.Unique = zq.ctx.Unique != nil && *zq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, zq.driver, _spec)
}

func (zq *ZooQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(zoo.Table, zoo.Columns, sqlgraph.NewFieldSpec(zoo.FieldID, field.TypeInt))
	_spec.From = zq.sql
	if unique := zq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if zq.path != nil {
		_spec.Unique = true
	}
	if fields := zq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, zoo.FieldID)
		for i := range fields {
			if fields[i] != zoo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := zq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := zq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := zq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := zq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (zq *ZooQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(zq.driver.Dialect())
	t1 := builder.Table(zoo.Table)
	columns := zq.ctx.Fields
	if len(columns) == 0 {
		columns = zoo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if zq.sql != nil {
		selector = zq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if zq.ctx.Unique != nil && *zq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range zq.predicates {
		p(selector)
	}
	for _, p := range zq.order {
		p(selector)
	}
	if offset := zq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := zq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ZooGroupBy is the group-by builder for Zoo entities.
type ZooGroupBy struct {
	selector
	build *ZooQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (zgb *ZooGroupBy) Aggregate(fns ...AggregateFunc) *ZooGroupBy {
	zgb.fns = append(zgb.fns, fns...)
	return zgb
}

// Scan applies the selector query and scans the result into the given value.
func (zgb *ZooGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, zgb.build.ctx, "GroupBy")
	if err := zgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ZooQuery, *ZooGroupBy](ctx, zgb.build, zgb, zgb.build.inters, v)
}

func (zgb *ZooGroupBy) sqlScan(ctx context.Context, root *ZooQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(zgb.fns))
	for _, fn := range zgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*zgb.flds)+len(zgb.fns))
		for _, f := range *zgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*zgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := zgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ZooSelect is the builder for selecting fields of Zoo entities.
type ZooSelect struct {
	*ZooQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (zs *ZooSelect) Aggregate(fns ...AggregateFunc) *ZooSelect {
	zs.fns = append(zs.fns, fns...)
	return zs
}

// Scan applies the selector query and scans the result into the given value.
func (zs *ZooSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, zs.ctx, "Select")
	if err := zs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ZooQuery, *ZooSelect](ctx, zs.ZooQuery, zs, zs.inters, v)
}

func (zs *ZooSelect) sqlScan(ctx context.Context, root *ZooQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(zs.fns))
	for _, fn := range zs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*zs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := zs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
