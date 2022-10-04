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
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
)

// UserActivityQuery is the builder for querying UserActivity entities.
type UserActivityQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserActivity
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserActivityQuery builder.
func (uaq *UserActivityQuery) Where(ps ...predicate.UserActivity) *UserActivityQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit adds a limit step to the query.
func (uaq *UserActivityQuery) Limit(limit int) *UserActivityQuery {
	uaq.limit = &limit
	return uaq
}

// Offset adds an offset step to the query.
func (uaq *UserActivityQuery) Offset(offset int) *UserActivityQuery {
	uaq.offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UserActivityQuery) Unique(unique bool) *UserActivityQuery {
	uaq.unique = &unique
	return uaq
}

// Order adds an order step to the query.
func (uaq *UserActivityQuery) Order(o ...OrderFunc) *UserActivityQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryUser chains the current query on the "user" edge.
func (uaq *UserActivityQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: uaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useractivity.Table, useractivity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useractivity.UserTable, useractivity.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserActivity entity from the query.
// Returns a *NotFoundError when no UserActivity was found.
func (uaq *UserActivityQuery) First(ctx context.Context) (*UserActivity, error) {
	nodes, err := uaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useractivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserActivityQuery) FirstX(ctx context.Context) *UserActivity {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserActivity ID from the query.
// Returns a *NotFoundError when no UserActivity ID was found.
func (uaq *UserActivityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useractivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UserActivityQuery) FirstIDX(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserActivity entity is found.
// Returns a *NotFoundError when no UserActivity entities are found.
func (uaq *UserActivityQuery) Only(ctx context.Context) (*UserActivity, error) {
	nodes, err := uaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useractivity.Label}
	default:
		return nil, &NotSingularError{useractivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserActivityQuery) OnlyX(ctx context.Context) *UserActivity {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserActivity ID in the query.
// Returns a *NotSingularError when more than one UserActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UserActivityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useractivity.Label}
	default:
		err = &NotSingularError{useractivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UserActivityQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserActivities.
func (uaq *UserActivityQuery) All(ctx context.Context) ([]*UserActivity, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserActivityQuery) AllX(ctx context.Context) []*UserActivity {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserActivity IDs.
func (uaq *UserActivityQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uaq.Select(useractivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserActivityQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserActivityQuery) Count(ctx context.Context) (int, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserActivityQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserActivityQuery) Exist(ctx context.Context) (bool, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserActivityQuery) Clone() *UserActivityQuery {
	if uaq == nil {
		return nil
	}
	return &UserActivityQuery{
		config:     uaq.config,
		limit:      uaq.limit,
		offset:     uaq.offset,
		order:      append([]OrderFunc{}, uaq.order...),
		predicates: append([]predicate.UserActivity{}, uaq.predicates...),
		withUser:   uaq.withUser.Clone(),
		// clone intermediate query.
		sql:    uaq.sql.Clone(),
		path:   uaq.path,
		unique: uaq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserActivityQuery) WithUser(opts ...func(*UserQuery)) *UserActivityQuery {
	query := &UserQuery{config: uaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaq.withUser = query
	return uaq
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
//	client.UserActivity.Query().
//		GroupBy(useractivity.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uaq *UserActivityQuery) GroupBy(field string, fields ...string) *UserActivityGroupBy {
	grbuild := &UserActivityGroupBy{config: uaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uaq.sqlQuery(ctx), nil
	}
	grbuild.label = useractivity.Label
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
//	client.UserActivity.Query().
//		Select(useractivity.FieldCreateTime).
//		Scan(ctx, &v)
func (uaq *UserActivityQuery) Select(fields ...string) *UserActivitySelect {
	uaq.fields = append(uaq.fields, fields...)
	selbuild := &UserActivitySelect{UserActivityQuery: uaq}
	selbuild.label = useractivity.Label
	selbuild.flds, selbuild.scan = &uaq.fields, selbuild.Scan
	return selbuild
}

func (uaq *UserActivityQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uaq.fields {
		if !useractivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UserActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserActivity, error) {
	var (
		nodes       = []*UserActivity{}
		withFKs     = uaq.withFKs
		_spec       = uaq.querySpec()
		loadedTypes = [1]bool{
			uaq.withUser != nil,
		}
	)
	if uaq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, useractivity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserActivity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserActivity{config: uaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uaq.withUser; query != nil {
		if err := uaq.loadUser(ctx, query, nodes, nil,
			func(n *UserActivity, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uaq *UserActivityQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserActivity, init func(*UserActivity), assign func(*UserActivity, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserActivity)
	for i := range nodes {
		if nodes[i].user_id == nil {
			continue
		}
		fk := *nodes[i].user_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uaq *UserActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	_spec.Node.Columns = uaq.fields
	if len(uaq.fields) > 0 {
		_spec.Unique = uaq.unique != nil && *uaq.unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserActivityQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := uaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (uaq *UserActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useractivity.Table,
			Columns: useractivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useractivity.FieldID,
			},
		},
		From:   uaq.sql,
		Unique: true,
	}
	if unique := uaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useractivity.FieldID)
		for i := range fields {
			if fields[i] != useractivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(useractivity.Table)
	columns := uaq.fields
	if len(columns) == 0 {
		columns = useractivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.unique != nil && *uaq.unique {
		selector.Distinct()
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserActivityGroupBy is the group-by builder for UserActivity entities.
type UserActivityGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserActivityGroupBy) Aggregate(fns ...AggregateFunc) *UserActivityGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the group-by query and scans the result into the given value.
func (uagb *UserActivityGroupBy) Scan(ctx context.Context, v any) error {
	query, err := uagb.path(ctx)
	if err != nil {
		return err
	}
	uagb.sql = query
	return uagb.sqlScan(ctx, v)
}

func (uagb *UserActivityGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range uagb.fields {
		if !useractivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uagb *UserActivityGroupBy) sqlQuery() *sql.Selector {
	selector := uagb.sql.Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uagb.fields)+len(uagb.fns))
		for _, f := range uagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uagb.fields...)...)
}

// UserActivitySelect is the builder for selecting fields of UserActivity entities.
type UserActivitySelect struct {
	*UserActivityQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UserActivitySelect) Scan(ctx context.Context, v any) error {
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	uas.sql = uas.UserActivityQuery.sqlQuery(ctx)
	return uas.sqlScan(ctx, v)
}

func (uas *UserActivitySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := uas.sql.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
