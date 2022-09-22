// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
)

// UserActivityUpdate is the builder for updating UserActivity entities.
type UserActivityUpdate struct {
	config
	hooks    []Hook
	mutation *UserActivityMutation
}

// Where appends a list predicates to the UserActivityUpdate builder.
func (uau *UserActivityUpdate) Where(ps ...predicate.UserActivity) *UserActivityUpdate {
	uau.mutation.Where(ps...)
	return uau
}

// SetUpdateTime sets the "update_time" field.
func (uau *UserActivityUpdate) SetUpdateTime(t time.Time) *UserActivityUpdate {
	uau.mutation.SetUpdateTime(t)
	return uau
}

// SetName sets the "name" field.
func (uau *UserActivityUpdate) SetName(s string) *UserActivityUpdate {
	uau.mutation.SetName(s)
	return uau
}

// SetURL sets the "url" field.
func (uau *UserActivityUpdate) SetURL(s string) *UserActivityUpdate {
	uau.mutation.SetURL(s)
	return uau
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uau *UserActivityUpdate) SetNillableURL(s *string) *UserActivityUpdate {
	if s != nil {
		uau.SetURL(*s)
	}
	return uau
}

// ClearURL clears the value of the "url" field.
func (uau *UserActivityUpdate) ClearURL() *UserActivityUpdate {
	uau.mutation.ClearURL()
	return uau
}

// SetIcon sets the "icon" field.
func (uau *UserActivityUpdate) SetIcon(s string) *UserActivityUpdate {
	uau.mutation.SetIcon(s)
	return uau
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (uau *UserActivityUpdate) SetNillableIcon(s *string) *UserActivityUpdate {
	if s != nil {
		uau.SetIcon(*s)
	}
	return uau
}

// ClearIcon clears the value of the "icon" field.
func (uau *UserActivityUpdate) ClearIcon() *UserActivityUpdate {
	uau.mutation.ClearIcon()
	return uau
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uau *UserActivityUpdate) SetUserID(id int) *UserActivityUpdate {
	uau.mutation.SetUserID(id)
	return uau
}

// SetUser sets the "user" edge to the User entity.
func (uau *UserActivityUpdate) SetUser(u *User) *UserActivityUpdate {
	return uau.SetUserID(u.ID)
}

// Mutation returns the UserActivityMutation object of the builder.
func (uau *UserActivityUpdate) Mutation() *UserActivityMutation {
	return uau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uau *UserActivityUpdate) ClearUser() *UserActivityUpdate {
	uau.mutation.ClearUser()
	return uau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uau *UserActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uau.defaults()
	if len(uau.hooks) == 0 {
		if err = uau.check(); err != nil {
			return 0, err
		}
		affected, err = uau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uau.check(); err != nil {
				return 0, err
			}
			uau.mutation = mutation
			affected, err = uau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uau.hooks) - 1; i >= 0; i-- {
			if uau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uau *UserActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := uau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uau *UserActivityUpdate) Exec(ctx context.Context) error {
	_, err := uau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uau *UserActivityUpdate) ExecX(ctx context.Context) {
	if err := uau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uau *UserActivityUpdate) defaults() {
	if _, ok := uau.mutation.UpdateTime(); !ok {
		v := useractivity.UpdateDefaultUpdateTime()
		uau.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uau *UserActivityUpdate) check() error {
	if v, ok := uau.mutation.Name(); ok {
		if err := useractivity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserActivity.name": %w`, err)}
		}
	}
	if v, ok := uau.mutation.URL(); ok {
		if err := useractivity.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "UserActivity.url": %w`, err)}
		}
	}
	if v, ok := uau.mutation.Icon(); ok {
		if err := useractivity.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "UserActivity.icon": %w`, err)}
		}
	}
	if _, ok := uau.mutation.UserID(); uau.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserActivity.user"`)
	}
	return nil
}

