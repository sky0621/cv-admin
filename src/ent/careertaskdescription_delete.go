// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careertaskdescription"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerTaskDescriptionDelete is the builder for deleting a CareerTaskDescription entity.
type CareerTaskDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *CareerTaskDescriptionMutation
}

// Where appends a list predicates to the CareerTaskDescriptionDelete builder.
func (ctdd *CareerTaskDescriptionDelete) Where(ps ...predicate.CareerTaskDescription) *CareerTaskDescriptionDelete {
	ctdd.mutation.Where(ps...)
	return ctdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctdd *CareerTaskDescriptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctdd.hooks) == 0 {
		affected, err = ctdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerTaskDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctdd.mutation = mutation
			affected, err = ctdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctdd.hooks) - 1; i >= 0; i-- {
			if ctdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdd *CareerTaskDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := ctdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctdd *CareerTaskDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: careertaskdescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careertaskdescription.FieldID,
			},
		},
	}
	if ps := ctdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CareerTaskDescriptionDeleteOne is the builder for deleting a single CareerTaskDescription entity.
type CareerTaskDescriptionDeleteOne struct {
	ctdd *CareerTaskDescriptionDelete
}

// Exec executes the deletion query.
func (ctddo *CareerTaskDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := ctddo.ctdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{careertaskdescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctddo *CareerTaskDescriptionDeleteOne) ExecX(ctx context.Context) {
	ctddo.ctdd.ExecX(ctx)
}
