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
)

// CareerSkillUpdate is the builder for updating CareerSkill entities.
type CareerSkillUpdate struct {
	config
	hooks    []Hook
	mutation *CareerSkillMutation
}

// Where appends a list predicates to the CareerSkillUpdate builder.
func (csu *CareerSkillUpdate) Where(ps ...predicate.CareerSkill) *CareerSkillUpdate {
	csu.mutation.Where(ps...)
	return csu
}

// SetUpdateTime sets the "update_time" field.
func (csu *CareerSkillUpdate) SetUpdateTime(t time.Time) *CareerSkillUpdate {
	csu.mutation.SetUpdateTime(t)
	return csu
}

// SetName sets the "name" field.
func (csu *CareerSkillUpdate) SetName(s string) *CareerSkillUpdate {
	csu.mutation.SetName(s)
	return csu
}

// SetURL sets the "url" field.
func (csu *CareerSkillUpdate) SetURL(s string) *CareerSkillUpdate {
	csu.mutation.SetURL(s)
	return csu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (csu *CareerSkillUpdate) SetNillableURL(s *string) *CareerSkillUpdate {
	if s != nil {
		csu.SetURL(*s)
	}
	return csu
}

// ClearURL clears the value of the "url" field.
func (csu *CareerSkillUpdate) ClearURL() *CareerSkillUpdate {
	csu.mutation.ClearURL()
	return csu
}

// SetVersion sets the "version" field.
func (csu *CareerSkillUpdate) SetVersion(s string) *CareerSkillUpdate {
	csu.mutation.SetVersion(s)
	return csu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (csu *CareerSkillUpdate) SetNillableVersion(s *string) *CareerSkillUpdate {
	if s != nil {
		csu.SetVersion(*s)
	}
	return csu
}

// ClearVersion clears the value of the "version" field.
func (csu *CareerSkillUpdate) ClearVersion() *CareerSkillUpdate {
	csu.mutation.ClearVersion()
	return csu
}

// SetCareerSkillGroupID sets the "careerSkillGroup" edge to the CareerSkillGroup entity by ID.
func (csu *CareerSkillUpdate) SetCareerSkillGroupID(id int) *CareerSkillUpdate {
	csu.mutation.SetCareerSkillGroupID(id)
	return csu
}

// SetCareerSkillGroup sets the "careerSkillGroup" edge to the CareerSkillGroup entity.
func (csu *CareerSkillUpdate) SetCareerSkillGroup(c *CareerSkillGroup) *CareerSkillUpdate {
	return csu.SetCareerSkillGroupID(c.ID)
}

// Mutation returns the CareerSkillMutation object of the builder.
func (csu *CareerSkillUpdate) Mutation() *CareerSkillMutation {
	return csu.mutation
}

// ClearCareerSkillGroup clears the "careerSkillGroup" edge to the CareerSkillGroup entity.
func (csu *CareerSkillUpdate) ClearCareerSkillGroup() *CareerSkillUpdate {
	csu.mutation.ClearCareerSkillGroup()
	return csu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *CareerSkillUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	csu.defaults()
	if len(csu.hooks) == 0 {
		if err = csu.check(); err != nil {
			return 0, err
		}
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csu.check(); err != nil {
				return 0, err
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			if csu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *CareerSkillUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *CareerSkillUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *CareerSkillUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csu *CareerSkillUpdate) defaults() {
	if _, ok := csu.mutation.UpdateTime(); !ok {
		v := careerskill.UpdateDefaultUpdateTime()
		csu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csu *CareerSkillUpdate) check() error {
	if v, ok := csu.mutation.Name(); ok {
		if err := careerskill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.name": %w`, err)}
		}
	}
	if v, ok := csu.mutation.URL(); ok {
		if err := careerskill.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.url": %w`, err)}
		}
	}
	if v, ok := csu.mutation.Version(); ok {
		if err := careerskill.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.version": %w`, err)}
		}
	}
	if _, ok := csu.mutation.CareerSkillGroupID(); csu.mutation.CareerSkillGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerSkill.careerSkillGroup"`)
	}
	return nil
}

func (csu *CareerSkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskill.Table,
			Columns: careerskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskill.FieldID,
			},
		},
	}
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.UpdateTime(); ok {
		_spec.SetField(careerskill.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := csu.mutation.Name(); ok {
		_spec.SetField(careerskill.FieldName, field.TypeString, value)
	}
	if value, ok := csu.mutation.URL(); ok {
		_spec.SetField(careerskill.FieldURL, field.TypeString, value)
	}
	if csu.mutation.URLCleared() {
		_spec.ClearField(careerskill.FieldURL, field.TypeString)
	}
	if value, ok := csu.mutation.Version(); ok {
		_spec.SetField(careerskill.FieldVersion, field.TypeString, value)
	}
	if csu.mutation.VersionCleared() {
		_spec.ClearField(careerskill.FieldVersion, field.TypeString)
	}
	if csu.mutation.CareerSkillGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskill.CareerSkillGroupTable,
			Columns: []string{careerskill.CareerSkillGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskillgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.CareerSkillGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskill.CareerSkillGroupTable,
			Columns: []string{careerskill.CareerSkillGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskillgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careerskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CareerSkillUpdateOne is the builder for updating a single CareerSkill entity.
type CareerSkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CareerSkillMutation
}

// SetUpdateTime sets the "update_time" field.
func (csuo *CareerSkillUpdateOne) SetUpdateTime(t time.Time) *CareerSkillUpdateOne {
	csuo.mutation.SetUpdateTime(t)
	return csuo
}

// SetName sets the "name" field.
func (csuo *CareerSkillUpdateOne) SetName(s string) *CareerSkillUpdateOne {
	csuo.mutation.SetName(s)
	return csuo
}

// SetURL sets the "url" field.
func (csuo *CareerSkillUpdateOne) SetURL(s string) *CareerSkillUpdateOne {
	csuo.mutation.SetURL(s)
	return csuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (csuo *CareerSkillUpdateOne) SetNillableURL(s *string) *CareerSkillUpdateOne {
	if s != nil {
		csuo.SetURL(*s)
	}
	return csuo
}

// ClearURL clears the value of the "url" field.
func (csuo *CareerSkillUpdateOne) ClearURL() *CareerSkillUpdateOne {
	csuo.mutation.ClearURL()
	return csuo
}

// SetVersion sets the "version" field.
func (csuo *CareerSkillUpdateOne) SetVersion(s string) *CareerSkillUpdateOne {
	csuo.mutation.SetVersion(s)
	return csuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (csuo *CareerSkillUpdateOne) SetNillableVersion(s *string) *CareerSkillUpdateOne {
	if s != nil {
		csuo.SetVersion(*s)
	}
	return csuo
}

// ClearVersion clears the value of the "version" field.
func (csuo *CareerSkillUpdateOne) ClearVersion() *CareerSkillUpdateOne {
	csuo.mutation.ClearVersion()
	return csuo
}

// SetCareerSkillGroupID sets the "careerSkillGroup" edge to the CareerSkillGroup entity by ID.
func (csuo *CareerSkillUpdateOne) SetCareerSkillGroupID(id int) *CareerSkillUpdateOne {
	csuo.mutation.SetCareerSkillGroupID(id)
	return csuo
}

// SetCareerSkillGroup sets the "careerSkillGroup" edge to the CareerSkillGroup entity.
func (csuo *CareerSkillUpdateOne) SetCareerSkillGroup(c *CareerSkillGroup) *CareerSkillUpdateOne {
	return csuo.SetCareerSkillGroupID(c.ID)
}

// Mutation returns the CareerSkillMutation object of the builder.
func (csuo *CareerSkillUpdateOne) Mutation() *CareerSkillMutation {
	return csuo.mutation
}

// ClearCareerSkillGroup clears the "careerSkillGroup" edge to the CareerSkillGroup entity.
func (csuo *CareerSkillUpdateOne) ClearCareerSkillGroup() *CareerSkillUpdateOne {
	csuo.mutation.ClearCareerSkillGroup()
	return csuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csuo *CareerSkillUpdateOne) Select(field string, fields ...string) *CareerSkillUpdateOne {
	csuo.fields = append([]string{field}, fields...)
	return csuo
}

// Save executes the query and returns the updated CareerSkill entity.
func (csuo *CareerSkillUpdateOne) Save(ctx context.Context) (*CareerSkill, error) {
	var (
		err  error
		node *CareerSkill
	)
	csuo.defaults()
	if len(csuo.hooks) == 0 {
		if err = csuo.check(); err != nil {
			return nil, err
		}
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csuo.check(); err != nil {
				return nil, err
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			if csuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, csuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CareerSkill)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CareerSkillMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *CareerSkillUpdateOne) SaveX(ctx context.Context) *CareerSkill {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *CareerSkillUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *CareerSkillUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csuo *CareerSkillUpdateOne) defaults() {
	if _, ok := csuo.mutation.UpdateTime(); !ok {
		v := careerskill.UpdateDefaultUpdateTime()
		csuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csuo *CareerSkillUpdateOne) check() error {
	if v, ok := csuo.mutation.Name(); ok {
		if err := careerskill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.name": %w`, err)}
		}
	}
	if v, ok := csuo.mutation.URL(); ok {
		if err := careerskill.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.url": %w`, err)}
		}
	}
	if v, ok := csuo.mutation.Version(); ok {
		if err := careerskill.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.version": %w`, err)}
		}
	}
	if _, ok := csuo.mutation.CareerSkillGroupID(); csuo.mutation.CareerSkillGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CareerSkill.careerSkillGroup"`)
	}
	return nil
}

func (csuo *CareerSkillUpdateOne) sqlSave(ctx context.Context) (_node *CareerSkill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   careerskill.Table,
			Columns: careerskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskill.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CareerSkill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, careerskill.FieldID)
		for _, f := range fields {
			if !careerskill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != careerskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.UpdateTime(); ok {
		_spec.SetField(careerskill.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := csuo.mutation.Name(); ok {
		_spec.SetField(careerskill.FieldName, field.TypeString, value)
	}
	if value, ok := csuo.mutation.URL(); ok {
		_spec.SetField(careerskill.FieldURL, field.TypeString, value)
	}
	if csuo.mutation.URLCleared() {
		_spec.ClearField(careerskill.FieldURL, field.TypeString)
	}
	if value, ok := csuo.mutation.Version(); ok {
		_spec.SetField(careerskill.FieldVersion, field.TypeString, value)
	}
	if csuo.mutation.VersionCleared() {
		_spec.ClearField(careerskill.FieldVersion, field.TypeString)
	}
	if csuo.mutation.CareerSkillGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskill.CareerSkillGroupTable,
			Columns: []string{careerskill.CareerSkillGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskillgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.CareerSkillGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careerskill.CareerSkillGroupTable,
			Columns: []string{careerskill.CareerSkillGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careerskillgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CareerSkill{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{careerskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
