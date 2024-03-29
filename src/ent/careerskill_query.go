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
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/skill"
)

// CareerSkillQuery is the builder for querying CareerSkill entities.
type CareerSkillQuery struct {
	config
	ctx                  *QueryContext
	order                []careerskill.OrderOption
	inters               []Interceptor
	predicates           []predicate.CareerSkill
	withCareerSkillGroup *CareerSkillGroupQuery
	withSkill            *SkillQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CareerSkillQuery builder.
func (csq *CareerSkillQuery) Where(ps ...predicate.CareerSkill) *CareerSkillQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit the number of records to be returned by this query.
func (csq *CareerSkillQuery) Limit(limit int) *CareerSkillQuery {
	csq.ctx.Limit = &limit
	return csq
}

// Offset to start from.
func (csq *CareerSkillQuery) Offset(offset int) *CareerSkillQuery {
	csq.ctx.Offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CareerSkillQuery) Unique(unique bool) *CareerSkillQuery {
	csq.ctx.Unique = &unique
	return csq
}

// Order specifies how the records should be ordered.
func (csq *CareerSkillQuery) Order(o ...careerskill.OrderOption) *CareerSkillQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryCareerSkillGroup chains the current query on the "careerSkillGroup" edge.
func (csq *CareerSkillQuery) QueryCareerSkillGroup() *CareerSkillGroupQuery {
	query := (&CareerSkillGroupClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(careerskill.Table, careerskill.FieldID, selector),
			sqlgraph.To(careerskillgroup.Table, careerskillgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, careerskill.CareerSkillGroupTable, careerskill.CareerSkillGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySkill chains the current query on the "skill" edge.
func (csq *CareerSkillQuery) QuerySkill() *SkillQuery {
	query := (&SkillClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(careerskill.Table, careerskill.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, careerskill.SkillTable, careerskill.SkillColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CareerSkill entity from the query.
// Returns a *NotFoundError when no CareerSkill was found.
func (csq *CareerSkillQuery) First(ctx context.Context) (*CareerSkill, error) {
	nodes, err := csq.Limit(1).All(setContextOp(ctx, csq.ctx, "First"))
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
	if ids, err = csq.Limit(1).IDs(setContextOp(ctx, csq.ctx, "FirstID")); err != nil {
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
	nodes, err := csq.Limit(2).All(setContextOp(ctx, csq.ctx, "Only"))
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
	if ids, err = csq.Limit(2).IDs(setContextOp(ctx, csq.ctx, "OnlyID")); err != nil {
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
	ctx = setContextOp(ctx, csq.ctx, "All")
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CareerSkill, *CareerSkillQuery]()
	return withInterceptors[[]*CareerSkill](ctx, csq, qr, csq.inters)
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
func (csq *CareerSkillQuery) IDs(ctx context.Context) (ids []int, err error) {
	if csq.ctx.Unique == nil && csq.path != nil {
		csq.Unique(true)
	}
	ctx = setContextOp(ctx, csq.ctx, "IDs")
	if err = csq.Select(careerskill.FieldID).Scan(ctx, &ids); err != nil {
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
	ctx = setContextOp(ctx, csq.ctx, "Count")
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csq, querierCount[*CareerSkillQuery](), csq.inters)
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
	ctx = setContextOp(ctx, csq.ctx, "Exist")
	switch _, err := csq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
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
		config:               csq.config,
		ctx:                  csq.ctx.Clone(),
		order:                append([]careerskill.OrderOption{}, csq.order...),
		inters:               append([]Interceptor{}, csq.inters...),
		predicates:           append([]predicate.CareerSkill{}, csq.predicates...),
		withCareerSkillGroup: csq.withCareerSkillGroup.Clone(),
		withSkill:            csq.withSkill.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithCareerSkillGroup tells the query-builder to eager-load the nodes that are connected to
// the "careerSkillGroup" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CareerSkillQuery) WithCareerSkillGroup(opts ...func(*CareerSkillGroupQuery)) *CareerSkillQuery {
	query := (&CareerSkillGroupClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withCareerSkillGroup = query
	return csq
}

// WithSkill tells the query-builder to eager-load the nodes that are connected to
// the "skill" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CareerSkillQuery) WithSkill(opts ...func(*SkillQuery)) *CareerSkillQuery {
	query := (&SkillClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withSkill = query
	return csq
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
	csq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CareerSkillGroupBy{build: csq}
	grbuild.flds = &csq.ctx.Fields
	grbuild.label = careerskill.Label
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
//	client.CareerSkill.Query().
//		Select(careerskill.FieldCreateTime).
//		Scan(ctx, &v)
func (csq *CareerSkillQuery) Select(fields ...string) *CareerSkillSelect {
	csq.ctx.Fields = append(csq.ctx.Fields, fields...)
	sbuild := &CareerSkillSelect{CareerSkillQuery: csq}
	sbuild.label = careerskill.Label
	sbuild.flds, sbuild.scan = &csq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CareerSkillSelect configured with the given aggregations.
func (csq *CareerSkillQuery) Aggregate(fns ...AggregateFunc) *CareerSkillSelect {
	return csq.Select().Aggregate(fns...)
}

func (csq *CareerSkillQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csq); err != nil {
				return err
			}
		}
	}
	for _, f := range csq.ctx.Fields {
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
		nodes       = []*CareerSkill{}
		withFKs     = csq.withFKs
		_spec       = csq.querySpec()
		loadedTypes = [2]bool{
			csq.withCareerSkillGroup != nil,
			csq.withSkill != nil,
		}
	)
	if csq.withCareerSkillGroup != nil || csq.withSkill != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, careerskill.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CareerSkill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CareerSkill{config: csq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := csq.withCareerSkillGroup; query != nil {
		if err := csq.loadCareerSkillGroup(ctx, query, nodes, nil,
			func(n *CareerSkill, e *CareerSkillGroup) { n.Edges.CareerSkillGroup = e }); err != nil {
			return nil, err
		}
	}
	if query := csq.withSkill; query != nil {
		if err := csq.loadSkill(ctx, query, nodes, nil,
			func(n *CareerSkill, e *Skill) { n.Edges.Skill = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (csq *CareerSkillQuery) loadCareerSkillGroup(ctx context.Context, query *CareerSkillGroupQuery, nodes []*CareerSkill, init func(*CareerSkill), assign func(*CareerSkill, *CareerSkillGroup)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CareerSkill)
	for i := range nodes {
		if nodes[i].career_skill_group_id == nil {
			continue
		}
		fk := *nodes[i].career_skill_group_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(careerskillgroup.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_skill_group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (csq *CareerSkillQuery) loadSkill(ctx context.Context, query *SkillQuery, nodes []*CareerSkill, init func(*CareerSkill), assign func(*CareerSkill, *Skill)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CareerSkill)
	for i := range nodes {
		if nodes[i].skill_id == nil {
			continue
		}
		fk := *nodes[i].skill_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(skill.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "skill_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (csq *CareerSkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Columns = csq.ctx.Fields
	if len(csq.ctx.Fields) > 0 {
		_spec.Unique = csq.ctx.Unique != nil && *csq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CareerSkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(careerskill.Table, careerskill.Columns, sqlgraph.NewFieldSpec(careerskill.FieldID, field.TypeInt))
	_spec.From = csq.sql
	if unique := csq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csq.path != nil {
		_spec.Unique = true
	}
	if fields := csq.ctx.Fields; len(fields) > 0 {
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
	if limit := csq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.ctx.Offset; offset != nil {
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
	columns := csq.ctx.Fields
	if len(columns) == 0 {
		columns = careerskill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.ctx.Unique != nil && *csq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CareerSkillGroupBy is the group-by builder for CareerSkill entities.
type CareerSkillGroupBy struct {
	selector
	build *CareerSkillQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CareerSkillGroupBy) Aggregate(fns ...AggregateFunc) *CareerSkillGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the selector query and scans the result into the given value.
func (csgb *CareerSkillGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csgb.build.ctx, "GroupBy")
	if err := csgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CareerSkillQuery, *CareerSkillGroupBy](ctx, csgb.build, csgb, csgb.build.inters, v)
}

func (csgb *CareerSkillGroupBy) sqlScan(ctx context.Context, root *CareerSkillQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csgb.flds)+len(csgb.fns))
		for _, f := range *csgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CareerSkillSelect is the builder for selecting fields of CareerSkill entities.
type CareerSkillSelect struct {
	*CareerSkillQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (css *CareerSkillSelect) Aggregate(fns ...AggregateFunc) *CareerSkillSelect {
	css.fns = append(css.fns, fns...)
	return css
}

// Scan applies the selector query and scans the result into the given value.
func (css *CareerSkillSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, css.ctx, "Select")
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CareerSkillQuery, *CareerSkillSelect](ctx, css.CareerSkillQuery, css, css.inters, v)
}

func (css *CareerSkillSelect) sqlScan(ctx context.Context, root *CareerSkillQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(css.fns))
	for _, fn := range css.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*css.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
