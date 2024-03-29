// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// UserCareerGroupQuery is the builder for querying UserCareerGroup entities.
type UserCareerGroupQuery struct {
	config
	ctx         *QueryContext
	order       []usercareergroup.OrderOption
	inters      []Interceptor
	predicates  []predicate.UserCareerGroup
	withUser    *UserQuery
	withCareers *UserCareerQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCareerGroupQuery builder.
func (ucgq *UserCareerGroupQuery) Where(ps ...predicate.UserCareerGroup) *UserCareerGroupQuery {
	ucgq.predicates = append(ucgq.predicates, ps...)
	return ucgq
}

// Limit the number of records to be returned by this query.
func (ucgq *UserCareerGroupQuery) Limit(limit int) *UserCareerGroupQuery {
	ucgq.ctx.Limit = &limit
	return ucgq
}

// Offset to start from.
func (ucgq *UserCareerGroupQuery) Offset(offset int) *UserCareerGroupQuery {
	ucgq.ctx.Offset = &offset
	return ucgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucgq *UserCareerGroupQuery) Unique(unique bool) *UserCareerGroupQuery {
	ucgq.ctx.Unique = &unique
	return ucgq
}

// Order specifies how the records should be ordered.
func (ucgq *UserCareerGroupQuery) Order(o ...usercareergroup.OrderOption) *UserCareerGroupQuery {
	ucgq.order = append(ucgq.order, o...)
	return ucgq
}

// QueryUser chains the current query on the "user" edge.
func (ucgq *UserCareerGroupQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ucgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareergroup.Table, usercareergroup.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usercareergroup.UserTable, usercareergroup.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCareers chains the current query on the "careers" edge.
func (ucgq *UserCareerGroupQuery) QueryCareers() *UserCareerQuery {
	query := (&UserCareerClient{config: ucgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareergroup.Table, usercareergroup.FieldID, selector),
			sqlgraph.To(usercareer.Table, usercareer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usercareergroup.CareersTable, usercareergroup.CareersColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCareerGroup entity from the query.
// Returns a *NotFoundError when no UserCareerGroup was found.
func (ucgq *UserCareerGroupQuery) First(ctx context.Context) (*UserCareerGroup, error) {
	nodes, err := ucgq.Limit(1).All(setContextOp(ctx, ucgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercareergroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) FirstX(ctx context.Context) *UserCareerGroup {
	node, err := ucgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCareerGroup ID from the query.
// Returns a *NotFoundError when no UserCareerGroup ID was found.
func (ucgq *UserCareerGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucgq.Limit(1).IDs(setContextOp(ctx, ucgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercareergroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := ucgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCareerGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCareerGroup entity is found.
// Returns a *NotFoundError when no UserCareerGroup entities are found.
func (ucgq *UserCareerGroupQuery) Only(ctx context.Context) (*UserCareerGroup, error) {
	nodes, err := ucgq.Limit(2).All(setContextOp(ctx, ucgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercareergroup.Label}
	default:
		return nil, &NotSingularError{usercareergroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) OnlyX(ctx context.Context) *UserCareerGroup {
	node, err := ucgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCareerGroup ID in the query.
// Returns a *NotSingularError when more than one UserCareerGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucgq *UserCareerGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucgq.Limit(2).IDs(setContextOp(ctx, ucgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercareergroup.Label}
	default:
		err = &NotSingularError{usercareergroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCareerGroups.
func (ucgq *UserCareerGroupQuery) All(ctx context.Context) ([]*UserCareerGroup, error) {
	ctx = setContextOp(ctx, ucgq.ctx, "All")
	if err := ucgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserCareerGroup, *UserCareerGroupQuery]()
	return withInterceptors[[]*UserCareerGroup](ctx, ucgq, qr, ucgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) AllX(ctx context.Context) []*UserCareerGroup {
	nodes, err := ucgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCareerGroup IDs.
func (ucgq *UserCareerGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ucgq.ctx.Unique == nil && ucgq.path != nil {
		ucgq.Unique(true)
	}
	ctx = setContextOp(ctx, ucgq.ctx, "IDs")
	if err = ucgq.Select(usercareergroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := ucgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucgq *UserCareerGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ucgq.ctx, "Count")
	if err := ucgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ucgq, querierCount[*UserCareerGroupQuery](), ucgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) CountX(ctx context.Context) int {
	count, err := ucgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucgq *UserCareerGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ucgq.ctx, "Exist")
	switch _, err := ucgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ucgq *UserCareerGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := ucgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCareerGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucgq *UserCareerGroupQuery) Clone() *UserCareerGroupQuery {
	if ucgq == nil {
		return nil
	}
	return &UserCareerGroupQuery{
		config:      ucgq.config,
		ctx:         ucgq.ctx.Clone(),
		order:       append([]usercareergroup.OrderOption{}, ucgq.order...),
		inters:      append([]Interceptor{}, ucgq.inters...),
		predicates:  append([]predicate.UserCareerGroup{}, ucgq.predicates...),
		withUser:    ucgq.withUser.Clone(),
		withCareers: ucgq.withCareers.Clone(),
		// clone intermediate query.
		sql:  ucgq.sql.Clone(),
		path: ucgq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ucgq *UserCareerGroupQuery) WithUser(opts ...func(*UserQuery)) *UserCareerGroupQuery {
	query := (&UserClient{config: ucgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucgq.withUser = query
	return ucgq
}

// WithCareers tells the query-builder to eager-load the nodes that are connected to
// the "careers" edge. The optional arguments are used to configure the query builder of the edge.
func (ucgq *UserCareerGroupQuery) WithCareers(opts ...func(*UserCareerQuery)) *UserCareerGroupQuery {
	query := (&UserCareerClient{config: ucgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucgq.withCareers = query
	return ucgq
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
//	client.UserCareerGroup.Query().
//		GroupBy(usercareergroup.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ucgq *UserCareerGroupQuery) GroupBy(field string, fields ...string) *UserCareerGroupGroupBy {
	ucgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserCareerGroupGroupBy{build: ucgq}
	grbuild.flds = &ucgq.ctx.Fields
	grbuild.label = usercareergroup.Label
	grbuild.scan = grbuild.Scan
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
//	client.UserCareerGroup.Query().
//		Select(usercareergroup.FieldCreateTime).
//		Scan(ctx, &v)
func (ucgq *UserCareerGroupQuery) Select(fields ...string) *UserCareerGroupSelect {
	ucgq.ctx.Fields = append(ucgq.ctx.Fields, fields...)
	sbuild := &UserCareerGroupSelect{UserCareerGroupQuery: ucgq}
	sbuild.label = usercareergroup.Label
	sbuild.flds, sbuild.scan = &ucgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserCareerGroupSelect configured with the given aggregations.
func (ucgq *UserCareerGroupQuery) Aggregate(fns ...AggregateFunc) *UserCareerGroupSelect {
	return ucgq.Select().Aggregate(fns...)
}

func (ucgq *UserCareerGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ucgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ucgq); err != nil {
				return err
			}
		}
	}
	for _, f := range ucgq.ctx.Fields {
		if !usercareergroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucgq.path != nil {
		prev, err := ucgq.path(ctx)
		if err != nil {
			return err
		}
		ucgq.sql = prev
	}
	return nil
}

func (ucgq *UserCareerGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCareerGroup, error) {
	var (
		nodes       = []*UserCareerGroup{}
		withFKs     = ucgq.withFKs
		_spec       = ucgq.querySpec()
		loadedTypes = [2]bool{
			ucgq.withUser != nil,
			ucgq.withCareers != nil,
		}
	)
	if ucgq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usercareergroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserCareerGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserCareerGroup{config: ucgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucgq.withUser; query != nil {
		if err := ucgq.loadUser(ctx, query, nodes, nil,
			func(n *UserCareerGroup, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ucgq.withCareers; query != nil {
		if err := ucgq.loadCareers(ctx, query, nodes,
			func(n *UserCareerGroup) { n.Edges.Careers = []*UserCareer{} },
			func(n *UserCareerGroup, e *UserCareer) { n.Edges.Careers = append(n.Edges.Careers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucgq *UserCareerGroupQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserCareerGroup, init func(*UserCareerGroup), assign func(*UserCareerGroup, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserCareerGroup)
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
	if len(ids) == 0 {
		return nil
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
func (ucgq *UserCareerGroupQuery) loadCareers(ctx context.Context, query *UserCareerQuery, nodes []*UserCareerGroup, init func(*UserCareerGroup), assign func(*UserCareerGroup, *UserCareer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserCareerGroup)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserCareer(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(usercareergroup.CareersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.career_group_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "career_group_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "career_group_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ucgq *UserCareerGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucgq.querySpec()
	_spec.Node.Columns = ucgq.ctx.Fields
	if len(ucgq.ctx.Fields) > 0 {
		_spec.Unique = ucgq.ctx.Unique != nil && *ucgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ucgq.driver, _spec)
}

func (ucgq *UserCareerGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usercareergroup.Table, usercareergroup.Columns, sqlgraph.NewFieldSpec(usercareergroup.FieldID, field.TypeInt))
	_spec.From = ucgq.sql
	if unique := ucgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ucgq.path != nil {
		_spec.Unique = true
	}
	if fields := ucgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercareergroup.FieldID)
		for i := range fields {
			if fields[i] != usercareergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucgq *UserCareerGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucgq.driver.Dialect())
	t1 := builder.Table(usercareergroup.Table)
	columns := ucgq.ctx.Fields
	if len(columns) == 0 {
		columns = usercareergroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucgq.sql != nil {
		selector = ucgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucgq.ctx.Unique != nil && *ucgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ucgq.predicates {
		p(selector)
	}
	for _, p := range ucgq.order {
		p(selector)
	}
	if offset := ucgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCareerGroupGroupBy is the group-by builder for UserCareerGroup entities.
type UserCareerGroupGroupBy struct {
	selector
	build *UserCareerGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucggb *UserCareerGroupGroupBy) Aggregate(fns ...AggregateFunc) *UserCareerGroupGroupBy {
	ucggb.fns = append(ucggb.fns, fns...)
	return ucggb
}

// Scan applies the selector query and scans the result into the given value.
func (ucggb *UserCareerGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucggb.build.ctx, "GroupBy")
	if err := ucggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCareerGroupQuery, *UserCareerGroupGroupBy](ctx, ucggb.build, ucggb, ucggb.build.inters, v)
}

func (ucggb *UserCareerGroupGroupBy) sqlScan(ctx context.Context, root *UserCareerGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ucggb.fns))
	for _, fn := range ucggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ucggb.flds)+len(ucggb.fns))
		for _, f := range *ucggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ucggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserCareerGroupSelect is the builder for selecting fields of UserCareerGroup entities.
type UserCareerGroupSelect struct {
	*UserCareerGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ucgs *UserCareerGroupSelect) Aggregate(fns ...AggregateFunc) *UserCareerGroupSelect {
	ucgs.fns = append(ucgs.fns, fns...)
	return ucgs
}

// Scan applies the selector query and scans the result into the given value.
func (ucgs *UserCareerGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucgs.ctx, "Select")
	if err := ucgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCareerGroupQuery, *UserCareerGroupSelect](ctx, ucgs.UserCareerGroupQuery, ucgs, ucgs.inters, v)
}

func (ucgs *UserCareerGroupSelect) sqlScan(ctx context.Context, root *UserCareerGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ucgs.fns))
	for _, fn := range ucgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ucgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
