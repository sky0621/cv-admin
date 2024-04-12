// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// UserCareerGroupUpdate is the builder for updating UserCareerGroup entities.
type UserCareerGroupUpdate struct {
	config
	hooks    []Hook
	mutation *UserCareerGroupMutation
}

// Where appends a list predicates to the UserCareerGroupUpdate builder.
func (ucgu *UserCareerGroupUpdate) Where(ps ...predicate.UserCareerGroup) *UserCareerGroupUpdate {
	ucgu.mutation.Where(ps...)
	return ucgu
}

// SetUpdateTime sets the "update_time" field.
func (ucgu *UserCareerGroupUpdate) SetUpdateTime(t time.Time) *UserCareerGroupUpdate {
	ucgu.mutation.SetUpdateTime(t)
	return ucgu
}

// SetLabel sets the "label" field.
func (ucgu *UserCareerGroupUpdate) SetLabel(s string) *UserCareerGroupUpdate {
	ucgu.mutation.SetLabel(s)
	return ucgu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (ucgu *UserCareerGroupUpdate) SetNillableLabel(s *string) *UserCareerGroupUpdate {
	if s != nil {
		ucgu.SetLabel(*s)
	}
	return ucgu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ucgu *UserCareerGroupUpdate) SetUserID(id int) *UserCareerGroupUpdate {
	ucgu.mutation.SetUserID(id)
	return ucgu
}

// SetUser sets the "user" edge to the User entity.
func (ucgu *UserCareerGroupUpdate) SetUser(u *User) *UserCareerGroupUpdate {
	return ucgu.SetUserID(u.ID)
}

// AddCareerIDs adds the "careers" edge to the UserCareer entity by IDs.
func (ucgu *UserCareerGroupUpdate) AddCareerIDs(ids ...int) *UserCareerGroupUpdate {
	ucgu.mutation.AddCareerIDs(ids...)
	return ucgu
}

// AddCareers adds the "careers" edges to the UserCareer entity.
func (ucgu *UserCareerGroupUpdate) AddCareers(u ...*UserCareer) *UserCareerGroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucgu.AddCareerIDs(ids...)
}

// Mutation returns the UserCareerGroupMutation object of the builder.
func (ucgu *UserCareerGroupUpdate) Mutation() *UserCareerGroupMutation {
	return ucgu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucgu *UserCareerGroupUpdate) ClearUser() *UserCareerGroupUpdate {
	ucgu.mutation.ClearUser()
	return ucgu
}

// ClearCareers clears all "careers" edges to the UserCareer entity.
func (ucgu *UserCareerGroupUpdate) ClearCareers() *UserCareerGroupUpdate {
	ucgu.mutation.ClearCareers()
	return ucgu
}

// RemoveCareerIDs removes the "careers" edge to UserCareer entities by IDs.
func (ucgu *UserCareerGroupUpdate) RemoveCareerIDs(ids ...int) *UserCareerGroupUpdate {
	ucgu.mutation.RemoveCareerIDs(ids...)
	return ucgu
}

