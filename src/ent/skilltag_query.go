// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/skilltag"
)

// SkillTagQuery is the builder for querying SkillTag entities.
type SkillTagQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SkillTag
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkillTagQuery builder.
func (stq *SkillTagQuery) Where(ps ...predicate.SkillTag) *SkillTagQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *SkillTagQuery) Limit(limit int) *SkillTagQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *SkillTagQuery) Offset(offset int) *SkillTagQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SkillTagQuery) Unique(unique bool) *SkillTagQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *SkillTagQuery) Order(o ...OrderFunc) *SkillTagQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first SkillTag entity from the query.
// Returns a *NotFoundError when no SkillTag was found.
func (stq *SkillTagQuery) First(ctx context.Context) (*SkillTag, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skilltag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SkillTagQuery) FirstX(ctx context.Context) *SkillTag {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SkillTag ID from the query.
// Returns a *NotFoundError when no SkillTag ID was found.
func (stq *SkillTagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skilltag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SkillTagQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SkillTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SkillTag entity is found.
// Returns a *NotFoundError when no SkillTag entities are found.
func (stq *SkillTagQuery) Only(ctx context.Context) (*SkillTag, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skilltag.Label}
	default:
		return nil, &NotSingularError{skilltag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SkillTagQuery) OnlyX(ctx context.Context) *SkillTag {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SkillTag ID in the query.
// Returns a *NotSingularError when more than one SkillTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SkillTagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skilltag.Label}
	default:
		err = &NotSingularError{skilltag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SkillTagQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SkillTags.
func (stq *SkillTagQuery) All(ctx context.Context) ([]*SkillTag, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *SkillTagQuery) AllX(ctx context.Context) []*SkillTag {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SkillTag IDs.
func (stq *SkillTagQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := stq.Select(skilltag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SkillTagQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SkillTagQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SkillTagQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SkillTagQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SkillTagQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkillTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SkillTagQuery) Clone() *SkillTagQuery {
	if stq == nil {
		return nil
	}
	return &SkillTagQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.SkillTag{}, stq.predicates...),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SkillTag.Query().
//		GroupBy(skilltag.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (stq *SkillTagQuery) GroupBy(field string, fields ...string) *SkillTagGroupBy {
	grbuild := &SkillTagGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = skilltag.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.SkillTag.Query().
//		Select(skilltag.FieldName).
//		Scan(ctx, &v)
func (stq *SkillTagQuery) Select(fields ...string) *SkillTagSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &SkillTagSelect{SkillTagQuery: stq}
	selbuild.label = skilltag.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SkillTagSelect configured with the given aggregations.
func (stq *SkillTagQuery) Aggregate(fns ...AggregateFunc) *SkillTagSelect {
	return stq.Select().Aggregate(fns...)
}

func (stq *SkillTagQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !skilltag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *SkillTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SkillTag, error) {
	var (
		nodes = []*SkillTag{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SkillTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SkillTag{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (stq *SkillTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SkillTagQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (stq *SkillTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skilltag.Table,
			Columns: skilltag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skilltag.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skilltag.FieldID)
		for i := range fields {
			if fields[i] != skilltag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SkillTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(skilltag.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = skilltag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkillTagGroupBy is the group-by builder for SkillTag entities.
type SkillTagGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SkillTagGroupBy) Aggregate(fns ...AggregateFunc) *SkillTagGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *SkillTagGroupBy) Scan(ctx context.Context, v any) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *SkillTagGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range stgb.fields {
		if !skilltag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *SkillTagGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// SkillTagSelect is the builder for selecting fields of SkillTag entities.
type SkillTagSelect struct {
	*SkillTagQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sts *SkillTagSelect) Aggregate(fns ...AggregateFunc) *SkillTagSelect {
	sts.fns = append(sts.fns, fns...)
	return sts
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SkillTagSelect) Scan(ctx context.Context, v any) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.SkillTagQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *SkillTagSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(sts.fns))
	for _, fn := range sts.fns {
		aggregation = append(aggregation, fn(sts.sql))
	}
	switch n := len(*sts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		sts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		sts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
