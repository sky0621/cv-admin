// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerTaskDelete is the builder for deleting a CareerTask entity.
type CareerTaskDelete struct {
	config
	hooks    []Hook
	mutation *CareerTaskMutation
}

// Where appends a list predicates to the CareerTaskDelete builder.
func (ctd *CareerTaskDelete) Where(ps ...predicate.CareerTask) *CareerTaskDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CareerTaskDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ctd.sqlExec, ctd.mutation, ctd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CareerTaskDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CareerTaskDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(careertask.Table, sqlgraph.NewFieldSpec(careertask.FieldID, field.TypeInt))
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ctd.mutation.done = true
	return affected, err
}

// CareerTaskDeleteOne is the builder for deleting a single CareerTask entity.
type CareerTaskDeleteOne struct {
	ctd *CareerTaskDelete
}

// Where appends a list predicates to the CareerTaskDelete builder.
func (ctdo *CareerTaskDeleteOne) Where(ps ...predicate.CareerTask) *CareerTaskDeleteOne {
	ctdo.ctd.mutation.Where(ps...)
	return ctdo
}

// Exec executes the deletion query.
func (ctdo *CareerTaskDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{careertask.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CareerTaskDeleteOne) ExecX(ctx context.Context) {
	if err := ctdo.Exec(ctx); err != nil {
		panic(err)
	}
}
