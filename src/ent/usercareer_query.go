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
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// UserCareerQuery is the builder for querying UserCareer entities.
type UserCareerQuery struct {
	config
	limit                  *int
	offset                 *int
	unique                 *bool
	order                  []OrderFunc
	fields                 []string
	predicates             []predicate.UserCareer
	withCareerGroup        *UserCareerGroupQuery
	withCareerDescriptions *UserCareerDescriptionQuery
	withCareerTasks        *CareerTaskQuery
	withCareerSkillGroups  *CareerSkillGroupQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCareerQuery builder.
func (ucq *UserCareerQuery) Where(ps ...predicate.UserCareer) *UserCareerQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit adds a limit step to the query.
func (ucq *UserCareerQuery) Limit(limit int) *UserCareerQuery {
	ucq.limit = &limit
	return ucq
}

// Offset adds an offset step to the query.
func (ucq *UserCareerQuery) Offset(offset int) *UserCareerQuery {
	ucq.offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserCareerQuery) Unique(unique bool) *UserCareerQuery {
	ucq.unique = &unique
	return ucq
}

// Order adds an order step to the query.
func (ucq *UserCareerQuery) Order(o ...OrderFunc) *UserCareerQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryCareerGroup chains the current query on the "careerGroup" edge.
func (ucq *UserCareerQuery) QueryCareerGroup() *UserCareerGroupQuery {
	query := &UserCareerGroupQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareer.Table, usercareer.FieldID, selector),
			sqlgraph.To(usercareergroup.Table, usercareergroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usercareer.CareerGroupTable, usercareer.CareerGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCareerDescriptions chains the current query on the "careerDescriptions" edge.
func (ucq *UserCareerQuery) QueryCareerDescriptions() *UserCareerDescriptionQuery {
	query := &UserCareerDescriptionQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareer.Table, usercareer.FieldID, selector),
			sqlgraph.To(usercareerdescription.Table, usercareerdescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usercareer.CareerDescriptionsTable, usercareer.CareerDescriptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCareerTasks chains the current query on the "careerTasks" edge.
func (ucq *UserCareerQuery) QueryCareerTasks() *CareerTaskQuery {
	query := &CareerTaskQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareer.Table, usercareer.FieldID, selector),
			sqlgraph.To(careertask.Table, careertask.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usercareer.CareerTasksTable, usercareer.CareerTasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCareerSkillGroups chains the current query on the "careerSkillGroups" edge.
func (ucq *UserCareerQuery) QueryCareerSkillGroups() *CareerSkillGroupQuery {
	query := &CareerSkillGroupQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercareer.Table, usercareer.FieldID, selector),
			sqlgraph.To(careerskillgroup.Table, careerskillgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usercareer.CareerSkillGroupsTable, usercareer.CareerSkillGroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCareer entity from the query.
// Returns a *NotFoundError when no UserCareer was found.
func (ucq *UserCareerQuery) First(ctx context.Context) (*UserCareer, error) {
	nodes, err := ucq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercareer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserCareerQuery) FirstX(ctx context.Context) *UserCareer {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCareer ID from the query.
// Returns a *NotFoundError when no UserCareer ID was found.
func (ucq *UserCareerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercareer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserCareerQuery) FirstIDX(ctx context.Context) int {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCareer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCareer entity is found.
// Returns a *NotFoundError when no UserCareer entities are found.
func (ucq *UserCareerQuery) Only(ctx context.Context) (*UserCareer, error) {
	nodes, err := ucq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercareer.Label}
	default:
		return nil, &NotSingularError{usercareer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserCareerQuery) OnlyX(ctx context.Context) *UserCareer {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCareer ID in the query.
// Returns a *NotSingularError when more than one UserCareer ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserCareerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercareer.Label}
	default:
		err = &NotSingularError{usercareer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserCareerQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCareers.
func (ucq *UserCareerQuery) All(ctx context.Context) ([]*UserCareer, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserCareerQuery) AllX(ctx context.Context) []*UserCareer {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCareer IDs.
func (ucq *UserCareerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ucq.Select(usercareer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserCareerQuery) IDsX(ctx context.Context) []int {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserCareerQuery) Count(ctx context.Context) (int, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserCareerQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserCareerQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserCareerQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCareerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserCareerQuery) Clone() *UserCareerQuery {
	if ucq == nil {
		return nil
	}
	return &UserCareerQuery{
		config:                 ucq.config,
		limit:                  ucq.limit,
		offset:                 ucq.offset,
		order:                  append([]OrderFunc{}, ucq.order...),
		predicates:             append([]predicate.UserCareer{}, ucq.predicates...),
		withCareerGroup:        ucq.withCareerGroup.Clone(),
		withCareerDescriptions: ucq.withCareerDescriptions.Clone(),
		withCareerTasks:        ucq.withCareerTasks.Clone(),
		withCareerSkillGroups:  ucq.withCareerSkillGroups.Clone(),
		// clone intermediate query.
		sql:    ucq.sql.Clone(),
		path:   ucq.path,
		unique: ucq.unique,
	}
}

// WithCareerGroup tells the query-builder to eager-load the nodes that are connected to
// the "careerGroup" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCareerQuery) WithCareerGroup(opts ...func(*UserCareerGroupQuery)) *UserCareerQuery {
	query := &UserCareerGroupQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withCareerGroup = query
	return ucq
}

// WithCareerDescriptions tells the query-builder to eager-load the nodes that are connected to
// the "careerDescriptions" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCareerQuery) WithCareerDescriptions(opts ...func(*UserCareerDescriptionQuery)) *UserCareerQuery {
	query := &UserCareerDescriptionQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withCareerDescriptions = query
	return ucq
}

// WithCareerTasks tells the query-builder to eager-load the nodes that are connected to
// the "careerTasks" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCareerQuery) WithCareerTasks(opts ...func(*CareerTaskQuery)) *UserCareerQuery {
	query := &CareerTaskQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withCareerTasks = query
	return ucq
}

// WithCareerSkillGroups tells the query-builder to eager-load the nodes that are connected to
// the "careerSkillGroups" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCareerQuery) WithCareerSkillGroups(opts ...func(*CareerSkillGroupQuery)) *UserCareerQuery {
	query := &CareerSkillGroupQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withCareerSkillGroups = query
	return ucq
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
//	client.UserCareer.Query().
//		GroupBy(usercareer.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ucq *UserCareerQuery) GroupBy(field string, fields ...string) *UserCareerGroupBy {
	grbuild := &UserCareerGroupBy{config: ucq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucq.sqlQuery(ctx), nil
	}
	grbuild.label = usercareer.Label
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
//	client.UserCareer.Query().
//		Select(usercareer.FieldCreateTime).
//		Scan(ctx, &v)
func (ucq *UserCareerQuery) Select(fields ...string) *UserCareerSelect {
	ucq.fields = append(ucq.fields, fields...)
	selbuild := &UserCareerSelect{UserCareerQuery: ucq}
	selbuild.label = usercareer.Label
	selbuild.flds, selbuild.scan = &ucq.fields, selbuild.Scan
	return selbuild
}

func (ucq *UserCareerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucq.fields {
		if !usercareer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserCareerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCareer, error) {
	var (
		nodes       = []*UserCareer{}
		withFKs     = ucq.withFKs
		_spec       = ucq.querySpec()
		loadedTypes = [4]bool{
			ucq.withCareerGroup != nil,
			ucq.withCareerDescriptions != nil,
			ucq.withCareerTasks != nil,
			ucq.withCareerSkillGroups != nil,
		}
	)
	if ucq.withCareerGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usercareer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserCareer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserCareer{config: ucq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucq.withCareerGroup; query != nil {
		if err := ucq.loadCareerGroup(ctx, query, nodes, nil,
			func(n *UserCareer, e *UserCareerGroup) { n.Edges.CareerGroup = e }); err != nil {
			return nil, err
		}
	}
	if query := ucq.withCareerDescriptions; query != nil {
		if err := ucq.loadCareerDescriptions(ctx, query, nodes,
			func(n *UserCareer) { n.Edges.CareerDescriptions = []*UserCareerDescription{} },
			func(n *UserCareer, e *UserCareerDescription) {
				n.Edges.CareerDescriptions = append(n.Edges.CareerDescriptions, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := ucq.withCareerTasks; query != nil {
		if err := ucq.loadCareerTasks(ctx, query, nodes,
			func(n *UserCareer) { n.Edges.CareerTasks = []*CareerTask{} },
			func(n *UserCareer, e *CareerTask) { n.Edges.CareerTasks = append(n.Edges.CareerTasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := ucq.withCareerSkillGroups; query != nil {
		if err := ucq.loadCareerSkillGroups(ctx, query, nodes,
			func(n *UserCareer) { n.Edges.CareerSkillGroups = []*CareerSkillGroup{} },
			func(n *UserCareer, e *CareerSkillGroup) {
				n.Edges.CareerSkillGroups = append(n.Edges.CareerSkillGroups, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucq *UserCareerQuery) loadCareerGroup(ctx context.Context, query *UserCareerGroupQuery, nodes []*UserCareer, init func(*UserCareer), assign func(*UserCareer, *UserCareerGroup)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserCareer)
	for i := range nodes {
		if nodes[i].career_group_id == nil {
			continue
		}
		fk := *nodes[i].career_group_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(usercareergroup.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ucq *UserCareerQuery) loadCareerDescriptions(ctx context.Context, query *UserCareerDescriptionQuery, nodes []*UserCareer, init func(*UserCareer), assign func(*UserCareer, *UserCareerDescription)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserCareer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserCareerDescription(func(s *sql.Selector) {
		s.Where(sql.InValues(usercareer.CareerDescriptionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.career_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "career_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ucq *UserCareerQuery) loadCareerTasks(ctx context.Context, query *CareerTaskQuery, nodes []*UserCareer, init func(*UserCareer), assign func(*UserCareer, *CareerTask)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserCareer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CareerTask(func(s *sql.Selector) {
		s.Where(sql.InValues(usercareer.CareerTasksColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.career_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "career_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ucq *UserCareerQuery) loadCareerSkillGroups(ctx context.Context, query *CareerSkillGroupQuery, nodes []*UserCareer, init func(*UserCareer), assign func(*UserCareer, *CareerSkillGroup)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserCareer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CareerSkillGroup(func(s *sql.Selector) {
		s.Where(sql.InValues(usercareer.CareerSkillGroupsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.career_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "career_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "career_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ucq *UserCareerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	_spec.Node.Columns = ucq.fields
	if len(ucq.fields) > 0 {
		_spec.Unique = ucq.unique != nil && *ucq.unique
	}
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserCareerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucq *UserCareerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercareer.Table,
			Columns: usercareer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usercareer.FieldID,
			},
		},
		From:   ucq.sql,
		Unique: true,
	}
	if unique := ucq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercareer.FieldID)
		for i := range fields {
			if fields[i] != usercareer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserCareerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(usercareer.Table)
	columns := ucq.fields
	if len(columns) == 0 {
		columns = usercareer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucq.unique != nil && *ucq.unique {
		selector.Distinct()
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCareerGroupBy is the group-by builder for UserCareer entities.
type UserCareerGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserCareerGroupBy) Aggregate(fns ...AggregateFunc) *UserCareerGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucgb *UserCareerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucgb.path(ctx)
	if err != nil {
		return err
	}
	ucgb.sql = query
	return ucgb.sqlScan(ctx, v)
}

func (ucgb *UserCareerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucgb.fields {
		if !usercareer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucgb *UserCareerGroupBy) sqlQuery() *sql.Selector {
	selector := ucgb.sql.Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucgb.fields)+len(ucgb.fns))
		for _, f := range ucgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucgb.fields...)...)
}

// UserCareerSelect is the builder for selecting fields of UserCareer entities.
type UserCareerSelect struct {
	*UserCareerQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserCareerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	ucs.sql = ucs.UserCareerQuery.sqlQuery(ctx)
	return ucs.sqlScan(ctx, v)
}

func (ucs *UserCareerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucs.sql.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
