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
	"github.com/sky0621/cv-admin/src/ent/usercareer"
)

// CareerTaskCreate is the builder for creating a CareerTask entity.
type CareerTaskCreate struct {
	config
	mutation *CareerTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ctc *CareerTaskCreate) SetCreateTime(t time.Time) *CareerTaskCreate {
	ctc.mutation.SetCreateTime(t)
	return ctc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ctc *CareerTaskCreate) SetNillableCreateTime(t *time.Time) *CareerTaskCreate {
	if t != nil {
		ctc.SetCreateTime(*t)
	}
	return ctc
}

// SetUpdateTime sets the "update_time" field.
func (ctc *CareerTaskCreate) SetUpdateTime(t time.Time) *CareerTaskCreate {
	ctc.mutation.SetUpdateTime(t)
	return ctc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ctc *CareerTaskCreate) SetNillableUpdateTime(t *time.Time) *CareerTaskCreate {
	if t != nil {
		ctc.SetUpdateTime(*t)
	}
	return ctc
}

// SetName sets the "name" field.
func (ctc *CareerTaskCreate) SetName(s string) *CareerTaskCreate {
	ctc.mutation.SetName(s)
	return ctc
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ctc *CareerTaskCreate) SetCareerID(id int) *CareerTaskCreate {
	ctc.mutation.SetCareerID(id)
	return ctc
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ctc *CareerTaskCreate) SetCareer(u *UserCareer) *CareerTaskCreate {
	return ctc.SetCareerID(u.ID)
}

// AddCareertaskdescriptionIDs adds the "careertaskdescriptions" edge to the CareerTaskDescription entity by IDs.
func (ctc *CareerTaskCreate) AddCareertaskdescriptionIDs(ids ...int) *CareerTaskCreate {
	ctc.mutation.AddCareertaskdescriptionIDs(ids...)
	return ctc
}

// AddCareertaskdescriptions adds the "careertaskdescriptions" edges to the CareerTaskDescription entity.
func (ctc *CareerTaskCreate) AddCareertaskdescriptions(c ...*CareerTaskDescription) *CareerTaskCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctc.AddCareertaskdescriptionIDs(ids...)
}

// Mutation returns the CareerTaskMutation object of the builder.
func (ctc *CareerTaskCreate) Mutation() *CareerTaskMutation {
	return ctc.mutation
}

// Save creates the CareerTask in the database.
func (ctc *CareerTaskCreate) Save(ctx context.Context) (*CareerTask, error) {
	var (
		err  error
		node *CareerTask
	)
	ctc.defaults()
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CareerTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			if node, err = ctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			if ctc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CareerTask)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CareerTaskMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CareerTaskCreate) SaveX(ctx context.Context) *CareerTask {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CareerTaskCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CareerTaskCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctc *CareerTaskCreate) defaults() {
	if _, ok := ctc.mutation.CreateTime(); !ok {
		v := careertask.DefaultCreateTime()
		ctc.mutation.SetCreateTime(v)
	}
	if _, ok := ctc.mutation.UpdateTime(); !ok {
		v := careertask.DefaultUpdateTime()
		ctc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CareerTaskCreate) check() error {
	if _, ok := ctc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "CareerTask.create_time"`)}
	}
	if _, ok := ctc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "CareerTask.update_time"`)}
	}
	if _, ok := ctc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CareerTask.name"`)}
	}
	if v, ok := ctc.mutation.Name(); ok {
		if err := careertask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CareerTask.name": %w`, err)}
		}
	}
	if _, ok := ctc.mutation.CareerID(); !ok {
		return &ValidationError{Name: "career", err: errors.New(`ent: missing required edge "CareerTask.career"`)}
	}
	return nil
}

