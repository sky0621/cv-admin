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
	"github.com/sky0621/cv-admin/src/ent/careerskill"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/skill"
)

// SkillQuery is the builder for querying Skill entities.
type SkillQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Skill
	withCareerSkills *CareerSkillQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkillQuery builder.
func (sq *SkillQuery) Where(ps ...predicate.Skill) *SkillQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SkillQuery) Limit(limit int) *SkillQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SkillQuery) Offset(offset int) *SkillQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SkillQuery) Unique(unique bool) *SkillQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SkillQuery) Order(o ...OrderFunc) *SkillQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCareerSkills chains the current query on the "careerSkills" edge.
func (sq *SkillQuery) QueryCareerSkills() *CareerSkillQuery {
	query := &CareerSkillQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skill.Table, skill.FieldID, selector),
			sqlgraph.To(careerskill.Table, careerskill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, skill.CareerSkillsTable, skill.CareerSkillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Skill entity from the query.
// Returns a *NotFoundError when no Skill was found.
func (sq *SkillQuery) First(ctx context.Context) (*Skill, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SkillQuery) FirstX(ctx context.Context) *Skill {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Skill ID from the query.
// Returns a *NotFoundError when no Skill ID was found.
func (sq *SkillQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SkillQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Skill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Skill entity is found.
// Returns a *NotFoundError when no Skill entities are found.
func (sq *SkillQuery) Only(ctx context.Context) (*Skill, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skill.Label}
	default:
		return nil, &NotSingularError{skill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SkillQuery) OnlyX(ctx context.Context) *Skill {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Skill ID in the query.
// Returns a *NotSingularError when more than one Skill ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SkillQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skill.Label}
	default:
		err = &NotSingularError{skill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SkillQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Skills.
func (sq *SkillQuery) All(ctx context.Context) ([]*Skill, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SkillQuery) AllX(ctx context.Context) []*Skill {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Skill IDs.
func (sq *SkillQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(skill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SkillQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SkillQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SkillQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SkillQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SkillQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SkillQuery) Clone() *SkillQuery {
	if sq == nil {
		return nil
	}
	return &SkillQuery{
		config:           sq.config,
		limit:            sq.limit,
		offset:           sq.offset,
		order:            append([]OrderFunc{}, sq.order...),
		predicates:       append([]predicate.Skill{}, sq.predicates...),
		withCareerSkills: sq.withCareerSkills.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithCareerSkills tells the query-builder to eager-load the nodes that are connected to
// the "careerSkills" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SkillQuery) WithCareerSkills(opts ...func(*CareerSkillQuery)) *SkillQuery {
	query := &CareerSkillQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withCareerSkills = query
	return sq
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
//	client.Skill.Query().
//		GroupBy(skill.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SkillQuery) GroupBy(field string, fields ...string) *SkillGroupBy {
	grbuild := &SkillGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = skill.Label
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
//	client.Skill.Query().
//		Select(skill.FieldCreateTime).
//		Scan(ctx, &v)
func (sq *SkillQuery) Select(fields ...string) *SkillSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SkillSelect{SkillQuery: sq}
	selbuild.label = skill.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SkillSelect configured with the given aggregations.
func (sq *SkillQuery) Aggregate(fns ...AggregateFunc) *SkillSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SkillQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !skill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SkillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Skill, error) {
	var (
		nodes       = []*Skill{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withCareerSkills != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Skill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Skill{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCareerSkills; query != nil {
		if err := sq.loadCareerSkills(ctx, query, nodes,
			func(n *Skill) { n.Edges.CareerSkills = []*CareerSkill{} },
			func(n *Skill, e *CareerSkill) { n.Edges.CareerSkills = append(n.Edges.CareerSkills, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SkillQuery) loadCareerSkills(ctx context.Context, query *CareerSkillQuery, nodes []*Skill, init func(*Skill), assign func(*Skill, *CareerSkill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Skill)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CareerSkill(func(s *sql.Selector) {
		s.Where(sql.InValues(skill.CareerSkillsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.skill_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "skill_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "skill_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SkillQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skill.Table,
			Columns: skill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skill.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skill.FieldID)
		for i := range fields {
			if fields[i] != skill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SkillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(skill.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = skill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkillGroupBy is the group-by builder for Skill entities.
type SkillGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SkillGroupBy) Aggregate(fns ...AggregateFunc) *SkillGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SkillGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SkillGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !skill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SkillGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SkillSelect is the builder for selecting fields of Skill entities.
type SkillSelect struct {
	*SkillQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SkillSelect) Aggregate(fns ...AggregateFunc) *SkillSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SkillSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SkillQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SkillSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(ss.sql))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}