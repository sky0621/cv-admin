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
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
)

// UserCareerDescriptionQuery is the builder for querying UserCareerDescription entities.
type UserCareerDescriptionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserCareerDescription
	withCareer *UserCareerQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCareerDescriptionQuery builder.
func (ucdq *UserCareerDescriptionQuery) Where(ps ...predicate.UserCareerDescription) *UserCareerDescriptionQuery {
	ucdq.predicates = append(ucdq.predicates, ps...)
	return ucdq
}

// Limit adds a limit step to the query.
func (ucdq *UserCareerDescriptionQuery) Limit(limit int) *UserCareerDescriptionQuery {
	ucdq.limit = &limit
	return ucdq
}

// Offset adds an offset step to the query.
func (ucdq *UserCareerDescriptionQuery) Offset(offset int) *UserCareerDescriptionQuery {
	ucdq.offset = &offset
	return ucdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucdq *UserCareerDescriptionQuery) Unique(unique bool) *UserCareerDescriptionQuery {
	ucdq.unique = &unique
	return ucdq
}

// Order adds an order step to the query.
func (ucdq *UserCareerDescriptionQuery) Order(o ...OrderFunc) *UserCareerDescriptionQuery {
	ucdq.order = append(ucdq.order, o...)
	return ucdq
}

