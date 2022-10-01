// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/careertaskdescription"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerTaskDescriptionQuery is the builder for querying CareerTaskDescription entities.
type CareerTaskDescriptionQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.CareerTaskDescription
	withCareerTask *CareerTaskQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CareerTaskDescriptionQuery builder.
func (ctdq *CareerTaskDescriptionQuery) Where(ps ...predicate.CareerTaskDescription) *CareerTaskDescriptionQuery {
	ctdq.predicates = append(ctdq.predicates, ps...)
	return ctdq
}

// Limit adds a limit step to the query.
func (ctdq *CareerTaskDescriptionQuery) Limit(limit int) *CareerTaskDescriptionQuery {
	ctdq.limit = &limit
	return ctdq
}

// Offset adds an offset step to the query.
func (ctdq *CareerTaskDescriptionQuery) Offset(offset int) *CareerTaskDescriptionQuery {
	ctdq.offset = &offset
	return ctdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctdq *CareerTaskDescriptionQuery) Unique(unique bool) *CareerTaskDescriptionQuery {
	ctdq.unique = &unique
	return ctdq
}

// Order adds an order step to the query.
func (ctdq *CareerTaskDescriptionQuery) Order(o ...OrderFunc) *CareerTaskDescriptionQuery {
	ctdq.order = append(ctdq.order, o...)
	return ctdq
}