func (ctc *CareerTaskCreate) sqlSave(ctx context.Context) (*CareerTask, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctc *CareerTaskCreate) createSpec() (*CareerTask, *sqlgraph.CreateSpec) {
	var (
		_node = &CareerTask{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: careertask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: careertask.FieldID,
			},
		}
	)
	_spec.OnConflict = ctc.conflict
	if value, ok := ctc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: careertask.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ctc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: careertask.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ctc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: careertask.FieldName,
		})
		_node.Name = value
	}
	if nodes := ctc.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   careertask.CareerTable,
			Columns: []string{careertask.CareerColumn},
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
		_node.career_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ctc.mutation.CareertaskdescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careertask.CareertaskdescriptionsTable,
			Columns: []string{careertask.CareertaskdescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careertaskdescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CareerTask.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CareerTaskUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ctc *CareerTaskCreate) OnConflict(opts ...sql.ConflictOption) *CareerTaskUpsertOne {
	ctc.conflict = opts
	return &CareerTaskUpsertOne{
		create: ctc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctc *CareerTaskCreate) OnConflictColumns(columns ...string) *CareerTaskUpsertOne {
	ctc.conflict = append(ctc.conflict, sql.ConflictColumns(columns...))
	return &CareerTaskUpsertOne{
		create: ctc,
	}
}

type (
	// CareerTaskUpsertOne is the builder for "upsert"-ing
	//  one CareerTask node.
	CareerTaskUpsertOne struct {
		create *CareerTaskCreate
	}

	// CareerTaskUpsert is the "OnConflict" setter.
	CareerTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *CareerTaskUpsert) SetCreateTime(v time.Time) *CareerTaskUpsert {
	u.Set(careertask.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CareerTaskUpsert) UpdateCreateTime() *CareerTaskUpsert {
	u.SetExcluded(careertask.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CareerTaskUpsert) SetUpdateTime(v time.Time) *CareerTaskUpsert {
	u.Set(careertask.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerTaskUpsert) UpdateUpdateTime() *CareerTaskUpsert {
	u.SetExcluded(careertask.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *CareerTaskUpsert) SetName(v string) *CareerTaskUpsert {
	u.Set(careertask.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerTaskUpsert) UpdateName() *CareerTaskUpsert {
	u.SetExcluded(careertask.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CareerTaskUpsertOne) UpdateNewValues() *CareerTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(careertask.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CareerTaskUpsertOne) Ignore() *CareerTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CareerTaskUpsertOne) DoNothing() *CareerTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CareerTaskCreate.OnConflict
// documentation for more info.
func (u *CareerTaskUpsertOne) Update(set func(*CareerTaskUpsert)) *CareerTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CareerTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CareerTaskUpsertOne) SetCreateTime(v time.Time) *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CareerTaskUpsertOne) UpdateCreateTime() *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CareerTaskUpsertOne) SetUpdateTime(v time.Time) *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerTaskUpsertOne) UpdateUpdateTime() *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CareerTaskUpsertOne) SetName(v string) *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerTaskUpsertOne) UpdateName() *CareerTaskUpsertOne {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *CareerTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CareerTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CareerTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CareerTaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CareerTaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CareerTaskCreateBulk is the builder for creating many CareerTask entities in bulk.
type CareerTaskCreateBulk struct {
	config
	builders []*CareerTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the CareerTask entities in the database.
func (ctcb *CareerTaskCreateBulk) Save(ctx context.Context) ([]*CareerTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CareerTask, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CareerTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CareerTaskCreateBulk) SaveX(ctx context.Context) []*CareerTask {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CareerTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CareerTaskCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CareerTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CareerTaskUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ctcb *CareerTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *CareerTaskUpsertBulk {
	ctcb.conflict = opts
	return &CareerTaskUpsertBulk{
		create: ctcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctcb *CareerTaskCreateBulk) OnConflictColumns(columns ...string) *CareerTaskUpsertBulk {
	ctcb.conflict = append(ctcb.conflict, sql.ConflictColumns(columns...))
	return &CareerTaskUpsertBulk{
		create: ctcb,
	}
}

// CareerTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of CareerTask nodes.
type CareerTaskUpsertBulk struct {
	create *CareerTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CareerTaskUpsertBulk) UpdateNewValues() *CareerTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(careertask.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CareerTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CareerTaskUpsertBulk) Ignore() *CareerTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CareerTaskUpsertBulk) DoNothing() *CareerTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CareerTaskCreateBulk.OnConflict
// documentation for more info.
func (u *CareerTaskUpsertBulk) Update(set func(*CareerTaskUpsert)) *CareerTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CareerTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CareerTaskUpsertBulk) SetCreateTime(v time.Time) *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CareerTaskUpsertBulk) UpdateCreateTime() *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CareerTaskUpsertBulk) SetUpdateTime(v time.Time) *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CareerTaskUpsertBulk) UpdateUpdateTime() *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CareerTaskUpsertBulk) SetName(v string) *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CareerTaskUpsertBulk) UpdateName() *CareerTaskUpsertBulk {
	return u.Update(func(s *CareerTaskUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *CareerTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CareerTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CareerTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CareerTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
