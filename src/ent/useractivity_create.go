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
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
)

// UserActivityCreate is the builder for creating a UserActivity entity.
type UserActivityCreate struct {
	config
	mutation *UserActivityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uac *UserActivityCreate) SetCreateTime(t time.Time) *UserActivityCreate {
	uac.mutation.SetCreateTime(t)
	return uac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uac *UserActivityCreate) SetNillableCreateTime(t *time.Time) *UserActivityCreate {
	if t != nil {
		uac.SetCreateTime(*t)
	}
	return uac
}

// SetUpdateTime sets the "update_time" field.
func (uac *UserActivityCreate) SetUpdateTime(t time.Time) *UserActivityCreate {
	uac.mutation.SetUpdateTime(t)
	return uac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uac *UserActivityCreate) SetNillableUpdateTime(t *time.Time) *UserActivityCreate {
	if t != nil {
		uac.SetUpdateTime(*t)
	}
	return uac
}

// SetName sets the "name" field.
func (uac *UserActivityCreate) SetName(s string) *UserActivityCreate {
	uac.mutation.SetName(s)
	return uac
}

// SetURL sets the "url" field.
func (uac *UserActivityCreate) SetURL(s string) *UserActivityCreate {
	uac.mutation.SetURL(s)
	return uac
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uac *UserActivityCreate) SetNillableURL(s *string) *UserActivityCreate {
	if s != nil {
		uac.SetURL(*s)
	}
	return uac
}

// SetIcon sets the "icon" field.
func (uac *UserActivityCreate) SetIcon(s string) *UserActivityCreate {
	uac.mutation.SetIcon(s)
	return uac
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (uac *UserActivityCreate) SetNillableIcon(s *string) *UserActivityCreate {
	if s != nil {
		uac.SetIcon(*s)
	}
	return uac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uac *UserActivityCreate) SetUserID(id int) *UserActivityCreate {
	uac.mutation.SetUserID(id)
	return uac
}

// SetUser sets the "user" edge to the User entity.
func (uac *UserActivityCreate) SetUser(u *User) *UserActivityCreate {
	return uac.SetUserID(u.ID)
}

// Mutation returns the UserActivityMutation object of the builder.
func (uac *UserActivityCreate) Mutation() *UserActivityMutation {
	return uac.mutation
}

// Save creates the UserActivity in the database.
func (uac *UserActivityCreate) Save(ctx context.Context) (*UserActivity, error) {
	uac.defaults()
	return withHooks(ctx, uac.sqlSave, uac.mutation, uac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uac *UserActivityCreate) SaveX(ctx context.Context) *UserActivity {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UserActivityCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UserActivityCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uac *UserActivityCreate) defaults() {
	if _, ok := uac.mutation.CreateTime(); !ok {
		v := useractivity.DefaultCreateTime()
		uac.mutation.SetCreateTime(v)
	}
	if _, ok := uac.mutation.UpdateTime(); !ok {
		v := useractivity.DefaultUpdateTime()
		uac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uac *UserActivityCreate) check() error {
	if _, ok := uac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserActivity.create_time"`)}
	}
	if _, ok := uac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserActivity.update_time"`)}
	}
	if _, ok := uac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserActivity.name"`)}
	}
	if v, ok := uac.mutation.Name(); ok {
		if err := useractivity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserActivity.name": %w`, err)}
		}
	}
	if v, ok := uac.mutation.URL(); ok {
		if err := useractivity.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "UserActivity.url": %w`, err)}
		}
	}
	if v, ok := uac.mutation.Icon(); ok {
		if err := useractivity.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "UserActivity.icon": %w`, err)}
		}
	}
	if _, ok := uac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserActivity.user"`)}
	}
	return nil
}

func (uac *UserActivityCreate) sqlSave(ctx context.Context) (*UserActivity, error) {
	if err := uac.check(); err != nil {
		return nil, err
	}
	_node, _spec := uac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uac.mutation.id = &_node.ID
	uac.mutation.done = true
	return _node, nil
}

