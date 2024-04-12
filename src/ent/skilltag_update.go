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
	"github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/skilltag"
)

// SkillTagUpdate is the builder for updating SkillTag entities.
type SkillTagUpdate struct {
	config
	hooks    []Hook
	mutation *SkillTagMutation
}

// Where appends a list predicates to the SkillTagUpdate builder.
func (stu *SkillTagUpdate) Where(ps ...predicate.SkillTag) *SkillTagUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetName sets the "name" field.
func (stu *SkillTagUpdate) SetName(s string) *SkillTagUpdate {
	stu.mutation.SetName(s)
	return stu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (stu *SkillTagUpdate) SetNillableName(s *string) *SkillTagUpdate {
	if s != nil {
		stu.SetName(*s)
	}
	return stu
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (stu *SkillTagUpdate) AddSkillIDs(ids ...int) *SkillTagUpdate {
	stu.mutation.AddSkillIDs(ids...)
	return stu
}

// AddSkills adds the "skills" edges to the Skill entity.
func (stu *SkillTagUpdate) AddSkills(s ...*Skill) *SkillTagUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddSkillIDs(ids...)
}

// Mutation returns the SkillTagMutation object of the builder.
func (stu *SkillTagUpdate) Mutation() *SkillTagMutation {
	return stu.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (stu *SkillTagUpdate) ClearSkills() *SkillTagUpdate {
	stu.mutation.ClearSkills()
	return stu
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (stu *SkillTagUpdate) RemoveSkillIDs(ids ...int) *SkillTagUpdate {
	stu.mutation.RemoveSkillIDs(ids...)
	return stu
}

// RemoveSkills removes "skills" edges to Skill entities.
func (stu *SkillTagUpdate) RemoveSkills(s ...*Skill) *SkillTagUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveSkillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SkillTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, stu.sqlSave, stu.mutation, stu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SkillTagUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SkillTagUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SkillTagUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *SkillTagUpdate) check() error {
	if v, ok := stu.mutation.Name(); ok {
		if err := skilltag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillTag.name": %w`, err)}
		}
	}
	return nil
}

func (stu *SkillTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := stu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(skilltag.Table, skilltag.Columns, sqlgraph.NewFieldSpec(skilltag.FieldID, field.TypeInt))
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.Name(); ok {
		_spec.SetField(skilltag.FieldName, field.TypeString, value)
	}
	if stu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !stu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	stu.mutation.done = true
	return n, nil
}

// SkillTagUpdateOne is the builder for updating a single SkillTag entity.
type SkillTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillTagMutation
}

// SetName sets the "name" field.
func (stuo *SkillTagUpdateOne) SetName(s string) *SkillTagUpdateOne {
	stuo.mutation.SetName(s)
	return stuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (stuo *SkillTagUpdateOne) SetNillableName(s *string) *SkillTagUpdateOne {
	if s != nil {
		stuo.SetName(*s)
	}
	return stuo
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (stuo *SkillTagUpdateOne) AddSkillIDs(ids ...int) *SkillTagUpdateOne {
	stuo.mutation.AddSkillIDs(ids...)
	return stuo
}

// AddSkills adds the "skills" edges to the Skill entity.
func (stuo *SkillTagUpdateOne) AddSkills(s ...*Skill) *SkillTagUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddSkillIDs(ids...)
}

// Mutation returns the SkillTagMutation object of the builder.
func (stuo *SkillTagUpdateOne) Mutation() *SkillTagMutation {
	return stuo.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (stuo *SkillTagUpdateOne) ClearSkills() *SkillTagUpdateOne {
	stuo.mutation.ClearSkills()
	return stuo
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (stuo *SkillTagUpdateOne) RemoveSkillIDs(ids ...int) *SkillTagUpdateOne {
	stuo.mutation.RemoveSkillIDs(ids...)
	return stuo
}

// RemoveSkills removes "skills" edges to Skill entities.
func (stuo *SkillTagUpdateOne) RemoveSkills(s ...*Skill) *SkillTagUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveSkillIDs(ids...)
}

// Where appends a list predicates to the SkillTagUpdate builder.
func (stuo *SkillTagUpdateOne) Where(ps ...predicate.SkillTag) *SkillTagUpdateOne {
	stuo.mutation.Where(ps...)
	return stuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SkillTagUpdateOne) Select(field string, fields ...string) *SkillTagUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SkillTag entity.
func (stuo *SkillTagUpdateOne) Save(ctx context.Context) (*SkillTag, error) {
	return withHooks(ctx, stuo.sqlSave, stuo.mutation, stuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SkillTagUpdateOne) SaveX(ctx context.Context) *SkillTag {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SkillTagUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SkillTagUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *SkillTagUpdateOne) check() error {
	if v, ok := stuo.mutation.Name(); ok {
		if err := skilltag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillTag.name": %w`, err)}
		}
	}
	return nil
}

func (stuo *SkillTagUpdateOne) sqlSave(ctx context.Context) (_node *SkillTag, err error) {
	if err := stuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(skilltag.Table, skilltag.Columns, sqlgraph.NewFieldSpec(skilltag.FieldID, field.TypeInt))
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkillTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skilltag.FieldID)
		for _, f := range fields {
			if !skilltag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skilltag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.Name(); ok {
		_spec.SetField(skilltag.FieldName, field.TypeString, value)
	}
	if stuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !stuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltag.SkillsTable,
			Columns: []string{skilltag.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SkillTag{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	stuo.mutation.done = true
	return _node, nil
}
