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
	"github.com/sky0621/cv-admin/src/ent/usernote"
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
)

// UserNoteQuery is the builder for querying UserNote entities.
type UserNoteQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.UserNote
	withUser      *UserQuery
	withNoteItems *UserNoteItemQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserNoteQuery builder.
func (unq *UserNoteQuery) Where(ps ...predicate.UserNote) *UserNoteQuery {
	unq.predicates = append(unq.predicates, ps...)
	return unq
}

// Limit adds a limit step to the query.
func (unq *UserNoteQuery) Limit(limit int) *UserNoteQuery {
	unq.limit = &limit
	return unq
}

// Offset adds an offset step to the query.
func (unq *UserNoteQuery) Offset(offset int) *UserNoteQuery {
	unq.offset = &offset
	return unq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (unq *UserNoteQuery) Unique(unique bool) *UserNoteQuery {
	unq.unique = &unique
	return unq
}

// Order adds an order step to the query.
func (unq *UserNoteQuery) Order(o ...OrderFunc) *UserNoteQuery {
	unq.order = append(unq.order, o...)
	return unq
}

// QueryUser chains the current query on the "user" edge.
func (unq *UserNoteQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: unq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := unq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernote.Table, usernote.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usernote.UserTable, usernote.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(unq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNoteItems chains the current query on the "noteItems" edge.
func (unq *UserNoteQuery) QueryNoteItems() *UserNoteItemQuery {
	query := &UserNoteItemQuery{config: unq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := unq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernote.Table, usernote.FieldID, selector),
			sqlgraph.To(usernoteitem.Table, usernoteitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usernote.NoteItemsTable, usernote.NoteItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(unq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserNote entity from the query.
// Returns a *NotFoundError when no UserNote was found.
func (unq *UserNoteQuery) First(ctx context.Context) (*UserNote, error) {
	nodes, err := unq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usernote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (unq *UserNoteQuery) FirstX(ctx context.Context) *UserNote {
	node, err := unq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserNote ID from the query.
// Returns a *NotFoundError when no UserNote ID was found.
func (unq *UserNoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = unq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usernote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (unq *UserNoteQuery) FirstIDX(ctx context.Context) int {
	id, err := unq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserNote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserNote entity is found.
// Returns a *NotFoundError when no UserNote entities are found.
func (unq *UserNoteQuery) Only(ctx context.Context) (*UserNote, error) {
	nodes, err := unq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usernote.Label}
	default:
		return nil, &NotSingularError{usernote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (unq *UserNoteQuery) OnlyX(ctx context.Context) *UserNote {
	node, err := unq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserNote ID in the query.
// Returns a *NotSingularError when more than one UserNote ID is found.
// Returns a *NotFoundError when no entities are found.
func (unq *UserNoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = unq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usernote.Label}
	default:
		err = &NotSingularError{usernote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (unq *UserNoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := unq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserNotes.
func (unq *UserNoteQuery) All(ctx context.Context) ([]*UserNote, error) {
	if err := unq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return unq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (unq *UserNoteQuery) AllX(ctx context.Context) []*UserNote {
	nodes, err := unq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserNote IDs.
func (unq *UserNoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := unq.Select(usernote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (unq *UserNoteQuery) IDsX(ctx context.Context) []int {
	ids, err := unq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (unq *UserNoteQuery) Count(ctx context.Context) (int, error) {
	if err := unq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return unq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (unq *UserNoteQuery) CountX(ctx context.Context) int {
	count, err := unq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (unq *UserNoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := unq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return unq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (unq *UserNoteQuery) ExistX(ctx context.Context) bool {
	exist, err := unq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserNoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (unq *UserNoteQuery) Clone() *UserNoteQuery {
	if unq == nil {
		return nil
	}
	return &UserNoteQuery{
		config:        unq.config,
		limit:         unq.limit,
		offset:        unq.offset,
		order:         append([]OrderFunc{}, unq.order...),
		predicates:    append([]predicate.UserNote{}, unq.predicates...),
		withUser:      unq.withUser.Clone(),
		withNoteItems: unq.withNoteItems.Clone(),
		// clone intermediate query.
		sql:    unq.sql.Clone(),
		path:   unq.path,
		unique: unq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (unq *UserNoteQuery) WithUser(opts ...func(*UserQuery)) *UserNoteQuery {
	query := &UserQuery{config: unq.config}
	for _, opt := range opts {
		opt(query)
	}
	unq.withUser = query
	return unq
}

// WithNoteItems tells the query-builder to eager-load the nodes that are connected to
// the "noteItems" edge. The optional arguments are used to configure the query builder of the edge.
func (unq *UserNoteQuery) WithNoteItems(opts ...func(*UserNoteItemQuery)) *UserNoteQuery {
	query := &UserNoteItemQuery{config: unq.config}
	for _, opt := range opts {
		opt(query)
	}
	unq.withNoteItems = query
	return unq
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
//	client.UserNote.Query().
//		GroupBy(usernote.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (unq *UserNoteQuery) GroupBy(field string, fields ...string) *UserNoteGroupBy {
	grbuild := &UserNoteGroupBy{config: unq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return unq.sqlQuery(ctx), nil
	}
	grbuild.label = usernote.Label
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
//	client.UserNote.Query().
//		Select(usernote.FieldCreateTime).
//		Scan(ctx, &v)
func (unq *UserNoteQuery) Select(fields ...string) *UserNoteSelect {
	unq.fields = append(unq.fields, fields...)
	selbuild := &UserNoteSelect{UserNoteQuery: unq}
	selbuild.label = usernote.Label
	selbuild.flds, selbuild.scan = &unq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a UserNoteSelect configured with the given aggregations.
func (unq *UserNoteQuery) Aggregate(fns ...AggregateFunc) *UserNoteSelect {
	return unq.Select().Aggregate(fns...)
}

func (unq *UserNoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range unq.fields {
		if !usernote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if unq.path != nil {
		prev, err := unq.path(ctx)
		if err != nil {
			return err
		}
		unq.sql = prev
	}
	return nil
}

func (unq *UserNoteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserNote, error) {
	var (
		nodes       = []*UserNote{}
		withFKs     = unq.withFKs
		_spec       = unq.querySpec()
		loadedTypes = [2]bool{
			unq.withUser != nil,
			unq.withNoteItems != nil,
		}
	)
	if unq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usernote.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserNote).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserNote{config: unq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, unq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := unq.withUser; query != nil {
		if err := unq.loadUser(ctx, query, nodes, nil,
			func(n *UserNote, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := unq.withNoteItems; query != nil {
		if err := unq.loadNoteItems(ctx, query, nodes,
			func(n *UserNote) { n.Edges.NoteItems = []*UserNoteItem{} },
			func(n *UserNote, e *UserNoteItem) { n.Edges.NoteItems = append(n.Edges.NoteItems, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (unq *UserNoteQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserNote, init func(*UserNote), assign func(*UserNote, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserNote)
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
func (unq *UserNoteQuery) loadNoteItems(ctx context.Context, query *UserNoteItemQuery, nodes []*UserNote, init func(*UserNote), assign func(*UserNote, *UserNoteItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserNote)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserNoteItem(func(s *sql.Selector) {
		s.Where(sql.InValues(usernote.NoteItemsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_note_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_note_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_note_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (unq *UserNoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := unq.querySpec()
	_spec.Node.Columns = unq.fields
	if len(unq.fields) > 0 {
		_spec.Unique = unq.unique != nil && *unq.unique
	}
	return sqlgraph.CountNodes(ctx, unq.driver, _spec)
}

func (unq *UserNoteQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := unq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (unq *UserNoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernote.Table,
			Columns: usernote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernote.FieldID,
			},
		},
		From:   unq.sql,
		Unique: true,
	}
	if unique := unq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := unq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernote.FieldID)
		for i := range fields {
			if fields[i] != usernote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := unq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := unq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := unq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := unq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (unq *UserNoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(unq.driver.Dialect())
	t1 := builder.Table(usernote.Table)
	columns := unq.fields
	if len(columns) == 0 {
		columns = usernote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if unq.sql != nil {
		selector = unq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if unq.unique != nil && *unq.unique {
		selector.Distinct()
	}
	for _, p := range unq.predicates {
		p(selector)
	}
	for _, p := range unq.order {
		p(selector)
	}
	if offset := unq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := unq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserNoteGroupBy is the group-by builder for UserNote entities.
type UserNoteGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ungb *UserNoteGroupBy) Aggregate(fns ...AggregateFunc) *UserNoteGroupBy {
	ungb.fns = append(ungb.fns, fns...)
	return ungb
}

// Scan applies the group-by query and scans the result into the given value.
func (ungb *UserNoteGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ungb.path(ctx)
	if err != nil {
		return err
	}
	ungb.sql = query
	return ungb.sqlScan(ctx, v)
}

func (ungb *UserNoteGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ungb.fields {
		if !usernote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ungb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ungb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ungb *UserNoteGroupBy) sqlQuery() *sql.Selector {
	selector := ungb.sql.Select()
	aggregation := make([]string, 0, len(ungb.fns))
	for _, fn := range ungb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ungb.fields)+len(ungb.fns))
		for _, f := range ungb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ungb.fields...)...)
}

// UserNoteSelect is the builder for selecting fields of UserNote entities.
type UserNoteSelect struct {
	*UserNoteQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uns *UserNoteSelect) Aggregate(fns ...AggregateFunc) *UserNoteSelect {
	uns.fns = append(uns.fns, fns...)
	return uns
}

// Scan applies the selector query and scans the result into the given value.
func (uns *UserNoteSelect) Scan(ctx context.Context, v any) error {
	if err := uns.prepareQuery(ctx); err != nil {
		return err
	}
	uns.sql = uns.UserNoteQuery.sqlQuery(ctx)
	return uns.sqlScan(ctx, v)
}

func (uns *UserNoteSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(uns.fns))
	for _, fn := range uns.fns {
		aggregation = append(aggregation, fn(uns.sql))
	}
	switch n := len(*uns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		uns.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		uns.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := uns.sql.Query()
	if err := uns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