// RemoveCareers removes "careers" edges to UserCareer entities.
func (ucgu *UserCareerGroupUpdate) RemoveCareers(u ...*UserCareer) *UserCareerGroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucgu.RemoveCareerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucgu *UserCareerGroupUpdate) Save(ctx context.Context) (int, error) {
	ucgu.defaults()
	return withHooks(ctx, ucgu.sqlSave, ucgu.mutation, ucgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucgu *UserCareerGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ucgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucgu *UserCareerGroupUpdate) Exec(ctx context.Context) error {
	_, err := ucgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucgu *UserCareerGroupUpdate) ExecX(ctx context.Context) {
	if err := ucgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucgu *UserCareerGroupUpdate) defaults() {
	if _, ok := ucgu.mutation.UpdateTime(); !ok {
		v := usercareergroup.UpdateDefaultUpdateTime()
		ucgu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucgu *UserCareerGroupUpdate) check() error {
	if v, ok := ucgu.mutation.Label(); ok {
		if err := usercareergroup.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "UserCareerGroup.label": %w`, err)}
		}
	}
	if _, ok := ucgu.mutation.UserID(); ucgu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCareerGroup.user"`)
	}
	return nil
}

func (ucgu *UserCareerGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucgu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercareergroup.Table, usercareergroup.Columns, sqlgraph.NewFieldSpec(usercareergroup.FieldID, field.TypeInt))
	if ps := ucgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucgu.mutation.UpdateTime(); ok {
		_spec.SetField(usercareergroup.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ucgu.mutation.Label(); ok {
		_spec.SetField(usercareergroup.FieldLabel, field.TypeString, value)
	}
	if ucgu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareergroup.UserTable,
			Columns: []string{usercareergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucgu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareergroup.UserTable,
			Columns: []string{usercareergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucgu.mutation.CareersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucgu.mutation.RemovedCareersIDs(); len(nodes) > 0 && !ucgu.mutation.CareersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucgu.mutation.CareersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ucgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercareergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucgu.mutation.done = true
	return n, nil
}

// UserCareerGroupUpdateOne is the builder for updating a single UserCareerGroup entity.
type UserCareerGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCareerGroupMutation
}

// SetUpdateTime sets the "update_time" field.
func (ucguo *UserCareerGroupUpdateOne) SetUpdateTime(t time.Time) *UserCareerGroupUpdateOne {
	ucguo.mutation.SetUpdateTime(t)
	return ucguo
}

// SetLabel sets the "label" field.
func (ucguo *UserCareerGroupUpdateOne) SetLabel(s string) *UserCareerGroupUpdateOne {
	ucguo.mutation.SetLabel(s)
	return ucguo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (ucguo *UserCareerGroupUpdateOne) SetNillableLabel(s *string) *UserCareerGroupUpdateOne {
	if s != nil {
		ucguo.SetLabel(*s)
	}
	return ucguo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ucguo *UserCareerGroupUpdateOne) SetUserID(id int) *UserCareerGroupUpdateOne {
	ucguo.mutation.SetUserID(id)
	return ucguo
}

// SetUser sets the "user" edge to the User entity.
func (ucguo *UserCareerGroupUpdateOne) SetUser(u *User) *UserCareerGroupUpdateOne {
	return ucguo.SetUserID(u.ID)
}

// AddCareerIDs adds the "careers" edge to the UserCareer entity by IDs.
func (ucguo *UserCareerGroupUpdateOne) AddCareerIDs(ids ...int) *UserCareerGroupUpdateOne {
	ucguo.mutation.AddCareerIDs(ids...)
	return ucguo
}

// AddCareers adds the "careers" edges to the UserCareer entity.
func (ucguo *UserCareerGroupUpdateOne) AddCareers(u ...*UserCareer) *UserCareerGroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucguo.AddCareerIDs(ids...)
}

// Mutation returns the UserCareerGroupMutation object of the builder.
func (ucguo *UserCareerGroupUpdateOne) Mutation() *UserCareerGroupMutation {
	return ucguo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucguo *UserCareerGroupUpdateOne) ClearUser() *UserCareerGroupUpdateOne {
	ucguo.mutation.ClearUser()
	return ucguo
}

// ClearCareers clears all "careers" edges to the UserCareer entity.
func (ucguo *UserCareerGroupUpdateOne) ClearCareers() *UserCareerGroupUpdateOne {
	ucguo.mutation.ClearCareers()
	return ucguo
}

// RemoveCareerIDs removes the "careers" edge to UserCareer entities by IDs.
func (ucguo *UserCareerGroupUpdateOne) RemoveCareerIDs(ids ...int) *UserCareerGroupUpdateOne {
	ucguo.mutation.RemoveCareerIDs(ids...)
	return ucguo
}

// RemoveCareers removes "careers" edges to UserCareer entities.
func (ucguo *UserCareerGroupUpdateOne) RemoveCareers(u ...*UserCareer) *UserCareerGroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucguo.RemoveCareerIDs(ids...)
}

// Where appends a list predicates to the UserCareerGroupUpdate builder.
func (ucguo *UserCareerGroupUpdateOne) Where(ps ...predicate.UserCareerGroup) *UserCareerGroupUpdateOne {
	ucguo.mutation.Where(ps...)
	return ucguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucguo *UserCareerGroupUpdateOne) Select(field string, fields ...string) *UserCareerGroupUpdateOne {
	ucguo.fields = append([]string{field}, fields...)
	return ucguo
}

// Save executes the query and returns the updated UserCareerGroup entity.
func (ucguo *UserCareerGroupUpdateOne) Save(ctx context.Context) (*UserCareerGroup, error) {
	ucguo.defaults()
	return withHooks(ctx, ucguo.sqlSave, ucguo.mutation, ucguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucguo *UserCareerGroupUpdateOne) SaveX(ctx context.Context) *UserCareerGroup {
	node, err := ucguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucguo *UserCareerGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := ucguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucguo *UserCareerGroupUpdateOne) ExecX(ctx context.Context) {
	if err := ucguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucguo *UserCareerGroupUpdateOne) defaults() {
	if _, ok := ucguo.mutation.UpdateTime(); !ok {
		v := usercareergroup.UpdateDefaultUpdateTime()
		ucguo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucguo *UserCareerGroupUpdateOne) check() error {
	if v, ok := ucguo.mutation.Label(); ok {
		if err := usercareergroup.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "UserCareerGroup.label": %w`, err)}
		}
	}
	if _, ok := ucguo.mutation.UserID(); ucguo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCareerGroup.user"`)
	}
	return nil
}

func (ucguo *UserCareerGroupUpdateOne) sqlSave(ctx context.Context) (_node *UserCareerGroup, err error) {
	if err := ucguo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercareergroup.Table, usercareergroup.Columns, sqlgraph.NewFieldSpec(usercareergroup.FieldID, field.TypeInt))
	id, ok := ucguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCareerGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercareergroup.FieldID)
		for _, f := range fields {
			if !usercareergroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercareergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucguo.mutation.UpdateTime(); ok {
		_spec.SetField(usercareergroup.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ucguo.mutation.Label(); ok {
		_spec.SetField(usercareergroup.FieldLabel, field.TypeString, value)
	}
	if ucguo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareergroup.UserTable,
			Columns: []string{usercareergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucguo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareergroup.UserTable,
			Columns: []string{usercareergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucguo.mutation.CareersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucguo.mutation.RemovedCareersIDs(); len(nodes) > 0 && !ucguo.mutation.CareersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucguo.mutation.CareersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareergroup.CareersTable,
			Columns: []string{usercareergroup.CareersColumn},
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
	_node = &UserCareerGroup{config: ucguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercareergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucguo.mutation.done = true
	return _node, nil
}