// QueryCareer chains the current query on the "career" edge.
func (ucdq *UserCareerDescriptionQuery) QueryCareer() *UserCareerQuery {
	query := &UserCareerQuery{config: ucdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareerdescription.Table, usercareerdescription.FieldID, selector),
			sqlgraph.To(usercareer.Table, usercareer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usercareerdescription.CareerTable, usercareerdescription.CareerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCareerDescription entity from the query.
// Returns a *NotFoundError when no UserCareerDescription was found.
func (ucdq *UserCareerDescriptionQuery) First(ctx context.Context) (*UserCareerDescription, error) {
	nodes, err := ucdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercareerdescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) FirstX(ctx context.Context) *UserCareerDescription {
	node, err := ucdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCareerDescription ID from the query.
// Returns a *NotFoundError when no UserCareerDescription ID was found.
func (ucdq *UserCareerDescriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercareerdescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := ucdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCareerDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCareerDescription entity is found.
// Returns a *NotFoundError when no UserCareerDescription entities are found.
func (ucdq *UserCareerDescriptionQuery) Only(ctx context.Context) (*UserCareerDescription, error) {
	nodes, err := ucdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercareerdescription.Label}
	default:
		return nil, &NotSingularError{usercareerdescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) OnlyX(ctx context.Context) *UserCareerDescription {
	node, err := ucdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCareerDescription ID in the query.
// Returns a *NotSingularError when more than one UserCareerDescription ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucdq *UserCareerDescriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercareerdescription.Label}
	default:
		err = &NotSingularError{usercareerdescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCareerDescriptions.
func (ucdq *UserCareerDescriptionQuery) All(ctx context.Context) ([]*UserCareerDescription, error) {
	if err := ucdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) AllX(ctx context.Context) []*UserCareerDescription {
	nodes, err := ucdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCareerDescription IDs.
func (ucdq *UserCareerDescriptionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ucdq.Select(usercareerdescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := ucdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucdq *UserCareerDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := ucdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) CountX(ctx context.Context) int {
	count, err := ucdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucdq *UserCareerDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucdq *UserCareerDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := ucdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCareerDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucdq *UserCareerDescriptionQuery) Clone() *UserCareerDescriptionQuery {
	if ucdq == nil {
		return nil
	}
	return &UserCareerDescriptionQuery{
		config:     ucdq.config,
		limit:      ucdq.limit,
		offset:     ucdq.offset,
		order:      append([]OrderFunc{}, ucdq.order...),
		predicates: append([]predicate.UserCareerDescription{}, ucdq.predicates...),
		withCareer: ucdq.withCareer.Clone(),
		// clone intermediate query.
		sql:    ucdq.sql.Clone(),
		path:   ucdq.path,
		unique: ucdq.unique,
	}
}

// WithCareer tells the query-builder to eager-load the nodes that are connected to
// the "career" edge. The optional arguments are used to configure the query builder of the edge.
func (ucdq *UserCareerDescriptionQuery) WithCareer(opts ...func(*UserCareerQuery)) *UserCareerDescriptionQuery {
	query := &UserCareerQuery{config: ucdq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucdq.withCareer = query
	return ucdq
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
//	client.UserCareerDescription.Query().
//		GroupBy(usercareerdescription.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ucdq *UserCareerDescriptionQuery) GroupBy(field string, fields ...string) *UserCareerDescriptionGroupBy {
	grbuild := &UserCareerDescriptionGroupBy{config: ucdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucdq.sqlQuery(ctx), nil
	}
	grbuild.label = usercareerdescription.Label
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
//	client.UserCareerDescription.Query().
//		Select(usercareerdescription.FieldDescription).
//		Scan(ctx, &v)
func (ucdq *UserCareerDescriptionQuery) Select(fields ...string) *UserCareerDescriptionSelect {
	ucdq.fields = append(ucdq.fields, fields...)
	selbuild := &UserCareerDescriptionSelect{UserCareerDescriptionQuery: ucdq}
	selbuild.label = usercareerdescription.Label
	selbuild.flds, selbuild.scan = &ucdq.fields, selbuild.Scan
	return selbuild
}

func (ucdq *UserCareerDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucdq.fields {
		if !usercareerdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucdq.path != nil {
		prev, err := ucdq.path(ctx)
		if err != nil {
			return err
		}
		ucdq.sql = prev
	}
	return nil
}

func (ucdq *UserCareerDescriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCareerDescription, error) {
	var (
		nodes       = []*UserCareerDescription{}
		withFKs     = ucdq.withFKs
		_spec       = ucdq.querySpec()
		loadedTypes = [1]bool{
			ucdq.withCareer != nil,
		}
	)
	if ucdq.withCareer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usercareerdescription.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserCareerDescription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserCareerDescription{config: ucdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucdq.withCareer; query != nil {
		if err := ucdq.loadCareer(ctx, query, nodes, nil,
			func(n *UserCareerDescription, e *UserCareer) { n.Edges.Career = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucdq *UserCareerDescriptionQuery) loadCareer(ctx context.Context, query *UserCareerQuery, nodes []*UserCareerDescription, init func(*UserCareerDescription), assign func(*UserCareerDescription, *UserCareer)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserCareerDescription)
	for i := range nodes {
		if nodes[i].career_id == nil {
			continue
		}
		fk := *nodes[i].career_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(usercareer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ucdq *UserCareerDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucdq.querySpec()
	_spec.Node.Columns = ucdq.fields
	if len(ucdq.fields) > 0 {
		_spec.Unique = ucdq.unique != nil && *ucdq.unique
	}
	return sqlgraph.CountNodes(ctx, ucdq.driver, _spec)
}

func (ucdq *UserCareerDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucdq *UserCareerDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercareerdescription.Table,
			Columns: usercareerdescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usercareerdescription.FieldID,
			},
		},
		From:   ucdq.sql,
		Unique: true,
	}
	if unique := ucdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercareerdescription.FieldID)
		for i := range fields {
			if fields[i] != usercareerdescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucdq *UserCareerDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucdq.driver.Dialect())
	t1 := builder.Table(usercareerdescription.Table)
	columns := ucdq.fields
	if len(columns) == 0 {
		columns = usercareerdescription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucdq.sql != nil {
		selector = ucdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucdq.unique != nil && *ucdq.unique {
		selector.Distinct()
	}
	for _, p := range ucdq.predicates {
		p(selector)
	}
	for _, p := range ucdq.order {
		p(selector)
	}
	if offset := ucdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCareerDescriptionGroupBy is the group-by builder for UserCareerDescription entities.
type UserCareerDescriptionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucdgb *UserCareerDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *UserCareerDescriptionGroupBy {
	ucdgb.fns = append(ucdgb.fns, fns...)
	return ucdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucdgb *UserCareerDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucdgb.path(ctx)
	if err != nil {
		return err
	}
	ucdgb.sql = query
	return ucdgb.sqlScan(ctx, v)
}

func (ucdgb *UserCareerDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucdgb.fields {
		if !usercareerdescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucdgb *UserCareerDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := ucdgb.sql.Select()
	aggregation := make([]string, 0, len(ucdgb.fns))
	for _, fn := range ucdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucdgb.fields)+len(ucdgb.fns))
		for _, f := range ucdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucdgb.fields...)...)
}

// UserCareerDescriptionSelect is the builder for selecting fields of UserCareerDescription entities.
type UserCareerDescriptionSelect struct {
	*UserCareerDescriptionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucds *UserCareerDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucds.prepareQuery(ctx); err != nil {
		return err
	}
	ucds.sql = ucds.UserCareerDescriptionQuery.sqlQuery(ctx)
	return ucds.sqlScan(ctx, v)
}

func (ucds *UserCareerDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucds.sql.Query()
	if err := ucds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
