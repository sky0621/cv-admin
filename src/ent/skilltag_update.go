// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/skilltag"
)

// SkillTagUpdate is the builder for updating SkillTag entities.
type SkillTagUpdate struct {
	config
	hooks    []Hook
	mutation *SkillTagMutation
}

// Where appends a list predicates to the SkillTagUpdate builder.
func (stu *SkillTagUpdate) Where(ps ...predicate.SkillTag) *SkillTagUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetName sets the "name" field.
func (stu *SkillTagUpdate) SetName(s string) *SkillTagUpdate {
	stu.mutation.SetName(s)
	return stu
}

// SetKey sets the "key" field.
func (stu *SkillTagUpdate) SetKey(s string) *SkillTagUpdate {
	stu.mutation.SetKey(s)
	return stu
}

// Mutation returns the SkillTagMutation object of the builder.
func (stu *SkillTagUpdate) Mutation() *SkillTagMutation {
	return stu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SkillTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		if err = stu.check(); err != nil {
			return 0, err
		}
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stu.check(); err != nil {
				return 0, err
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			if stu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SkillTagUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SkillTagUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SkillTagUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *SkillTagUpdate) check() error {
	if v, ok := stu.mutation.Name(); ok {
		if err := skilltag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillTag.name": %w`, err)}
		}
	}
	if v, ok := stu.mutation.Key(); ok {
		if err := skilltag.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "SkillTag.key": %w`, err)}
		}
	}
	return nil
}

func (stu *SkillTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skilltag.Table,
			Columns: skilltag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skilltag.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.Name(); ok {
		_spec.SetField(skilltag.FieldName, field.TypeString, value)
	}
	if value, ok := stu.mutation.Key(); ok {
		_spec.SetField(skilltag.FieldKey, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SkillTagUpdateOne is the builder for updating a single SkillTag entity.
type SkillTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillTagMutation
}

// SetName sets the "name" field.
func (stuo *SkillTagUpdateOne) SetName(s string) *SkillTagUpdateOne {
	stuo.mutation.SetName(s)
	return stuo
}

// SetKey sets the "key" field.
func (stuo *SkillTagUpdateOne) SetKey(s string) *SkillTagUpdateOne {
	stuo.mutation.SetKey(s)
	return stuo
}

// Mutation returns the SkillTagMutation object of the builder.
func (stuo *SkillTagUpdateOne) Mutation() *SkillTagMutation {
	return stuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SkillTagUpdateOne) Select(field string, fields ...string) *SkillTagUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SkillTag entity.
func (stuo *SkillTagUpdateOne) Save(ctx context.Context) (*SkillTag, error) {
	var (
		err  error
		node *SkillTag
	)
	if len(stuo.hooks) == 0 {
		if err = stuo.check(); err != nil {
			return nil, err
		}
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stuo.check(); err != nil {
				return nil, err
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			if stuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, stuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SkillTag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SkillTagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SkillTagUpdateOne) SaveX(ctx context.Context) *SkillTag {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SkillTagUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SkillTagUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *SkillTagUpdateOne) check() error {
	if v, ok := stuo.mutation.Name(); ok {
		if err := skilltag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillTag.name": %w`, err)}
		}
	}
	if v, ok := stuo.mutation.Key(); ok {
		if err := skilltag.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "SkillTag.key": %w`, err)}
		}
	}
	return nil
}

func (stuo *SkillTagUpdateOne) sqlSave(ctx context.Context) (_node *SkillTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skilltag.Table,
			Columns: skilltag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skilltag.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkillTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skilltag.FieldID)
		for _, f := range fields {
			if !skilltag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skilltag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.Name(); ok {
		_spec.SetField(skilltag.FieldName, field.TypeString, value)
	}
	if value, ok := stuo.mutation.Key(); ok {
		_spec.SetField(skilltag.FieldKey, field.TypeString, value)
	}
	_node = &SkillTag{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