func (uau *UserActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useractivity.Table,
			Columns: useractivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useractivity.FieldID,
			},
		},
	}
	if ps := uau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uau.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useractivity.FieldUpdateTime,
		})
	}
	if value, ok := uau.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldName,
		})
	}
	if value, ok := uau.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldURL,
		})
	}
	if uau.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: useractivity.FieldURL,
		})
	}
	if value, ok := uau.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldIcon,
		})
	}
	if uau.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: useractivity.FieldIcon,
		})
	}
	if uau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useractivity.UserTable,
			Columns: []string{useractivity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useractivity.UserTable,
			Columns: []string{useractivity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{useractivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserActivityUpdateOne is the builder for updating a single UserActivity entity.
type UserActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserActivityMutation
}

// SetUpdateTime sets the "update_time" field.
func (uauo *UserActivityUpdateOne) SetUpdateTime(t time.Time) *UserActivityUpdateOne {
	uauo.mutation.SetUpdateTime(t)
	return uauo
}

// SetName sets the "name" field.
func (uauo *UserActivityUpdateOne) SetName(s string) *UserActivityUpdateOne {
	uauo.mutation.SetName(s)
	return uauo
}

// SetURL sets the "url" field.
func (uauo *UserActivityUpdateOne) SetURL(s string) *UserActivityUpdateOne {
	uauo.mutation.SetURL(s)
	return uauo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uauo *UserActivityUpdateOne) SetNillableURL(s *string) *UserActivityUpdateOne {
	if s != nil {
		uauo.SetURL(*s)
	}
	return uauo
}

// ClearURL clears the value of the "url" field.
func (uauo *UserActivityUpdateOne) ClearURL() *UserActivityUpdateOne {
	uauo.mutation.ClearURL()
	return uauo
}

// SetIcon sets the "icon" field.
func (uauo *UserActivityUpdateOne) SetIcon(s string) *UserActivityUpdateOne {
	uauo.mutation.SetIcon(s)
	return uauo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (uauo *UserActivityUpdateOne) SetNillableIcon(s *string) *UserActivityUpdateOne {
	if s != nil {
		uauo.SetIcon(*s)
	}
	return uauo
}

// ClearIcon clears the value of the "icon" field.
func (uauo *UserActivityUpdateOne) ClearIcon() *UserActivityUpdateOne {
	uauo.mutation.ClearIcon()
	return uauo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uauo *UserActivityUpdateOne) SetUserID(id int) *UserActivityUpdateOne {
	uauo.mutation.SetUserID(id)
	return uauo
}

// SetUser sets the "user" edge to the User entity.
func (uauo *UserActivityUpdateOne) SetUser(u *User) *UserActivityUpdateOne {
	return uauo.SetUserID(u.ID)
}

// Mutation returns the UserActivityMutation object of the builder.
func (uauo *UserActivityUpdateOne) Mutation() *UserActivityMutation {
	return uauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uauo *UserActivityUpdateOne) ClearUser() *UserActivityUpdateOne {
	uauo.mutation.ClearUser()
	return uauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uauo *UserActivityUpdateOne) Select(field string, fields ...string) *UserActivityUpdateOne {
	uauo.fields = append([]string{field}, fields...)
	return uauo
}

// Save executes the query and returns the updated UserActivity entity.
func (uauo *UserActivityUpdateOne) Save(ctx context.Context) (*UserActivity, error) {
	var (
		err  error
		node *UserActivity
	)
	uauo.defaults()
	if len(uauo.hooks) == 0 {
		if err = uauo.check(); err != nil {
			return nil, err
		}
		node, err = uauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uauo.check(); err != nil {
				return nil, err
			}
			uauo.mutation = mutation
			node, err = uauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uauo.hooks) - 1; i >= 0; i-- {
			if uauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserActivity)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserActivityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uauo *UserActivityUpdateOne) SaveX(ctx context.Context) *UserActivity {
	node, err := uauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uauo *UserActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := uauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uauo *UserActivityUpdateOne) ExecX(ctx context.Context) {
	if err := uauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uauo *UserActivityUpdateOne) defaults() {
	if _, ok := uauo.mutation.UpdateTime(); !ok {
		v := useractivity.UpdateDefaultUpdateTime()
		uauo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uauo *UserActivityUpdateOne) check() error {
	if v, ok := uauo.mutation.Name(); ok {
		if err := useractivity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserActivity.name": %w`, err)}
		}
	}
	if v, ok := uauo.mutation.URL(); ok {
		if err := useractivity.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "UserActivity.url": %w`, err)}
		}
	}
	if v, ok := uauo.mutation.Icon(); ok {
		if err := useractivity.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "UserActivity.icon": %w`, err)}
		}
	}
	if _, ok := uauo.mutation.UserID(); uauo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserActivity.user"`)
	}
	return nil
}

func (uauo *UserActivityUpdateOne) sqlSave(ctx context.Context) (_node *UserActivity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useractivity.Table,
			Columns: useractivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useractivity.FieldID,
			},
		},
	}
	id, ok := uauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserActivity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useractivity.FieldID)
		for _, f := range fields {
			if !useractivity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != useractivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uauo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useractivity.FieldUpdateTime,
		})
	}
	if value, ok := uauo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldName,
		})
	}
	if value, ok := uauo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldURL,
		})
	}
	if uauo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: useractivity.FieldURL,
		})
	}
	if value, ok := uauo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useractivity.FieldIcon,
		})
	}
	if uauo.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: useractivity.FieldIcon,
		})
	}
	if uauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useractivity.UserTable,
			Columns: []string{useractivity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useractivity.UserTable,
			Columns: []string{useractivity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserActivity{config: uauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{useractivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
