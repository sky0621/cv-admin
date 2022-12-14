// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// CareerSkillGroupDelete is the builder for deleting a CareerSkillGroup entity.
type CareerSkillGroupDelete struct {
	config
	hooks    []Hook
	mutation *CareerSkillGroupMutation
}

// Where appends a list predicates to the CareerSkillGroupDelete builder.
func (csgd *CareerSkillGroupDelete) Where(ps ...predicate.CareerSkillGroup) *CareerSkillGroupDelete {
	csgd.mutation.Where(ps...)
	return csgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csgd *CareerSkillGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csgd.hooks) == 0 {
		affected, err = csgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csgd.mutation = mutation
			affected, err = csgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csgd.hooks) - 1; i >= 0; i-- {
			if csgd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (csgd *CareerSkillGroupDelete) ExecX(ctx context.Context) int {
	n, err := csgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csgd *CareerSkillGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: careerskillgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskillgroup.FieldID,
			},
		},
	}
	if ps := csgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, csgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CareerSkillGroupDeleteOne is the builder for deleting a single CareerSkillGroup entity.
type CareerSkillGroupDeleteOne struct {
	csgd *CareerSkillGroupDelete
}

// Exec executes the deletion query.
func (csgdo *CareerSkillGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := csgdo.csgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{careerskillgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csgdo *CareerSkillGroupDeleteOne) ExecX(ctx context.Context) {
	csgdo.csgd.ExecX(ctx)
}
