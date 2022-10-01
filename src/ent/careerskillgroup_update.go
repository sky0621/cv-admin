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
	"github.com/sky0621/cv-admin/src/ent/careerskill"
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
)

// CareerSkillGroupUpdate is the builder for updating CareerSkillGroup entities.
type CareerSkillGroupUpdate struct {
	config
	hooks    []Hook
	mutation *CareerSkillGroupMutation
}

// Where appends a list predicates to the CareerSkillGroupUpdate builder.
func (csgu *CareerSkillGroupUpdate) Where(ps ...predicate.CareerSkillGroup) *CareerSkillGroupUpdate {
	csgu.mutation.Where(ps...)
	return csgu
}

// SetUpdateTime sets the "update_time" field.
func (csgu *CareerSkillGroupUpdate) SetUpdateTime(t time.Time) *CareerSkillGroupUpdate {
	csgu.mutation.SetUpdateTime(t)
	return csgu
}

// SetLabel sets the "label" field.
func (csgu *CareerSkillGroupUpdate) SetLabel(s string) *CareerSkillGroupUpdate {
	csgu.mutation.SetLabel(s)
	return csgu
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (csgu *CareerSkillGroupUpdate) SetCareerID(id int) *CareerSkillGroupUpdate {
	csgu.mutation.SetCareerID(id)
	return csgu
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (csgu *CareerSkillGroupUpdate) SetCareer(u *UserCareer) *CareerSkillGroupUpdate {
	return csgu.SetCareerID(u.ID)
}

// AddCareerSkillIDs adds the "careerSkills" edge to the CareerSkill entity by IDs.
func (csgu *CareerSkillGroupUpdate) AddCareerSkillIDs(ids ...int) *CareerSkillGroupUpdate {
	csgu.mutation.AddCareerSkillIDs(ids...)
	return csgu
}

// AddCareerSkills adds the "careerSkills" edges to the CareerSkill entity.
func (csgu *CareerSkillGroupUpdate) AddCareerSkills(c ...*CareerSkill) *CareerSkillGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csgu.AddCareerSkillIDs(ids...)
}

// Mutation returns the CareerSkillGroupMutation object of the builder.
func (csgu *CareerSkillGroupUpdate) Mutation() *CareerSkillGroupMutation {
	return csgu.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (csgu *CareerSkillGroupUpdate) ClearCareer() *CareerSkillGroupUpdate {
	csgu.mutation.ClearCareer()
	return csgu
}

// ClearCareerSkills clears all "careerSkills" edges to the CareerSkill entity.
func (csgu *CareerSkillGroupUpdate) ClearCareerSkills() *CareerSkillGroupUpdate {
	csgu.mutation.ClearCareerSkills()
	return csgu
}

// RemoveCareerSkillIDs removes the "careerSkills" edge to CareerSkill entities by IDs.
func (csgu *CareerSkillGroupUpdate) RemoveCareerSkillIDs(ids ...int) *CareerSkillGroupUpdate {
	csgu.mutation.RemoveCareerSkillIDs(ids...)
	return csgu
}

// RemoveCareerSkills removes "careerSkills" edges to CareerSkill entities.
func (csgu *CareerSkillGroupUpdate) RemoveCareerSkills(c ...*CareerSkill) *CareerSkillGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csgu.RemoveCareerSkillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csgu *CareerSkillGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	csgu.defaults()
	if len(csgu.hooks) == 0 {
		if err = csgu.check(); err != nil {
			return 0, err
		}
		affected, err = csgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csgu.check(); err != nil {
				return 0, err
			}
			csgu.mutation = mutation
			affected, err = csgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csgu.hooks) - 1; i >= 0; i-- {
			if csgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csgu *CareerSkillGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := csgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csgu *CareerSkillGroupUpdate) Exec(ctx context.Context) error {
	_, err := csgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csgu *CareerSkillGroupUpdate) ExecX(ctx context.Context) {
	if err := csgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csgu *CareerSkillGroupUpdate) defaults() {
	if _, ok := csgu.mutation.UpdateTime(); !ok {
		v := careerskillgroup.UpdateDefaultUpdateTime()
		csgu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csgu *CareerSkillGroupUpdate) check() error {
	if v, ok := csgu.mutation.Label(); ok {
		if err := careerskillgroup.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "CareerSkillGroup.label": %w`, err)}
		}
	}
	if _, ok := csgu.mutation.CareerID(); csgu.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerSkillGroup.career"`)
	}
	return nil
}

func (csgu *CareerSkillGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskillgroup.Table,
			Columns: careerskillgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskillgroup.FieldID,
			},
		},
	}
	if ps := csgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csgu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: careerskillgroup.FieldUpdateTime,
		})
	}
	if value, ok := csgu.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: careerskillgroup.FieldLabel,
		})
	}
	if csgu.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskillgroup.CareerTable,
			Columns: []string{careerskillgroup.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csgu.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskillgroup.CareerTable,
			Columns: []string{careerskillgroup.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csgu.mutation.CareerSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csgu.mutation.RemovedCareerSkillsIDs(); len(nodes) > 0 && !csgu.mutation.CareerSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csgu.mutation.CareerSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careerskillgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CareerSkillGroupUpdateOne is the builder for updating a single CareerSkillGroup entity.
type CareerSkillGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CareerSkillGroupMutation
}

// SetUpdateTime sets the "update_time" field.
func (csguo *CareerSkillGroupUpdateOne) SetUpdateTime(t time.Time) *CareerSkillGroupUpdateOne {
	csguo.mutation.SetUpdateTime(t)
	return csguo
}

// SetLabel sets the "label" field.
func (csguo *CareerSkillGroupUpdateOne) SetLabel(s string) *CareerSkillGroupUpdateOne {
	csguo.mutation.SetLabel(s)
	return csguo
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (csguo *CareerSkillGroupUpdateOne) SetCareerID(id int) *CareerSkillGroupUpdateOne {
	csguo.mutation.SetCareerID(id)
	return csguo
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (csguo *CareerSkillGroupUpdateOne) SetCareer(u *UserCareer) *CareerSkillGroupUpdateOne {
	return csguo.SetCareerID(u.ID)
}

// AddCareerSkillIDs adds the "careerSkills" edge to the CareerSkill entity by IDs.
func (csguo *CareerSkillGroupUpdateOne) AddCareerSkillIDs(ids ...int) *CareerSkillGroupUpdateOne {
	csguo.mutation.AddCareerSkillIDs(ids...)
	return csguo
}

// AddCareerSkills adds the "careerSkills" edges to the CareerSkill entity.
func (csguo *CareerSkillGroupUpdateOne) AddCareerSkills(c ...*CareerSkill) *CareerSkillGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csguo.AddCareerSkillIDs(ids...)
}

// Mutation returns the CareerSkillGroupMutation object of the builder.
func (csguo *CareerSkillGroupUpdateOne) Mutation() *CareerSkillGroupMutation {
	return csguo.mutation
}

// ClearCareer clears the "career" edge to the UserCareer entity.
func (csguo *CareerSkillGroupUpdateOne) ClearCareer() *CareerSkillGroupUpdateOne {
	csguo.mutation.ClearCareer()
	return csguo
}

// ClearCareerSkills clears all "careerSkills" edges to the CareerSkill entity.
func (csguo *CareerSkillGroupUpdateOne) ClearCareerSkills() *CareerSkillGroupUpdateOne {
	csguo.mutation.ClearCareerSkills()
	return csguo
}

// RemoveCareerSkillIDs removes the "careerSkills" edge to CareerSkill entities by IDs.
func (csguo *CareerSkillGroupUpdateOne) RemoveCareerSkillIDs(ids ...int) *CareerSkillGroupUpdateOne {
	csguo.mutation.RemoveCareerSkillIDs(ids...)
	return csguo
}

// RemoveCareerSkills removes "careerSkills" edges to CareerSkill entities.
func (csguo *CareerSkillGroupUpdateOne) RemoveCareerSkills(c ...*CareerSkill) *CareerSkillGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csguo.RemoveCareerSkillIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csguo *CareerSkillGroupUpdateOne) Select(field string, fields ...string) *CareerSkillGroupUpdateOne {
	csguo.fields = append([]string{field}, fields...)
	return csguo
}

// Save executes the query and returns the updated CareerSkillGroup entity.
func (csguo *CareerSkillGroupUpdateOne) Save(ctx context.Context) (*CareerSkillGroup, error) {
	var (
		err  error
		node *CareerSkillGroup
	)
	csguo.defaults()
	if len(csguo.hooks) == 0 {
		if err = csguo.check(); err != nil {
			return nil, err
		}
		node, err = csguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csguo.check(); err != nil {
				return nil, err
			}
			csguo.mutation = mutation
			node, err = csguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csguo.hooks) - 1; i >= 0; i-- {
			if csguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, csguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CareerSkillGroup)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CareerSkillGroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csguo *CareerSkillGroupUpdateOne) SaveX(ctx context.Context) *CareerSkillGroup {
	node, err := csguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csguo *CareerSkillGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := csguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csguo *CareerSkillGroupUpdateOne) ExecX(ctx context.Context) {
	if err := csguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csguo *CareerSkillGroupUpdateOne) defaults() {
	if _, ok := csguo.mutation.UpdateTime(); !ok {
		v := careerskillgroup.UpdateDefaultUpdateTime()
		csguo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csguo *CareerSkillGroupUpdateOne) check() error {
	if v, ok := csguo.mutation.Label(); ok {
		if err := careerskillgroup.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "CareerSkillGroup.label": %w`, err)}
		}
	}
	if _, ok := csguo.mutation.CareerID(); csguo.mutation.CareerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerSkillGroup.career"`)
	}
	return nil
}

func (csguo *CareerSkillGroupUpdateOne) sqlSave(ctx context.Context) (_node *CareerSkillGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskillgroup.Table,
			Columns: careerskillgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskillgroup.FieldID,
			},
		},
	}
	id, ok := csguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CareerSkillGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careerskillgroup.FieldID)
		for _, f := range fields {
			if !careerskillgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != careerskillgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csguo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: careerskillgroup.FieldUpdateTime,
		})
	}
	if value, ok := csguo.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: careerskillgroup.FieldLabel,
		})
	}
	if csguo.mutation.CareerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskillgroup.CareerTable,
			Columns: []string{careerskillgroup.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csguo.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskillgroup.CareerTable,
			Columns: []string{careerskillgroup.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csguo.mutation.CareerSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csguo.mutation.RemovedCareerSkillsIDs(); len(nodes) > 0 && !csguo.mutation.CareerSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csguo.mutation.CareerSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careerskillgroup.CareerSkillsTable,
			Columns: []string{careerskillgroup.CareerSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CareerSkillGroup{config: csguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careerskillgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
