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
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/careertaskdescription"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
)

// CareerTaskUpdate is the builder for updating CareerTask entities.
type CareerTaskUpdate struct {
	config
	hooks    []Hook
	mutation *CareerTaskMutation
}

// Where appends a list predicates to the CareerTaskUpdate builder.
func (ctu *CareerTaskUpdate) Where(ps ...predicate.CareerTask) *CareerTaskUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetUpdateTime sets the "update_time" field.
func (ctu *CareerTaskUpdate) SetUpdateTime(t time.Time) *CareerTaskUpdate {
	ctu.mutation.SetUpdateTime(t)
	return ctu
}

// SetName sets the "name" field.
func (ctu *CareerTaskUpdate) SetName(s string) *CareerTaskUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ctu *CareerTaskUpdate) SetCareerID(id int) *CareerTaskUpdate {
	ctu.mutation.SetCareerID(id)
	return ctu
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ctu *CareerTaskUpdate) SetCareer(u *UserCareer) *CareerTaskUpdate {
	return ctu.SetCareerID(u.ID)
}

// AddCareerTaskDescriptionIDs adds the "careerTaskDescriptions" edge to the CareerTaskDescription entity by IDs.
func (ctu *CareerTaskUpdate) AddCareerTaskDescriptionIDs(ids ...int) *CareerTaskUpdate {
	ctu.mutation.AddCareerTaskDescriptionIDs(ids...)
	return ctu
}

// AddCareerTaskDescriptions adds the "careerTaskDescriptions" edges to the CareerTaskDescription entity.
func (ctu *CareerTaskUpdate) AddCareerTaskDescriptions(c ...*CareerTaskDescription) *CareerTaskUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddCareerTaskDescriptionIDs(ids...)
}

// Mutation returns the CareerTaskMutation object of the builder.
func (ctu *CareerTaskUpdate) Mutation() *CareerTaskMutation {
	return ctu.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (ctu *CareerTaskUpdate) ClearCareer() *CareerTaskUpdate {
	ctu.mutation.ClearCareer()
	return ctu
}

// ClearCareerTaskDescriptions clears all "careerTaskDescriptions" edges to the CareerTaskDescription entity.
func (ctu *CareerTaskUpdate) ClearCareerTaskDescriptions() *CareerTaskUpdate {
	ctu.mutation.ClearCareerTaskDescriptions()
	return ctu
}

// RemoveCareerTaskDescriptionIDs removes the "careerTaskDescriptions" edge to CareerTaskDescription entities by IDs.
func (ctu *CareerTaskUpdate) RemoveCareerTaskDescriptionIDs(ids ...int) *CareerTaskUpdate {
	ctu.mutation.RemoveCareerTaskDescriptionIDs(ids...)
	return ctu
}

// RemoveCareerTaskDescriptions removes "careerTaskDescriptions" edges to CareerTaskDescription entities.
func (ctu *CareerTaskUpdate) RemoveCareerTaskDescriptions(c ...*CareerTaskDescription) *CareerTaskUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveCareerTaskDescriptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CareerTaskUpdate) Save(ctx context.Context) (int, error) {
	ctu.defaults()
	return withHooks(ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CareerTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CareerTaskUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CareerTaskUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CareerTaskUpdate) defaults() {
	if _, ok := ctu.mutation.UpdateTime(); !ok {
		v := careertask.UpdateDefaultUpdateTime()
		ctu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *CareerTaskUpdate) check() error {
	if v, ok := ctu.mutation.Name(); ok {
		if err := careertask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerTask.name": %w`, err)}
		}
	}
	if _, ok := ctu.mutation.CareerID(); ctu.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerTask.career"`)
	}
	return nil
}

