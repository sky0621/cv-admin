// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/predicate"
	"github.com/sky0621/cv-admin/src/interfaces/ent/ent/usernote"
)

// UserNoteDelete is the builder for deleting a UserNote entity.
type UserNoteDelete struct {
	config
	hooks    []Hook
	mutation *UserNoteMutation
}

// Where appends a list predicates to the UserNoteDelete builder.
func (und *UserNoteDelete) Where(ps ...predicate.UserNote) *UserNoteDelete {
	und.mutation.Where(ps...)
	return und
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (und *UserNoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(und.hooks) == 0 {
		affected, err = und.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			und.mutation = mutation
			affected, err = und.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(und.hooks) - 1; i >= 0; i-- {
			if und.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = und.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, und.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (und *UserNoteDelete) ExecX(ctx context.Context) int {
	n, err := und.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (und *UserNoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usernote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernote.FieldID,
			},
		},
	}
	if ps := und.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, und.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserNoteDeleteOne is the builder for deleting a single UserNote entity.
type UserNoteDeleteOne struct {
	und *UserNoteDelete
}

// Exec executes the deletion query.
func (undo *UserNoteDeleteOne) Exec(ctx context.Context) error {
	n, err := undo.und.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usernote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (undo *UserNoteDeleteOne) ExecX(ctx context.Context) {
	undo.und.ExecX(ctx)
}
