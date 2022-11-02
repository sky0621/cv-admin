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
)

// CareerSkillCreate is the builder for creating a CareerSkill entity.
type CareerSkillCreate struct {
	config
	mutation *CareerSkillMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (csc *CareerSkillCreate) SetCreateTime(t time.Time) *CareerSkillCreate {
	csc.mutation.SetCreateTime(t)
	return csc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (csc *CareerSkillCreate) SetNillableCreateTime(t *time.Time) *CareerSkillCreate {
	if t != nil {
		csc.SetCreateTime(*t)
	}
	return csc
}

// SetUpdateTime sets the "update_time" field.
func (csc *CareerSkillCreate) SetUpdateTime(t time.Time) *CareerSkillCreate {
	csc.mutation.SetUpdateTime(t)
	return csc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (csc *CareerSkillCreate) SetNillableUpdateTime(t *time.Time) *CareerSkillCreate {
	if t != nil {
		csc.SetUpdateTime(*t)
	}
	return csc
}

// SetName sets the "name" field.
func (csc *CareerSkillCreate) SetName(s string) *CareerSkillCreate {
	csc.mutation.SetName(s)
	return csc
}

// SetURL sets the "url" field.
func (csc *CareerSkillCreate) SetURL(s string) *CareerSkillCreate {
	csc.mutation.SetURL(s)
	return csc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (csc *CareerSkillCreate) SetNillableURL(s *string) *CareerSkillCreate {
	if s != nil {
		csc.SetURL(*s)
	}
	return csc
}

// SetVersion sets the "version" field.
func (csc *CareerSkillCreate) SetVersion(s string) *CareerSkillCreate {
	csc.mutation.SetVersion(s)
	return csc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (csc *CareerSkillCreate) SetNillableVersion(s *string) *CareerSkillCreate {
	if s != nil {
		csc.SetVersion(*s)
	}
	return csc
}

// SetCareerSkillGroupID sets the "careerSkillGroup" edge to the CareerSkillGroup entity by ID.
func (csc *CareerSkillCreate) SetCareerSkillGroupID(id int) *CareerSkillCreate {
	csc.mutation.SetCareerSkillGroupID(id)
	return csc
}

// SetCareerSkillGroup sets the "careerSkillGroup" edge to the CareerSkillGroup entity.
func (csc *CareerSkillCreate) SetCareerSkillGroup(c *CareerSkillGroup) *CareerSkillCreate {
	return csc.SetCareerSkillGroupID(c.ID)
}

// Mutation returns the CareerSkillMutation object of the builder.
func (csc *CareerSkillCreate) Mutation() *CareerSkillMutation {
	return csc.mutation
}

// Save creates the CareerSkill in the database.
func (csc *CareerSkillCreate) Save(ctx context.Context) (*CareerSkill, error) {
	var (
		err  error
		node *CareerSkill
	)
	csc.defaults()
	if len(csc.hooks) == 0 {
		if err = csc.check(); err != nil {
			return nil, err
		}
		node, err = csc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerSkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csc.check(); err != nil {
				return nil, err
			}
			csc.mutation = mutation
			if node, err = csc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(csc.hooks) - 1; i >= 0; i-- {
			if csc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, csc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (csc *CareerSkillCreate) SaveX(ctx context.Context) *CareerSkill {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csc *CareerSkillCreate) Exec(ctx context.Context) error {
	_, err := csc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csc *CareerSkillCreate) ExecX(ctx context.Context) {
	if err := csc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csc *CareerSkillCreate) defaults() {
	if _, ok := csc.mutation.CreateTime(); !ok {
		v := careerskill.DefaultCreateTime()
		csc.mutation.SetCreateTime(v)
	}
	if _, ok := csc.mutation.UpdateTime(); !ok {
		v := careerskill.DefaultUpdateTime()
		csc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csc *CareerSkillCreate) check() error {
	if _, ok := csc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "CareerSkill.create_time"`)}
	}
	if _, ok := csc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "CareerSkill.update_time"`)}
	}
	if _, ok := csc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CareerSkill.name"`)}
	}
	if v, ok := csc.mutation.Name(); ok {
		if err := careerskill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.name": %w`, err)}
		}
	}
	if v, ok := csc.mutation.URL(); ok {
		if err := careerskill.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.url": %w`, err)}
		}
	}
	if v, ok := csc.mutation.Version(); ok {
		if err := careerskill.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "CareerSkill.version": %w`, err)}
		}
	}
	if _, ok := csc.mutation.CareerSkillGroupID(); !ok {
		return &ValidationError{Name: "careerSkillGroup", err: errors.New(`ent: missing required edge "CareerSkill.careerSkillGroup"`)}
	}
	return nil
}

func (csc *CareerSkillCreate) sqlSave(ctx context.Context) (*CareerSkill, error) {
	_node, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (csc *CareerSkillCreate) createSpec() (*CareerSkill, *sqlgraph.CreateSpec) {
	var (
		_node = &CareerSkill{config: csc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: careerskill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careerskill.FieldID,
			},
		}
	)
	_spec.OnConflict = csc.conflict
	if value, ok := csc.mutation.CreateTime(); ok {
		_spec.SetField(careerskill.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := csc.mutation.UpdateTime(); ok {
		_spec.SetField(careerskill.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := csc.mutation.Name(); ok {
		_spec.SetField(careerskill.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := csc.mutation.URL(); ok {
		_spec.SetField(careerskill.FieldURL, field.TypeString, value)
		_node.URL = &value
	}
	if value, ok := csc.mutation.Version(); ok {
		_spec.SetField(careerskill.FieldVersion, field.TypeString, value)
		_node.Version = &value
	}
	if nodes := csc.mutation.CareerSkillGroupIDs(); len(nodes) > 0 {
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
		_node.career_skill_group_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CareerSkill.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CareerSkillUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (csc *CareerSkillCreate) OnConflict(opts ...sql.ConflictOption) *CareerSkillUpsertOne {
	csc.conflict = opts
	return &CareerSkillUpsertOne{
		create: csc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (csc *CareerSkillCreate) OnConflictColumns(columns ...string) *CareerSkillUpsertOne {
	csc.conflict = append(csc.conflict, sql.ConflictColumns(columns...))
	return &CareerSkillUpsertOne{
		create: csc,
	}
}

type (
	// CareerSkillUpsertOne is the builder for "upsert"-ing
	//  one CareerSkill node.
	CareerSkillUpsertOne struct {
		create *CareerSkillCreate
	}

	// CareerSkillUpsert is the "OnConflict" setter.
	CareerSkillUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *CareerSkillUpsert) SetUpdateTime(v time.Time) *CareerSkillUpsert {
	u.Set(careerskill.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerSkillUpsert) UpdateUpdateTime() *CareerSkillUpsert {
	u.SetExcluded(careerskill.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *CareerSkillUpsert) SetName(v string) *CareerSkillUpsert {
	u.Set(careerskill.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerSkillUpsert) UpdateName() *CareerSkillUpsert {
	u.SetExcluded(careerskill.FieldName)
	return u
}

// SetURL sets the "url" field.
func (u *CareerSkillUpsert) SetURL(v string) *CareerSkillUpsert {
	u.Set(careerskill.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CareerSkillUpsert) UpdateURL() *CareerSkillUpsert {
	u.SetExcluded(careerskill.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *CareerSkillUpsert) ClearURL() *CareerSkillUpsert {
	u.SetNull(careerskill.FieldURL)
	return u
}

// SetVersion sets the "version" field.
func (u *CareerSkillUpsert) SetVersion(v string) *CareerSkillUpsert {
	u.Set(careerskill.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *CareerSkillUpsert) UpdateVersion() *CareerSkillUpsert {
	u.SetExcluded(careerskill.FieldVersion)
	return u
}

// ClearVersion clears the value of the "version" field.
func (u *CareerSkillUpsert) ClearVersion() *CareerSkillUpsert {
	u.SetNull(careerskill.FieldVersion)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CareerSkillUpsertOne) UpdateNewValues() *CareerSkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(careerskill.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CareerSkillUpsertOne) Ignore() *CareerSkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CareerSkillUpsertOne) DoNothing() *CareerSkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CareerSkillCreate.OnConflict
// documentation for more info.
func (u *CareerSkillUpsertOne) Update(set func(*CareerSkillUpsert)) *CareerSkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CareerSkillUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CareerSkillUpsertOne) SetUpdateTime(v time.Time) *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerSkillUpsertOne) UpdateUpdateTime() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CareerSkillUpsertOne) SetName(v string) *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerSkillUpsertOne) UpdateName() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *CareerSkillUpsertOne) SetURL(v string) *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CareerSkillUpsertOne) UpdateURL() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *CareerSkillUpsertOne) ClearURL() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.ClearURL()
	})
}

// SetVersion sets the "version" field.
func (u *CareerSkillUpsertOne) SetVersion(v string) *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *CareerSkillUpsertOne) UpdateVersion() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *CareerSkillUpsertOne) ClearVersion() *CareerSkillUpsertOne {
	return u.Update(func(s *CareerSkillUpsert) {
		s.ClearVersion()
	})
}

// Exec executes the query.
func (u *CareerSkillUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CareerSkillCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CareerSkillUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CareerSkillUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CareerSkillUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CareerSkillCreateBulk is the builder for creating many CareerSkill entities in bulk.
type CareerSkillCreateBulk struct {
	config
	builders []*CareerSkillCreate
	conflict []sql.ConflictOption
}

// Save creates the CareerSkill entities in the database.
func (cscb *CareerSkillCreateBulk) Save(ctx context.Context) ([]*CareerSkill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cscb.builders))
	nodes := make([]*CareerSkill, len(cscb.builders))
	mutators := make([]Mutator, len(cscb.builders))
	for i := range cscb.builders {
		func(i int, root context.Context) {
			builder := cscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CareerSkillMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cscb *CareerSkillCreateBulk) SaveX(ctx context.Context) []*CareerSkill {
	v, err := cscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cscb *CareerSkillCreateBulk) Exec(ctx context.Context) error {
	_, err := cscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscb *CareerSkillCreateBulk) ExecX(ctx context.Context) {
	if err := cscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CareerSkill.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CareerSkillUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (cscb *CareerSkillCreateBulk) OnConflict(opts ...sql.ConflictOption) *CareerSkillUpsertBulk {
	cscb.conflict = opts
	return &CareerSkillUpsertBulk{
		create: cscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cscb *CareerSkillCreateBulk) OnConflictColumns(columns ...string) *CareerSkillUpsertBulk {
	cscb.conflict = append(cscb.conflict, sql.ConflictColumns(columns...))
	return &CareerSkillUpsertBulk{
		create: cscb,
	}
}

// CareerSkillUpsertBulk is the builder for "upsert"-ing
// a bulk of CareerSkill nodes.
type CareerSkillUpsertBulk struct {
	create *CareerSkillCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CareerSkillUpsertBulk) UpdateNewValues() *CareerSkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(careerskill.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CareerSkill.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CareerSkillUpsertBulk) Ignore() *CareerSkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CareerSkillUpsertBulk) DoNothing() *CareerSkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CareerSkillCreateBulk.OnConflict
// documentation for more info.
func (u *CareerSkillUpsertBulk) Update(set func(*CareerSkillUpsert)) *CareerSkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CareerSkillUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CareerSkillUpsertBulk) SetUpdateTime(v time.Time) *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerSkillUpsertBulk) UpdateUpdateTime() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CareerSkillUpsertBulk) SetName(v string) *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerSkillUpsertBulk) UpdateName() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *CareerSkillUpsertBulk) SetURL(v string) *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CareerSkillUpsertBulk) UpdateURL() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *CareerSkillUpsertBulk) ClearURL() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.ClearURL()
	})
}

// SetVersion sets the "version" field.
func (u *CareerSkillUpsertBulk) SetVersion(v string) *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *CareerSkillUpsertBulk) UpdateVersion() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *CareerSkillUpsertBulk) ClearVersion() *CareerSkillUpsertBulk {
	return u.Update(func(s *CareerSkillUpsert) {
		s.ClearVersion()
	})
}

// Exec executes the query.
func (u *CareerSkillUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CareerSkillCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CareerSkillCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CareerSkillUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
