// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usersolution"
)

// UserSolutionDelete is the builder for deleting a UserSolution entity.
type UserSolutionDelete struct {
	config
	hooks    []Hook
	mutation *UserSolutionMutation
}

// Where appends a list predicates to the UserSolutionDelete builder.
func (usd *UserSolutionDelete) Where(ps ...predicate.UserSolution) *UserSolutionDelete {
	usd.mutation.Where(ps...)
	return usd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (usd *UserSolutionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, usd.sqlExec, usd.mutation, usd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (usd *UserSolutionDelete) ExecX(ctx context.Context) int {
	n, err := usd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (usd *UserSolutionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usersolution.Table, sqlgraph.NewFieldSpec(usersolution.FieldID, field.TypeInt))
	if ps := usd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, usd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	usd.mutation.done = true
	return affected, err
}

// UserSolutionDeleteOne is the builder for deleting a single UserSolution entity.
type UserSolutionDeleteOne struct {
	usd *UserSolutionDelete
}

// Where appends a list predicates to the UserSolutionDelete builder.
func (usdo *UserSolutionDeleteOne) Where(ps ...predicate.UserSolution) *UserSolutionDeleteOne {
	usdo.usd.mutation.Where(ps...)
	return usdo
}

// Exec executes the deletion query.
func (usdo *UserSolutionDeleteOne) Exec(ctx context.Context) error {
	n, err := usdo.usd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usersolution.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (usdo *UserSolutionDeleteOne) ExecX(ctx context.Context) {
	if err := usdo.Exec(ctx); err != nil {
		panic(err)
	}
}
