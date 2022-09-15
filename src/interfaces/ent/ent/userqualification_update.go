// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/predicate"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/userqualification"
)

// UserQualificationUpdate is the builder for updating UserQualification entities.
type UserQualificationUpdate struct {
	config
	hooks    []Hook
	mutation *UserQualificationMutation
}

// Where appends a list predicates to the UserQualificationUpdate builder.
func (uqu *UserQualificationUpdate) Where(ps ...predicate.UserQualification) *UserQualificationUpdate {
	uqu.mutation.Where(ps...)
	return uqu
}

// Mutation returns the UserQualificationMutation object of the builder.
func (uqu *UserQualificationUpdate) Mutation() *UserQualificationMutation {
	return uqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uqu *UserQualificationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uqu.hooks) == 0 {
		affected, err = uqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserQualificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uqu.mutation = mutation
			affected, err = uqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uqu.hooks) - 1; i >= 0; i-- {
			if uqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uqu *UserQualificationUpdate) SaveX(ctx context.Context) int {
	affected, err := uqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uqu *UserQualificationUpdate) Exec(ctx context.Context) error {
	_, err := uqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uqu *UserQualificationUpdate) ExecX(ctx context.Context) {
	if err := uqu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uqu *UserQualificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userqualification.Table,
			Columns: userqualification.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userqualification.FieldID,
			},
		},
	}
	if ps := uqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userqualification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserQualificationUpdateOne is the builder for updating a single UserQualification entity.
type UserQualificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserQualificationMutation
}

// Mutation returns the UserQualificationMutation object of the builder.
func (uquo *UserQualificationUpdateOne) Mutation() *UserQualificationMutation {
	return uquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uquo *UserQualificationUpdateOne) Select(field string, fields ...string) *UserQualificationUpdateOne {
	uquo.fields = append([]string{field}, fields...)
	return uquo
}

// Save executes the query and returns the updated UserQualification entity.
func (uquo *UserQualificationUpdateOne) Save(ctx context.Context) (*UserQualification, error) {
	var (
		err  error
		node *UserQualification
	)
	if len(uquo.hooks) == 0 {
		node, err = uquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserQualificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uquo.mutation = mutation
			node, err = uquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uquo.hooks) - 1; i >= 0; i-- {
			if uquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uquo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uquo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserQualification)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserQualificationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uquo *UserQualificationUpdateOne) SaveX(ctx context.Context) *UserQualification {
	node, err := uquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uquo *UserQualificationUpdateOne) Exec(ctx context.Context) error {
	_, err := uquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uquo *UserQualificationUpdateOne) ExecX(ctx context.Context) {
	if err := uquo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uquo *UserQualificationUpdateOne) sqlSave(ctx context.Context) (_node *UserQualification, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userqualification.Table,
			Columns: userqualification.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userqualification.FieldID,
			},
		},
	}
	id, ok := uquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserQualification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userqualification.FieldID)
		for _, f := range fields {
			if !userqualification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userqualification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &UserQualification{config: uquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userqualification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
