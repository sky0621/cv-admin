// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerSkillGroupQuery is the builder for querying CareerSkillGroup entities.
type CareerSkillGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CareerSkillGroup
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CareerSkillGroupQuery builder.
func (csgq *CareerSkillGroupQuery) Where(ps ...predicate.CareerSkillGroup) *CareerSkillGroupQuery {
	csgq.predicates = append(csgq.predicates, ps...)
	return csgq
}

// Limit adds a limit step to the query.
func (csgq *CareerSkillGroupQuery) Limit(limit int) *CareerSkillGroupQuery {
	csgq.limit = &limit
	return csgq
}

// Offset adds an offset step to the query.
func (csgq *CareerSkillGroupQuery) Offset(offset int) *CareerSkillGroupQuery {
	csgq.offset = &offset
	return csgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csgq *CareerSkillGroupQuery) Unique(unique bool) *CareerSkillGroupQuery {
	csgq.unique = &unique
	return csgq
}

// Order adds an order step to the query.
func (csgq *CareerSkillGroupQuery) Order(o ...OrderFunc) *CareerSkillGroupQuery {
	csgq.order = append(csgq.order, o...)
	return csgq
}

// First returns the first CareerSkillGroup entity from the query.
// Returns a *NotFoundError when no CareerSkillGroup was found.
func (csgq *CareerSkillGroupQuery) First(ctx context.Context) (*CareerSkillGroup, error) {
	nodes, err := csgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{careerskillgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) FirstX(ctx context.Context) *CareerSkillGroup {
	node, err := csgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CareerSkillGroup ID from the query.
// Returns a *NotFoundError when no CareerSkillGroup ID was found.
func (csgq *CareerSkillGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{careerskillgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := csgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CareerSkillGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CareerSkillGroup entity is found.
// Returns a *NotFoundError when no CareerSkillGroup entities are found.
func (csgq *CareerSkillGroupQuery) Only(ctx context.Context) (*CareerSkillGroup, error) {
	nodes, err := csgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{careerskillgroup.Label}
	default:
		return nil, &NotSingularError{careerskillgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) OnlyX(ctx context.Context) *CareerSkillGroup {
	node, err := csgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CareerSkillGroup ID in the query.
// Returns a *NotSingularError when more than one CareerSkillGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (csgq *CareerSkillGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{careerskillgroup.Label}
	default:
		err = &NotSingularError{careerskillgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := csgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CareerSkillGroups.
func (csgq *CareerSkillGroupQuery) All(ctx context.Context) ([]*CareerSkillGroup, error) {
	if err := csgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) AllX(ctx context.Context) []*CareerSkillGroup {
	nodes, err := csgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CareerSkillGroup IDs.
func (csgq *CareerSkillGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := csgq.Select(careerskillgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := csgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csgq *CareerSkillGroupQuery) Count(ctx context.Context) (int, error) {
	if err := csgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) CountX(ctx context.Context) int {
	count, err := csgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csgq *CareerSkillGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := csgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csgq *CareerSkillGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := csgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CareerSkillGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csgq *CareerSkillGroupQuery) Clone() *CareerSkillGroupQuery {
	if csgq == nil {
		return nil
	}
	return &CareerSkillGroupQuery{
		config:     csgq.config,
		limit:      csgq.limit,
		offset:     csgq.offset,
		order:      append([]OrderFunc{}, csgq.order...),
		predicates: append([]predicate.CareerSkillGroup{}, csgq.predicates...),
		// clone intermediate query.
		sql:    csgq.sql.Clone(),
		path:   csgq.path,
		unique: csgq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (csgq *CareerSkillGroupQuery) GroupBy(field string, fields ...string) *CareerSkillGroupGroupBy {
	grbuild := &CareerSkillGroupGroupBy{config: csgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csgq.sqlQuery(ctx), nil
	}
	grbuild.label = careerskillgroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (csgq *CareerSkillGroupQuery) Select(fields ...string) *CareerSkillGroupSelect {
	csgq.fields = append(csgq.fields, fields...)
	selbuild := &CareerSkillGroupSelect{CareerSkillGroupQuery: csgq}
	selbuild.label = careerskillgroup.Label
	selbuild.flds, selbuild.scan = &csgq.fields, selbuild.Scan
	return selbuild
}

func (csgq *CareerSkillGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range csgq.fields {
		if !careerskillgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csgq.path != nil {
		prev, err := csgq.path(ctx)
		if err != nil {
			return err
		}
		csgq.sql = prev
	}
	return nil
}

func (csgq *CareerSkillGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CareerSkillGroup, error) {
	var (
		nodes = []*CareerSkillGroup{}
		_spec = csgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CareerSkillGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CareerSkillGroup{config: csgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (csgq *CareerSkillGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csgq.querySpec()
	_spec.Node.Columns = csgq.fields
	if len(csgq.fields) > 0 {
		_spec.Unique = csgq.unique != nil && *csgq.unique
	}
	return sqlgraph.CountNodes(ctx, csgq.driver, _spec)
}

func (csgq *CareerSkillGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (csgq *CareerSkillGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskillgroup.Table,
			Columns: careerskillgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskillgroup.FieldID,
			},
		},
		From:   csgq.sql,
		Unique: true,
	}
	if unique := csgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := csgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careerskillgroup.FieldID)
		for i := range fields {
			if fields[i] != careerskillgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csgq *CareerSkillGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csgq.driver.Dialect())
	t1 := builder.Table(careerskillgroup.Table)
	columns := csgq.fields
	if len(columns) == 0 {
		columns = careerskillgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csgq.sql != nil {
		selector = csgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csgq.unique != nil && *csgq.unique {
		selector.Distinct()
	}
	for _, p := range csgq.predicates {
		p(selector)
	}
	for _, p := range csgq.order {
		p(selector)
	}
	if offset := csgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CareerSkillGroupGroupBy is the group-by builder for CareerSkillGroup entities.
type CareerSkillGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csggb *CareerSkillGroupGroupBy) Aggregate(fns ...AggregateFunc) *CareerSkillGroupGroupBy {
	csggb.fns = append(csggb.fns, fns...)
	return csggb
}

// Scan applies the group-by query and scans the result into the given value.
func (csggb *CareerSkillGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csggb.path(ctx)
	if err != nil {
		return err
	}
	csggb.sql = query
	return csggb.sqlScan(ctx, v)
}

func (csggb *CareerSkillGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range csggb.fields {
		if !careerskillgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := csggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csggb *CareerSkillGroupGroupBy) sqlQuery() *sql.Selector {
	selector := csggb.sql.Select()
	aggregation := make([]string, 0, len(csggb.fns))
	for _, fn := range csggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(csggb.fields)+len(csggb.fns))
		for _, f := range csggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(csggb.fields...)...)
}

// CareerSkillGroupSelect is the builder for selecting fields of CareerSkillGroup entities.
type CareerSkillGroupSelect struct {
	*CareerSkillGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (csgs *CareerSkillGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := csgs.prepareQuery(ctx); err != nil {
		return err
	}
	csgs.sql = csgs.CareerSkillGroupQuery.sqlQuery(ctx)
	return csgs.sqlScan(ctx, v)
}

func (csgs *CareerSkillGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := csgs.sql.Query()
	if err := csgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}