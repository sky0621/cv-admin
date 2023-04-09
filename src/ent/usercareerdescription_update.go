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
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
)

// UserCareerDescriptionUpdate is the builder for updating UserCareerDescription entities.
type UserCareerDescriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserCareerDescriptionMutation
}

// Where appends a list predicates to the UserCareerDescriptionUpdate builder.
func (ucdu *UserCareerDescriptionUpdate) Where(ps ...predicate.UserCareerDescription) *UserCareerDescriptionUpdate {
	ucdu.mutation.Where(ps...)
	return ucdu
}

// SetDescription sets the "description" field.
func (ucdu *UserCareerDescriptionUpdate) SetDescription(s string) *UserCareerDescriptionUpdate {
	ucdu.mutation.SetDescription(s)
	return ucdu
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ucdu *UserCareerDescriptionUpdate) SetCareerID(id int) *UserCareerDescriptionUpdate {
	ucdu.mutation.SetCareerID(id)
	return ucdu
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ucdu *UserCareerDescriptionUpdate) SetCareer(u *UserCareer) *UserCareerDescriptionUpdate {
	return ucdu.SetCareerID(u.ID)
}

// Mutation returns the UserCareerDescriptionMutation object of the builder.
func (ucdu *UserCareerDescriptionUpdate) Mutation() *UserCareerDescriptionMutation {
	return ucdu.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (ucdu *UserCareerDescriptionUpdate) ClearCareer() *UserCareerDescriptionUpdate {
	ucdu.mutation.ClearCareer()
	return ucdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucdu *UserCareerDescriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserCareerDescriptionMutation](ctx, ucdu.sqlSave, ucdu.mutation, ucdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucdu *UserCareerDescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := ucdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucdu *UserCareerDescriptionUpdate) Exec(ctx context.Context) error {
	_, err := ucdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdu *UserCareerDescriptionUpdate) ExecX(ctx context.Context) {
	if err := ucdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucdu *UserCareerDescriptionUpdate) check() error {
	if v, ok := ucdu.mutation.Description(); ok {
		if err := usercareerdescription.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "UserCareerDescription.description": %w`, err)}
		}
	}
	if _, ok := ucdu.mutation.CareerID(); ucdu.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCareerDescription.career"`)
	}
	return nil
}

func (ucdu *UserCareerDescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercareerdescription.Table, usercareerdescription.Columns, sqlgraph.NewFieldSpec(usercareerdescription.FieldID, field.TypeInt))
	if ps := ucdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucdu.mutation.Description(); ok {
		_spec.SetField(usercareerdescription.FieldDescription, field.TypeString, value)
	}
	if ucdu.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareerdescription.CareerTable,
			Columns: []string{usercareerdescription.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucdu.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareerdescription.CareerTable,
			Columns: []string{usercareerdescription.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercareerdescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucdu.mutation.done = true
	return n, nil
}

// UserCareerDescriptionUpdateOne is the builder for updating a single UserCareerDescription entity.
type UserCareerDescriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCareerDescriptionMutation
}

// SetDescription sets the "description" field.
func (ucduo *UserCareerDescriptionUpdateOne) SetDescription(s string) *UserCareerDescriptionUpdateOne {
	ucduo.mutation.SetDescription(s)
	return ucduo
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ucduo *UserCareerDescriptionUpdateOne) SetCareerID(id int) *UserCareerDescriptionUpdateOne {
	ucduo.mutation.SetCareerID(id)
	return ucduo
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ucduo *UserCareerDescriptionUpdateOne) SetCareer(u *UserCareer) *UserCareerDescriptionUpdateOne {
	return ucduo.SetCareerID(u.ID)
}

// Mutation returns the UserCareerDescriptionMutation object of the builder.
func (ucduo *UserCareerDescriptionUpdateOne) Mutation() *UserCareerDescriptionMutation {
	return ucduo.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (ucduo *UserCareerDescriptionUpdateOne) ClearCareer() *UserCareerDescriptionUpdateOne {
	ucduo.mutation.ClearCareer()
	return ucduo
}

// Where appends a list predicates to the UserCareerDescriptionUpdate builder.
func (ucduo *UserCareerDescriptionUpdateOne) Where(ps ...predicate.UserCareerDescription) *UserCareerDescriptionUpdateOne {
	ucduo.mutation.Where(ps...)
	return ucduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucduo *UserCareerDescriptionUpdateOne) Select(field string, fields ...string) *UserCareerDescriptionUpdateOne {
	ucduo.fields = append([]string{field}, fields...)
	return ucduo
}

// Save executes the query and returns the updated UserCareerDescription entity.
func (ucduo *UserCareerDescriptionUpdateOne) Save(ctx context.Context) (*UserCareerDescription, error) {
	return withHooks[*UserCareerDescription, UserCareerDescriptionMutation](ctx, ucduo.sqlSave, ucduo.mutation, ucduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucduo *UserCareerDescriptionUpdateOne) SaveX(ctx context.Context) *UserCareerDescription {
	node, err := ucduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucduo *UserCareerDescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := ucduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucduo *UserCareerDescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := ucduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucduo *UserCareerDescriptionUpdateOne) check() error {
	if v, ok := ucduo.mutation.Description(); ok {
		if err := usercareerdescription.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "UserCareerDescription.description": %w`, err)}
		}
	}
	if _, ok := ucduo.mutation.CareerID(); ucduo.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCareerDescription.career"`)
	}
	return nil
}

func (ucduo *UserCareerDescriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserCareerDescription, err error) {
	if err := ucduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercareerdescription.Table, usercareerdescription.Columns, sqlgraph.NewFieldSpec(usercareerdescription.FieldID, field.TypeInt))
	id, ok := ucduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCareerDescription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercareerdescription.FieldID)
		for _, f := range fields {
			if !usercareerdescription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercareerdescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucduo.mutation.Description(); ok {
		_spec.SetField(usercareerdescription.FieldDescription, field.TypeString, value)
	}
	if ucduo.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareerdescription.CareerTable,
			Columns: []string{usercareerdescription.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucduo.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareerdescription.CareerTable,
			Columns: []string{usercareerdescription.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCareerDescription{config: ucduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercareerdescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucduo.mutation.done = true
	return _node, nil
}
