// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careerskill"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerSkillQuery is the builder for querying CareerSkill entities.
type CareerSkillQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CareerSkill
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CareerSkillQuery builder.
func (csq *CareerSkillQuery) Where(ps ...predicate.CareerSkill) *CareerSkillQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit adds a limit step to the query.
func (csq *CareerSkillQuery) Limit(limit int) *CareerSkillQuery {
	csq.limit = &limit
	return csq
}

// Offset adds an offset step to the query.
func (csq *CareerSkillQuery) Offset(offset int) *CareerSkillQuery {
	csq.offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CareerSkillQuery) Unique(unique bool) *CareerSkillQuery {
	csq.unique = &unique
	return csq
}

// Order adds an order step to the query.
func (csq *CareerSkillQuery) Order(o ...OrderFunc) *CareerSkillQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// First returns the first CareerSkill entity from the query.
// Returns a *NotFoundError when no CareerSkill was found.
func (csq *CareerSkillQuery) First(ctx context.Context) (*CareerSkill, error) {
	nodes, err := csq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{careerskill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CareerSkillQuery) FirstX(ctx context.Context) *CareerSkill {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CareerSkill ID from the query.
// Returns a *NotFoundError when no CareerSkill ID was found.
func (csq *CareerSkillQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{careerskill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CareerSkillQuery) FirstIDX(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CareerSkill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CareerSkill entity is found.
// Returns a *NotFoundError when no CareerSkill entities are found.
func (csq *CareerSkillQuery) Only(ctx context.Context) (*CareerSkill, error) {
	nodes, err := csq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{careerskill.Label}
	default:
		return nil, &NotSingularError{careerskill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CareerSkillQuery) OnlyX(ctx context.Context) *CareerSkill {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CareerSkill ID in the query.
// Returns a *NotSingularError when more than one CareerSkill ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *CareerSkillQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{careerskill.Label}
	default:
		err = &NotSingularError{careerskill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CareerSkillQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CareerSkills.
func (csq *CareerSkillQuery) All(ctx context.Context) ([]*CareerSkill, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csq *CareerSkillQuery) AllX(ctx context.Context) []*CareerSkill {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CareerSkill IDs.
func (csq *CareerSkillQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := csq.Select(careerskill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CareerSkillQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CareerSkillQuery) Count(ctx context.Context) (int, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CareerSkillQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CareerSkillQuery) Exist(ctx context.Context) (bool, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CareerSkillQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CareerSkillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CareerSkillQuery) Clone() *CareerSkillQuery {
	if csq == nil {
		return nil
	}
	return &CareerSkillQuery{
		config:     csq.config,
		limit:      csq.limit,
		offset:     csq.offset,
		order:      append([]OrderFunc{}, csq.order...),
		predicates: append([]predicate.CareerSkill{}, csq.predicates...),
		// clone intermediate query.
		sql:    csq.sql.Clone(),
		path:   csq.path,
		unique: csq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CareerSkill.Query().
//		GroupBy(careerskill.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (csq *CareerSkillQuery) GroupBy(field string, fields ...string) *CareerSkillGroupBy {
	grbuild := &CareerSkillGroupBy{config: csq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(ctx), nil
	}
	grbuild.label = careerskill.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.CareerSkill.Query().
//		Select(careerskill.FieldCreateTime).
//		Scan(ctx, &v)
func (csq *CareerSkillQuery) Select(fields ...string) *CareerSkillSelect {
	csq.fields = append(csq.fields, fields...)
	selbuild := &CareerSkillSelect{CareerSkillQuery: csq}
	selbuild.label = careerskill.Label
	selbuild.flds, selbuild.scan = &csq.fields, selbuild.Scan
	return selbuild
}

func (csq *CareerSkillQuery) prepareQuery(ctx context.Context) error {
	for _, f := range csq.fields {
		if !careerskill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CareerSkillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CareerSkill, error) {
	var (
		nodes = []*CareerSkill{}
		_spec = csq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CareerSkill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CareerSkill{config: csq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (csq *CareerSkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Columns = csq.fields
	if len(csq.fields) > 0 {
		_spec.Unique = csq.unique != nil && *csq.unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CareerSkillQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (csq *CareerSkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskill.Table,
			Columns: careerskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskill.FieldID,
			},
		},
		From:   csq.sql,
		Unique: true,
	}
	if unique := csq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := csq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careerskill.FieldID)
		for i := range fields {
			if fields[i] != careerskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CareerSkillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(careerskill.Table)
	columns := csq.fields
	if len(columns) == 0 {
		columns = careerskill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.unique != nil && *csq.unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CareerSkillGroupBy is the group-by builder for CareerSkill entities.
type CareerSkillGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CareerSkillGroupBy) Aggregate(fns ...AggregateFunc) *CareerSkillGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the group-by query and scans the result into the given value.
func (csgb *CareerSkillGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csgb.path(ctx)
	if err != nil {
		return err
	}
	csgb.sql = query
	return csgb.sqlScan(ctx, v)
}

func (csgb *CareerSkillGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range csgb.fields {
		if !careerskill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := csgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csgb *CareerSkillGroupBy) sqlQuery() *sql.Selector {
	selector := csgb.sql.Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(csgb.fields)+len(csgb.fns))
		for _, f := range csgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(csgb.fields...)...)
}

// CareerSkillSelect is the builder for selecting fields of CareerSkill entities.
type CareerSkillSelect struct {
	*CareerSkillQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (css *CareerSkillSelect) Scan(ctx context.Context, v interface{}) error {
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	css.sql = css.CareerSkillQuery.sqlQuery(ctx)
	return css.sqlScan(ctx, v)
}

func (css *CareerSkillSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := css.sql.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
