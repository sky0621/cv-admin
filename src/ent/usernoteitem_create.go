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
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
)

// UserNoteItemCreate is the builder for creating a UserNoteItem entity.
type UserNoteItemCreate struct {
	config
	mutation *UserNoteItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (unic *UserNoteItemCreate) SetCreateTime(t time.Time) *UserNoteItemCreate {
	unic.mutation.SetCreateTime(t)
	return unic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (unic *UserNoteItemCreate) SetNillableCreateTime(t *time.Time) *UserNoteItemCreate {
	if t != nil {
		unic.SetCreateTime(*t)
	}
	return unic
}

// SetUpdateTime sets the "update_time" field.
func (unic *UserNoteItemCreate) SetUpdateTime(t time.Time) *UserNoteItemCreate {
	unic.mutation.SetUpdateTime(t)
	return unic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (unic *UserNoteItemCreate) SetNillableUpdateTime(t *time.Time) *UserNoteItemCreate {
	if t != nil {
		unic.SetUpdateTime(*t)
	}
	return unic
}

// Mutation returns the UserNoteItemMutation object of the builder.
func (unic *UserNoteItemCreate) Mutation() *UserNoteItemMutation {
	return unic.mutation
}

// Save creates the UserNoteItem in the database.
func (unic *UserNoteItemCreate) Save(ctx context.Context) (*UserNoteItem, error) {
	var (
		err  error
		node *UserNoteItem
	)
	unic.defaults()
	if len(unic.hooks) == 0 {
		if err = unic.check(); err != nil {
			return nil, err
		}
		node, err = unic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNoteItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = unic.check(); err != nil {
				return nil, err
			}
			unic.mutation = mutation
			if node, err = unic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(unic.hooks) - 1; i >= 0; i-- {
			if unic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = unic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, unic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserNoteItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserNoteItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (unic *UserNoteItemCreate) SaveX(ctx context.Context) *UserNoteItem {
	v, err := unic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (unic *UserNoteItemCreate) Exec(ctx context.Context) error {
	_, err := unic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (unic *UserNoteItemCreate) ExecX(ctx context.Context) {
	if err := unic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (unic *UserNoteItemCreate) defaults() {
	if _, ok := unic.mutation.CreateTime(); !ok {
		v := usernoteitem.DefaultCreateTime()
		unic.mutation.SetCreateTime(v)
	}
	if _, ok := unic.mutation.UpdateTime(); !ok {
		v := usernoteitem.DefaultUpdateTime()
		unic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (unic *UserNoteItemCreate) check() error {
	if _, ok := unic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserNoteItem.create_time"`)}
	}
	if _, ok := unic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserNoteItem.update_time"`)}
	}
	return nil
}

func (unic *UserNoteItemCreate) sqlSave(ctx context.Context) (*UserNoteItem, error) {
	_node, _spec := unic.createSpec()
	if err := sqlgraph.CreateNode(ctx, unic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (unic *UserNoteItemCreate) createSpec() (*UserNoteItem, *sqlgraph.CreateSpec) {
	var (
		_node = &UserNoteItem{config: unic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usernoteitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernoteitem.FieldID,
			},
		}
	)
	_spec.OnConflict = unic.conflict
	if value, ok := unic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usernoteitem.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := unic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usernoteitem.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserNoteItem.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserNoteItemUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (unic *UserNoteItemCreate) OnConflict(opts ...sql.ConflictOption) *UserNoteItemUpsertOne {
	unic.conflict = opts
	return &UserNoteItemUpsertOne{
		create: unic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (unic *UserNoteItemCreate) OnConflictColumns(columns ...string) *UserNoteItemUpsertOne {
	unic.conflict = append(unic.conflict, sql.ConflictColumns(columns...))
	return &UserNoteItemUpsertOne{
		create: unic,
	}
}

type (
	// UserNoteItemUpsertOne is the builder for "upsert"-ing
	//  one UserNoteItem node.
	UserNoteItemUpsertOne struct {
		create *UserNoteItemCreate
	}

	// UserNoteItemUpsert is the "OnConflict" setter.
	UserNoteItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *UserNoteItemUpsert) SetCreateTime(v time.Time) *UserNoteItemUpsert {
	u.Set(usernoteitem.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserNoteItemUpsert) UpdateCreateTime() *UserNoteItemUpsert {
	u.SetExcluded(usernoteitem.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserNoteItemUpsert) SetUpdateTime(v time.Time) *UserNoteItemUpsert {
	u.Set(usernoteitem.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserNoteItemUpsert) UpdateUpdateTime() *UserNoteItemUpsert {
	u.SetExcluded(usernoteitem.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserNoteItemUpsertOne) UpdateNewValues() *UserNoteItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(usernoteitem.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserNoteItemUpsertOne) Ignore() *UserNoteItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserNoteItemUpsertOne) DoNothing() *UserNoteItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserNoteItemCreate.OnConflict
// documentation for more info.
func (u *UserNoteItemUpsertOne) Update(set func(*UserNoteItemUpsert)) *UserNoteItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserNoteItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *UserNoteItemUpsertOne) SetCreateTime(v time.Time) *UserNoteItemUpsertOne {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserNoteItemUpsertOne) UpdateCreateTime() *UserNoteItemUpsertOne {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *UserNoteItemUpsertOne) SetUpdateTime(v time.Time) *UserNoteItemUpsertOne {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserNoteItemUpsertOne) UpdateUpdateTime() *UserNoteItemUpsertOne {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *UserNoteItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserNoteItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserNoteItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserNoteItemUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserNoteItemUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserNoteItemCreateBulk is the builder for creating many UserNoteItem entities in bulk.
type UserNoteItemCreateBulk struct {
	config
	builders []*UserNoteItemCreate
	conflict []sql.ConflictOption
}

// Save creates the UserNoteItem entities in the database.
func (unicb *UserNoteItemCreateBulk) Save(ctx context.Context) ([]*UserNoteItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(unicb.builders))
	nodes := make([]*UserNoteItem, len(unicb.builders))
	mutators := make([]Mutator, len(unicb.builders))
	for i := range unicb.builders {
		func(i int, root context.Context) {
			builder := unicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserNoteItemMutation)
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
					_, err = mutators[i+1].Mutate(root, unicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = unicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, unicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, unicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (unicb *UserNoteItemCreateBulk) SaveX(ctx context.Context) []*UserNoteItem {
	v, err := unicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (unicb *UserNoteItemCreateBulk) Exec(ctx context.Context) error {
	_, err := unicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (unicb *UserNoteItemCreateBulk) ExecX(ctx context.Context) {
	if err := unicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserNoteItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserNoteItemUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (unicb *UserNoteItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserNoteItemUpsertBulk {
	unicb.conflict = opts
	return &UserNoteItemUpsertBulk{
		create: unicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (unicb *UserNoteItemCreateBulk) OnConflictColumns(columns ...string) *UserNoteItemUpsertBulk {
	unicb.conflict = append(unicb.conflict, sql.ConflictColumns(columns...))
	return &UserNoteItemUpsertBulk{
		create: unicb,
	}
}

// UserNoteItemUpsertBulk is the builder for "upsert"-ing
// a bulk of UserNoteItem nodes.
type UserNoteItemUpsertBulk struct {
	create *UserNoteItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserNoteItemUpsertBulk) UpdateNewValues() *UserNoteItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(usernoteitem.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserNoteItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserNoteItemUpsertBulk) Ignore() *UserNoteItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserNoteItemUpsertBulk) DoNothing() *UserNoteItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserNoteItemCreateBulk.OnConflict
// documentation for more info.
func (u *UserNoteItemUpsertBulk) Update(set func(*UserNoteItemUpsert)) *UserNoteItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserNoteItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *UserNoteItemUpsertBulk) SetCreateTime(v time.Time) *UserNoteItemUpsertBulk {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserNoteItemUpsertBulk) UpdateCreateTime() *UserNoteItemUpsertBulk {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *UserNoteItemUpsertBulk) SetUpdateTime(v time.Time) *UserNoteItemUpsertBulk {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserNoteItemUpsertBulk) UpdateUpdateTime() *UserNoteItemUpsertBulk {
	return u.Update(func(s *UserNoteItemUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *UserNoteItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserNoteItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserNoteItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserNoteItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
