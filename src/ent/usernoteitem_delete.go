// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
)

// UserNoteItemDelete is the builder for deleting a UserNoteItem entity.
type UserNoteItemDelete struct {
	config
	hooks    []Hook
	mutation *UserNoteItemMutation
}

// Where appends a list predicates to the UserNoteItemDelete builder.
func (unid *UserNoteItemDelete) Where(ps ...predicate.UserNoteItem) *UserNoteItemDelete {
	unid.mutation.Where(ps...)
	return unid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (unid *UserNoteItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(unid.hooks) == 0 {
		affected, err = unid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNoteItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			unid.mutation = mutation
			affected, err = unid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(unid.hooks) - 1; i >= 0; i-- {
			if unid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = unid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, unid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (unid *UserNoteItemDelete) ExecX(ctx context.Context) int {
	n, err := unid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (unid *UserNoteItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usernoteitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernoteitem.FieldID,
			},
		},
	}
	if ps := unid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, unid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserNoteItemDeleteOne is the builder for deleting a single UserNoteItem entity.
type UserNoteItemDeleteOne struct {
	unid *UserNoteItemDelete
}

// Exec executes the deletion query.
func (unido *UserNoteItemDeleteOne) Exec(ctx context.Context) error {
	n, err := unido.unid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usernoteitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (unido *UserNoteItemDeleteOne) ExecX(ctx context.Context) {
	unido.unid.ExecX(ctx)
}
