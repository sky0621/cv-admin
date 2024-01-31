// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
)

// UserCareerDescriptionDelete is the builder for deleting a UserCareerDescription entity.
type UserCareerDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *UserCareerDescriptionMutation
}

// Where appends a list predicates to the UserCareerDescriptionDelete builder.
func (ucdd *UserCareerDescriptionDelete) Where(ps ...predicate.UserCareerDescription) *UserCareerDescriptionDelete {
	ucdd.mutation.Where(ps...)
	return ucdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucdd *UserCareerDescriptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucdd.sqlExec, ucdd.mutation, ucdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdd *UserCareerDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := ucdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucdd *UserCareerDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usercareerdescription.Table, sqlgraph.NewFieldSpec(usercareerdescription.FieldID, field.TypeInt))
	if ps := ucdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucdd.mutation.done = true
	return affected, err
}

// UserCareerDescriptionDeleteOne is the builder for deleting a single UserCareerDescription entity.
type UserCareerDescriptionDeleteOne struct {
	ucdd *UserCareerDescriptionDelete
}

// Where appends a list predicates to the UserCareerDescriptionDelete builder.
func (ucddo *UserCareerDescriptionDeleteOne) Where(ps ...predicate.UserCareerDescription) *UserCareerDescriptionDeleteOne {
	ucddo.ucdd.mutation.Where(ps...)
	return ucddo
}

// Exec executes the deletion query.
func (ucddo *UserCareerDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := ucddo.ucdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usercareerdescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucddo *UserCareerDescriptionDeleteOne) ExecX(ctx context.Context) {
	if err := ucddo.Exec(ctx); err != nil {
		panic(err)
	}
}