func (ctu *CareerTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(careertask.Table, careertask.Columns, sqlgraph.NewFieldSpec(careertask.FieldID, field.TypeInt))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.UpdateTime(); ok {
		_spec.SetField(careertask.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.SetField(careertask.FieldName, field.TypeString, value)
	}
	if ctu.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careertask.CareerTable,
			Columns: []string{careertask.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careertask.CareerTable,
			Columns: []string{careertask.CareerColumn},
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
	if ctu.mutation.CareerTaskDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedCareerTaskDescriptionsIDs(); len(nodes) > 0 && !ctu.mutation.CareerTaskDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CareerTaskDescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careertask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CareerTaskUpdateOne is the builder for updating a single CareerTask entity.
type CareerTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CareerTaskMutation
}

// SetUpdateTime sets the "update_time" field.
func (ctuo *CareerTaskUpdateOne) SetUpdateTime(t time.Time) *CareerTaskUpdateOne {
	ctuo.mutation.SetUpdateTime(t)
	return ctuo
}

// SetName sets the "name" field.
func (ctuo *CareerTaskUpdateOne) SetName(s string) *CareerTaskUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ctuo *CareerTaskUpdateOne) SetCareerID(id int) *CareerTaskUpdateOne {
	ctuo.mutation.SetCareerID(id)
	return ctuo
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ctuo *CareerTaskUpdateOne) SetCareer(u *UserCareer) *CareerTaskUpdateOne {
	return ctuo.SetCareerID(u.ID)
}

// AddCareerTaskDescriptionIDs adds the "careerTaskDescriptions" edge to the CareerTaskDescription entity by IDs.
func (ctuo *CareerTaskUpdateOne) AddCareerTaskDescriptionIDs(ids ...int) *CareerTaskUpdateOne {
	ctuo.mutation.AddCareerTaskDescriptionIDs(ids...)
	return ctuo
}

// AddCareerTaskDescriptions adds the "careerTaskDescriptions" edges to the CareerTaskDescription entity.
func (ctuo *CareerTaskUpdateOne) AddCareerTaskDescriptions(c ...*CareerTaskDescription) *CareerTaskUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddCareerTaskDescriptionIDs(ids...)
}

// Mutation returns the CareerTaskMutation object of the builder.
func (ctuo *CareerTaskUpdateOne) Mutation() *CareerTaskMutation {
	return ctuo.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (ctuo *CareerTaskUpdateOne) ClearCareer() *CareerTaskUpdateOne {
	ctuo.mutation.ClearCareer()
	return ctuo
}

// ClearCareerTaskDescriptions clears all "careerTaskDescriptions" edges to the CareerTaskDescription entity.
func (ctuo *CareerTaskUpdateOne) ClearCareerTaskDescriptions() *CareerTaskUpdateOne {
	ctuo.mutation.ClearCareerTaskDescriptions()
	return ctuo
}

// RemoveCareerTaskDescriptionIDs removes the "careerTaskDescriptions" edge to CareerTaskDescription entities by IDs.
func (ctuo *CareerTaskUpdateOne) RemoveCareerTaskDescriptionIDs(ids ...int) *CareerTaskUpdateOne {
	ctuo.mutation.RemoveCareerTaskDescriptionIDs(ids...)
	return ctuo
}

// RemoveCareerTaskDescriptions removes "careerTaskDescriptions" edges to CareerTaskDescription entities.
func (ctuo *CareerTaskUpdateOne) RemoveCareerTaskDescriptions(c ...*CareerTaskDescription) *CareerTaskUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveCareerTaskDescriptionIDs(ids...)
}

// Where appends a list predicates to the CareerTaskUpdate builder.
func (ctuo *CareerTaskUpdateOne) Where(ps ...predicate.CareerTask) *CareerTaskUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CareerTaskUpdateOne) Select(field string, fields ...string) *CareerTaskUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CareerTask entity.
func (ctuo *CareerTaskUpdateOne) Save(ctx context.Context) (*CareerTask, error) {
	ctuo.defaults()
	return withHooks(ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CareerTaskUpdateOne) SaveX(ctx context.Context) *CareerTask {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CareerTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CareerTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CareerTaskUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdateTime(); !ok {
		v := careertask.UpdateDefaultUpdateTime()
		ctuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *CareerTaskUpdateOne) check() error {
	if v, ok := ctuo.mutation.Name(); ok {
		if err := careertask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerTask.name": %w`, err)}
		}
	}
	if _, ok := ctuo.mutation.CareerID(); ctuo.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerTask.career"`)
	}
	return nil
}

func (ctuo *CareerTaskUpdateOne) sqlSave(ctx context.Context) (_node *CareerTask, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(careertask.Table, careertask.Columns, sqlgraph.NewFieldSpec(careertask.FieldID, field.TypeInt))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CareerTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careertask.FieldID)
		for _, f := range fields {
			if !careertask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != careertask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.UpdateTime(); ok {
		_spec.SetField(careertask.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.SetField(careertask.FieldName, field.TypeString, value)
	}
	if ctuo.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careertask.CareerTable,
			Columns: []string{careertask.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careertask.CareerTable,
			Columns: []string{careertask.CareerColumn},
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
	if ctuo.mutation.CareerTaskDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedCareerTaskDescriptionsIDs(); len(nodes) > 0 && !ctuo.mutation.CareerTaskDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CareerTaskDescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareerTaskDescriptionsTable,
			Columns: []string{careertask.CareerTaskDescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careertaskdescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CareerTask{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careertask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
