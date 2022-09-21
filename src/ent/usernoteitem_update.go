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
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
)

// UserNoteItemUpdate is the builder for updating UserNoteItem entities.
type UserNoteItemUpdate struct {
	config
	hooks    []Hook
	mutation *UserNoteItemMutation
}

// Where appends a list predicates to the UserNoteItemUpdate builder.
func (uniu *UserNoteItemUpdate) Where(ps ...predicate.UserNoteItem) *UserNoteItemUpdate {
	uniu.mutation.Where(ps...)
	return uniu
}

// SetUpdateTime sets the "update_time" field.
func (uniu *UserNoteItemUpdate) SetUpdateTime(t time.Time) *UserNoteItemUpdate {
	uniu.mutation.SetUpdateTime(t)
	return uniu
}

// Mutation returns the UserNoteItemMutation object of the builder.
func (uniu *UserNoteItemUpdate) Mutation() *UserNoteItemMutation {
	return uniu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uniu *UserNoteItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uniu.defaults()
	if len(uniu.hooks) == 0 {
		affected, err = uniu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNoteItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uniu.mutation = mutation
			affected, err = uniu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uniu.hooks) - 1; i >= 0; i-- {
			if uniu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uniu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uniu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uniu *UserNoteItemUpdate) SaveX(ctx context.Context) int {
	affected, err := uniu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uniu *UserNoteItemUpdate) Exec(ctx context.Context) error {
	_, err := uniu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uniu *UserNoteItemUpdate) ExecX(ctx context.Context) {
	if err := uniu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uniu *UserNoteItemUpdate) defaults() {
	if _, ok := uniu.mutation.UpdateTime(); !ok {
		v := usernoteitem.UpdateDefaultUpdateTime()
		uniu.mutation.SetUpdateTime(v)
	}
}

func (uniu *UserNoteItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernoteitem.Table,
			Columns: usernoteitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernoteitem.FieldID,
			},
		},
	}
	if ps := uniu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uniu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usernoteitem.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uniu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernoteitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserNoteItemUpdateOne is the builder for updating a single UserNoteItem entity.
type UserNoteItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserNoteItemMutation
}

// SetUpdateTime sets the "update_time" field.
func (uniuo *UserNoteItemUpdateOne) SetUpdateTime(t time.Time) *UserNoteItemUpdateOne {
	uniuo.mutation.SetUpdateTime(t)
	return uniuo
}

// Mutation returns the UserNoteItemMutation object of the builder.
func (uniuo *UserNoteItemUpdateOne) Mutation() *UserNoteItemMutation {
	return uniuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uniuo *UserNoteItemUpdateOne) Select(field string, fields ...string) *UserNoteItemUpdateOne {
	uniuo.fields = append([]string{field}, fields...)
	return uniuo
}

// Save executes the query and returns the updated UserNoteItem entity.
func (uniuo *UserNoteItemUpdateOne) Save(ctx context.Context) (*UserNoteItem, error) {
	var (
		err  error
		node *UserNoteItem
	)
	uniuo.defaults()
	if len(uniuo.hooks) == 0 {
		node, err = uniuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNoteItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uniuo.mutation = mutation
			node, err = uniuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uniuo.hooks) - 1; i >= 0; i-- {
			if uniuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uniuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uniuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserNoteItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserNoteItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uniuo *UserNoteItemUpdateOne) SaveX(ctx context.Context) *UserNoteItem {
	node, err := uniuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uniuo *UserNoteItemUpdateOne) Exec(ctx context.Context) error {
	_, err := uniuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uniuo *UserNoteItemUpdateOne) ExecX(ctx context.Context) {
	if err := uniuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uniuo *UserNoteItemUpdateOne) defaults() {
	if _, ok := uniuo.mutation.UpdateTime(); !ok {
		v := usernoteitem.UpdateDefaultUpdateTime()
		uniuo.mutation.SetUpdateTime(v)
	}
}

func (uniuo *UserNoteItemUpdateOne) sqlSave(ctx context.Context) (_node *UserNoteItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernoteitem.Table,
			Columns: usernoteitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernoteitem.FieldID,
			},
		},
	}
	id, ok := uniuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserNoteItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uniuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernoteitem.FieldID)
		for _, f := range fields {
			if !usernoteitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usernoteitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uniuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uniuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usernoteitem.FieldUpdateTime,
		})
	}
	_node = &UserNoteItem{config: uniuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uniuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernoteitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