// QueryCareerTask chains the current query on the "careerTask" edge.
func (ctdq *CareerTaskDescriptionQuery) QueryCareerTask() *CareerTaskQuery {
	query := &CareerTaskQuery{config: ctdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(careertaskdescription.Table, careertaskdescription.FieldID, selector),
			sqlgraph.To(careertask.Table, careertask.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, careertaskdescription.CareerTaskTable, careertaskdescription.CareerTaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CareerTaskDescription entity from the query.
// Returns a *NotFoundError when no CareerTaskDescription was found.
func (ctdq *CareerTaskDescriptionQuery) First(ctx context.Context) (*CareerTaskDescription, error) {
	nodes, err := ctdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{careertaskdescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) FirstX(ctx context.Context) *CareerTaskDescription {
	node, err := ctdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CareerTaskDescription ID from the query.
// Returns a *NotFoundError when no CareerTaskDescription ID was found.
func (ctdq *CareerTaskDescriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{careertaskdescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := ctdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CareerTaskDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CareerTaskDescription entity is found.
// Returns a *NotFoundError when no CareerTaskDescription entities are found.
func (ctdq *CareerTaskDescriptionQuery) Only(ctx context.Context) (*CareerTaskDescription, error) {
	nodes, err := ctdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{careertaskdescription.Label}
	default:
		return nil, &NotSingularError{careertaskdescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) OnlyX(ctx context.Context) *CareerTaskDescription {
	node, err := ctdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CareerTaskDescription ID in the query.
// Returns a *NotSingularError when more than one CareerTaskDescription ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctdq *CareerTaskDescriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{careertaskdescription.Label}
	default:
		err = &NotSingularError{careertaskdescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CareerTaskDescriptions.
func (ctdq *CareerTaskDescriptionQuery) All(ctx context.Context) ([]*CareerTaskDescription, error) {
	if err := ctdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) AllX(ctx context.Context) []*CareerTaskDescription {
	nodes, err := ctdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CareerTaskDescription IDs.
func (ctdq *CareerTaskDescriptionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctdq.Select(careertaskdescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := ctdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctdq *CareerTaskDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := ctdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) CountX(ctx context.Context) int {
	count, err := ctdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctdq *CareerTaskDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctdq *CareerTaskDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := ctdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CareerTaskDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctdq *CareerTaskDescriptionQuery) Clone() *CareerTaskDescriptionQuery {
	if ctdq == nil {
		return nil
	}
	return &CareerTaskDescriptionQuery{
		config:         ctdq.config,
		limit:          ctdq.limit,
		offset:         ctdq.offset,
		order:          append([]OrderFunc{}, ctdq.order...),
		predicates:     append([]predicate.CareerTaskDescription{}, ctdq.predicates...),
		withCareerTask: ctdq.withCareerTask.Clone(),
		// clone intermediate query.
		sql:    ctdq.sql.Clone(),
		path:   ctdq.path,
		unique: ctdq.unique,
	}
}

// WithCareerTask tells the query-builder to eager-load the nodes that are connected to
// the "careerTask" edge. The optional arguments are used to configure the query builder of the edge.
func (ctdq *CareerTaskDescriptionQuery) WithCareerTask(opts ...func(*CareerTaskQuery)) *CareerTaskDescriptionQuery {
	query := &CareerTaskQuery{config: ctdq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctdq.withCareerTask = query
	return ctdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CareerTaskDescription.Query().
//		GroupBy(careertaskdescription.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctdq *CareerTaskDescriptionQuery) GroupBy(field string, fields ...string) *CareerTaskDescriptionGroupBy {
	grbuild := &CareerTaskDescriptionGroupBy{config: ctdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctdq.sqlQuery(ctx), nil
	}
	grbuild.label = careertaskdescription.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.CareerTaskDescription.Query().
//		Select(careertaskdescription.FieldDescription).
//		Scan(ctx, &v)
func (ctdq *CareerTaskDescriptionQuery) Select(fields ...string) *CareerTaskDescriptionSelect {
	ctdq.fields = append(ctdq.fields, fields...)
	selbuild := &CareerTaskDescriptionSelect{CareerTaskDescriptionQuery: ctdq}
	selbuild.label = careertaskdescription.Label
	selbuild.flds, selbuild.scan = &ctdq.fields, selbuild.Scan
	return selbuild
}

func (ctdq *CareerTaskDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctdq.fields {
		if !careertaskdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctdq.path != nil {
		prev, err := ctdq.path(ctx)
		if err != nil {
			return err
		}
		ctdq.sql = prev
	}
	return nil
}

func (ctdq *CareerTaskDescriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CareerTaskDescription, error) {
	var (
		nodes       = []*CareerTaskDescription{}
		withFKs     = ctdq.withFKs
		_spec       = ctdq.querySpec()
		loadedTypes = [1]bool{
			ctdq.withCareerTask != nil,
		}
	)
	if ctdq.withCareerTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, careertaskdescription.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CareerTaskDescription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CareerTaskDescription{config: ctdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctdq.withCareerTask; query != nil {
		if err := ctdq.loadCareerTask(ctx, query, nodes, nil,
			func(n *CareerTaskDescription, e *CareerTask) { n.Edges.CareerTask = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctdq *CareerTaskDescriptionQuery) loadCareerTask(ctx context.Context, query *CareerTaskQuery, nodes []*CareerTaskDescription, init func(*CareerTaskDescription), assign func(*CareerTaskDescription, *CareerTask)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CareerTaskDescription)
	for i := range nodes {
		if nodes[i].career_task_id == nil {
			continue
		}
		fk := *nodes[i].career_task_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(careertask.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_task_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ctdq *CareerTaskDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctdq.querySpec()
	_spec.Node.Columns = ctdq.fields
	if len(ctdq.fields) > 0 {
		_spec.Unique = ctdq.unique != nil && *ctdq.unique
	}
	return sqlgraph.CountNodes(ctx, ctdq.driver, _spec)
}

func (ctdq *CareerTaskDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ctdq *CareerTaskDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careertaskdescription.Table,
			Columns: careertaskdescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careertaskdescription.FieldID,
			},
		},
		From:   ctdq.sql,
		Unique: true,
	}
	if unique := ctdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careertaskdescription.FieldID)
		for i := range fields {
			if fields[i] != careertaskdescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctdq *CareerTaskDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctdq.driver.Dialect())
	t1 := builder.Table(careertaskdescription.Table)
	columns := ctdq.fields
	if len(columns) == 0 {
		columns = careertaskdescription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctdq.sql != nil {
		selector = ctdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctdq.unique != nil && *ctdq.unique {
		selector.Distinct()
	}
	for _, p := range ctdq.predicates {
		p(selector)
	}
	for _, p := range ctdq.order {
		p(selector)
	}
	if offset := ctdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CareerTaskDescriptionGroupBy is the group-by builder for CareerTaskDescription entities.
type CareerTaskDescriptionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctdgb *CareerTaskDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *CareerTaskDescriptionGroupBy {
	ctdgb.fns = append(ctdgb.fns, fns...)
	return ctdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctdgb *CareerTaskDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctdgb.path(ctx)
	if err != nil {
		return err
	}
	ctdgb.sql = query
	return ctdgb.sqlScan(ctx, v)
}

func (ctdgb *CareerTaskDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctdgb.fields {
		if !careertaskdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctdgb *CareerTaskDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := ctdgb.sql.Select()
	aggregation := make([]string, 0, len(ctdgb.fns))
	for _, fn := range ctdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctdgb.fields)+len(ctdgb.fns))
		for _, f := range ctdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctdgb.fields...)...)
}

// CareerTaskDescriptionSelect is the builder for selecting fields of CareerTaskDescription entities.
type CareerTaskDescriptionSelect struct {
	*CareerTaskDescriptionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ctds *CareerTaskDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ctds.prepareQuery(ctx); err != nil {
		return err
	}
	ctds.sql = ctds.CareerTaskDescriptionQuery.sqlQuery(ctx)
	return ctds.sqlScan(ctx, v)
}

func (ctds *CareerTaskDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ctds.sql.Query()
	if err := ctds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