func (uac *UserActivityCreate) createSpec() (*UserActivity, *sqlgraph.CreateSpec) {
	var (
		_node = &UserActivity{config: uac.config}
		_spec = sqlgraph.NewCreateSpec(useractivity.Table, sqlgraph.NewFieldSpec(useractivity.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uac.conflict
	if value, ok := uac.mutation.CreateTime(); ok {
		_spec.SetField(useractivity.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uac.mutation.UpdateTime(); ok {
		_spec.SetField(useractivity.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uac.mutation.Name(); ok {
		_spec.SetField(useractivity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uac.mutation.URL(); ok {
		_spec.SetField(useractivity.FieldURL, field.TypeString, value)
		_node.URL = &value
	}
	if value, ok := uac.mutation.Icon(); ok {
		_spec.SetField(useractivity.FieldIcon, field.TypeString, value)
		_node.Icon = &value
	}
	if nodes := uac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useractivity.UserTable,
			Columns: []string{useractivity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserActivity.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserActivityUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uac *UserActivityCreate) OnConflict(opts ...sql.ConflictOption) *UserActivityUpsertOne {
	uac.conflict = opts
	return &UserActivityUpsertOne{
		create: uac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uac *UserActivityCreate) OnConflictColumns(columns ...string) *UserActivityUpsertOne {
	uac.conflict = append(uac.conflict, sql.ConflictColumns(columns...))
	return &UserActivityUpsertOne{
		create: uac,
	}
}

type (
	// UserActivityUpsertOne is the builder for "upsert"-ing
	//  one UserActivity node.
	UserActivityUpsertOne struct {
		create *UserActivityCreate
	}

	// UserActivityUpsert is the "OnConflict" setter.
	UserActivityUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserActivityUpsert) SetUpdateTime(v time.Time) *UserActivityUpsert {
	u.Set(useractivity.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserActivityUpsert) UpdateUpdateTime() *UserActivityUpsert {
	u.SetExcluded(useractivity.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *UserActivityUpsert) SetName(v string) *UserActivityUpsert {
	u.Set(useractivity.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserActivityUpsert) UpdateName() *UserActivityUpsert {
	u.SetExcluded(useractivity.FieldName)
	return u
}

// SetURL sets the "url" field.
func (u *UserActivityUpsert) SetURL(v string) *UserActivityUpsert {
	u.Set(useractivity.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UserActivityUpsert) UpdateURL() *UserActivityUpsert {
	u.SetExcluded(useractivity.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *UserActivityUpsert) ClearURL() *UserActivityUpsert {
	u.SetNull(useractivity.FieldURL)
	return u
}

// SetIcon sets the "icon" field.
func (u *UserActivityUpsert) SetIcon(v string) *UserActivityUpsert {
	u.Set(useractivity.FieldIcon, v)
	return u
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *UserActivityUpsert) UpdateIcon() *UserActivityUpsert {
	u.SetExcluded(useractivity.FieldIcon)
	return u
}

// ClearIcon clears the value of the "icon" field.
func (u *UserActivityUpsert) ClearIcon() *UserActivityUpsert {
	u.SetNull(useractivity.FieldIcon)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserActivityUpsertOne) UpdateNewValues() *UserActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(useractivity.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserActivityUpsertOne) Ignore() *UserActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserActivityUpsertOne) DoNothing() *UserActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserActivityCreate.OnConflict
// documentation for more info.
func (u *UserActivityUpsertOne) Update(set func(*UserActivityUpsert)) *UserActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserActivityUpsertOne) SetUpdateTime(v time.Time) *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserActivityUpsertOne) UpdateUpdateTime() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserActivityUpsertOne) SetName(v string) *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserActivityUpsertOne) UpdateName() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *UserActivityUpsertOne) SetURL(v string) *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UserActivityUpsertOne) UpdateURL() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *UserActivityUpsertOne) ClearURL() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.ClearURL()
	})
}

// SetIcon sets the "icon" field.
func (u *UserActivityUpsertOne) SetIcon(v string) *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *UserActivityUpsertOne) UpdateIcon() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *UserActivityUpsertOne) ClearIcon() *UserActivityUpsertOne {
	return u.Update(func(s *UserActivityUpsert) {
		s.ClearIcon()
	})
}

// Exec executes the query.
func (u *UserActivityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserActivityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserActivityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserActivityUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserActivityUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserActivityCreateBulk is the builder for creating many UserActivity entities in bulk.
type UserActivityCreateBulk struct {
	config
	err      error
	builders []*UserActivityCreate
	conflict []sql.ConflictOption
}

// Save creates the UserActivity entities in the database.
func (uacb *UserActivityCreateBulk) Save(ctx context.Context) ([]*UserActivity, error) {
	if uacb.err != nil {
		return nil, uacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UserActivity, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserActivityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uacb *UserActivityCreateBulk) SaveX(ctx context.Context) []*UserActivity {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UserActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UserActivityCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserActivity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserActivityUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uacb *UserActivityCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserActivityUpsertBulk {
	uacb.conflict = opts
	return &UserActivityUpsertBulk{
		create: uacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uacb *UserActivityCreateBulk) OnConflictColumns(columns ...string) *UserActivityUpsertBulk {
	uacb.conflict = append(uacb.conflict, sql.ConflictColumns(columns...))
	return &UserActivityUpsertBulk{
		create: uacb,
	}
}

// UserActivityUpsertBulk is the builder for "upsert"-ing
// a bulk of UserActivity nodes.
type UserActivityUpsertBulk struct {
	create *UserActivityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserActivityUpsertBulk) UpdateNewValues() *UserActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(useractivity.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserActivity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserActivityUpsertBulk) Ignore() *UserActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserActivityUpsertBulk) DoNothing() *UserActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserActivityCreateBulk.OnConflict
// documentation for more info.
func (u *UserActivityUpsertBulk) Update(set func(*UserActivityUpsert)) *UserActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserActivityUpsertBulk) SetUpdateTime(v time.Time) *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserActivityUpsertBulk) UpdateUpdateTime() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserActivityUpsertBulk) SetName(v string) *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserActivityUpsertBulk) UpdateName() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *UserActivityUpsertBulk) SetURL(v string) *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UserActivityUpsertBulk) UpdateURL() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *UserActivityUpsertBulk) ClearURL() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.ClearURL()
	})
}

// SetIcon sets the "icon" field.
func (u *UserActivityUpsertBulk) SetIcon(v string) *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *UserActivityUpsertBulk) UpdateIcon() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *UserActivityUpsertBulk) ClearIcon() *UserActivityUpsertBulk {
	return u.Update(func(s *UserActivityUpsert) {
		s.ClearIcon()
	})
}

// Exec executes the query.
func (u *UserActivityUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserActivityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserActivityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserActivityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
