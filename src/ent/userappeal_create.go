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
	"github.com/sky0621/cv-admin/src/ent/userappeal"
)

// UserAppealCreate is the builder for creating a UserAppeal entity.
type UserAppealCreate struct {
	config
	mutation *UserAppealMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uac *UserAppealCreate) SetCreateTime(t time.Time) *UserAppealCreate {
	uac.mutation.SetCreateTime(t)
	return uac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uac *UserAppealCreate) SetNillableCreateTime(t *time.Time) *UserAppealCreate {
	if t != nil {
		uac.SetCreateTime(*t)
	}
	return uac
}

// SetUpdateTime sets the "update_time" field.
func (uac *UserAppealCreate) SetUpdateTime(t time.Time) *UserAppealCreate {
	uac.mutation.SetUpdateTime(t)
	return uac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uac *UserAppealCreate) SetNillableUpdateTime(t *time.Time) *UserAppealCreate {
	if t != nil {
		uac.SetUpdateTime(*t)
	}
	return uac
}

// SetContent sets the "content" field.
func (uac *UserAppealCreate) SetContent(s string) *UserAppealCreate {
	uac.mutation.SetContent(s)
	return uac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uac *UserAppealCreate) SetUserID(id int) *UserAppealCreate {
	uac.mutation.SetUserID(id)
	return uac
}

// SetUser sets the "user" edge to the User entity.
func (uac *UserAppealCreate) SetUser(u *User) *UserAppealCreate {
	return uac.SetUserID(u.ID)
}

// Mutation returns the UserAppealMutation object of the builder.
func (uac *UserAppealCreate) Mutation() *UserAppealMutation {
	return uac.mutation
}

// Save creates the UserAppeal in the database.
func (uac *UserAppealCreate) Save(ctx context.Context) (*UserAppeal, error) {
	uac.defaults()
	return withHooks(ctx, uac.sqlSave, uac.mutation, uac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uac *UserAppealCreate) SaveX(ctx context.Context) *UserAppeal {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UserAppealCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UserAppealCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uac *UserAppealCreate) defaults() {
	if _, ok := uac.mutation.CreateTime(); !ok {
		v := userappeal.DefaultCreateTime()
		uac.mutation.SetCreateTime(v)
	}
	if _, ok := uac.mutation.UpdateTime(); !ok {
		v := userappeal.DefaultUpdateTime()
		uac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uac *UserAppealCreate) check() error {
	if _, ok := uac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserAppeal.create_time"`)}
	}
	if _, ok := uac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserAppeal.update_time"`)}
	}
	if _, ok := uac.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "UserAppeal.content"`)}
	}
	if v, ok := uac.mutation.Content(); ok {
		if err := userappeal.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "UserAppeal.content": %w`, err)}
		}
	}
	if _, ok := uac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserAppeal.user"`)}
	}
	return nil
}

func (uac *UserAppealCreate) sqlSave(ctx context.Context) (*UserAppeal, error) {
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

func (uac *UserAppealCreate) createSpec() (*UserAppeal, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAppeal{config: uac.config}
		_spec = sqlgraph.NewCreateSpec(userappeal.Table, sqlgraph.NewFieldSpec(userappeal.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uac.conflict
	if value, ok := uac.mutation.CreateTime(); ok {
		_spec.SetField(userappeal.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uac.mutation.UpdateTime(); ok {
		_spec.SetField(userappeal.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uac.mutation.Content(); ok {
		_spec.SetField(userappeal.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := uac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userappeal.UserTable,
			Columns: []string{userappeal.UserColumn},
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
//	client.UserAppeal.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserAppealUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uac *UserAppealCreate) OnConflict(opts ...sql.ConflictOption) *UserAppealUpsertOne {
	uac.conflict = opts
	return &UserAppealUpsertOne{
		create: uac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uac *UserAppealCreate) OnConflictColumns(columns ...string) *UserAppealUpsertOne {
	uac.conflict = append(uac.conflict, sql.ConflictColumns(columns...))
	return &UserAppealUpsertOne{
		create: uac,
	}
}

type (
	// UserAppealUpsertOne is the builder for "upsert"-ing
	//  one UserAppeal node.
	UserAppealUpsertOne struct {
		create *UserAppealCreate
	}

	// UserAppealUpsert is the "OnConflict" setter.
	UserAppealUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserAppealUpsert) SetUpdateTime(v time.Time) *UserAppealUpsert {
	u.Set(userappeal.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserAppealUpsert) UpdateUpdateTime() *UserAppealUpsert {
	u.SetExcluded(userappeal.FieldUpdateTime)
	return u
}

// SetContent sets the "content" field.
func (u *UserAppealUpsert) SetContent(v string) *UserAppealUpsert {
	u.Set(userappeal.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserAppealUpsert) UpdateContent() *UserAppealUpsert {
	u.SetExcluded(userappeal.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserAppealUpsertOne) UpdateNewValues() *UserAppealUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(userappeal.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserAppealUpsertOne) Ignore() *UserAppealUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserAppealUpsertOne) DoNothing() *UserAppealUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserAppealCreate.OnConflict
// documentation for more info.
func (u *UserAppealUpsertOne) Update(set func(*UserAppealUpsert)) *UserAppealUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserAppealUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserAppealUpsertOne) SetUpdateTime(v time.Time) *UserAppealUpsertOne {
	return u.Update(func(s *UserAppealUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserAppealUpsertOne) UpdateUpdateTime() *UserAppealUpsertOne {
	return u.Update(func(s *UserAppealUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetContent sets the "content" field.
func (u *UserAppealUpsertOne) SetContent(v string) *UserAppealUpsertOne {
	return u.Update(func(s *UserAppealUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserAppealUpsertOne) UpdateContent() *UserAppealUpsertOne {
	return u.Update(func(s *UserAppealUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *UserAppealUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserAppealCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserAppealUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserAppealUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserAppealUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserAppealCreateBulk is the builder for creating many UserAppeal entities in bulk.
type UserAppealCreateBulk struct {
	config
	err      error
	builders []*UserAppealCreate
	conflict []sql.ConflictOption
}

// Save creates the UserAppeal entities in the database.
func (uacb *UserAppealCreateBulk) Save(ctx context.Context) ([]*UserAppeal, error) {
	if uacb.err != nil {
		return nil, uacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UserAppeal, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAppealMutation)
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
func (uacb *UserAppealCreateBulk) SaveX(ctx context.Context) []*UserAppeal {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UserAppealCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UserAppealCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserAppeal.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserAppealUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uacb *UserAppealCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserAppealUpsertBulk {
	uacb.conflict = opts
	return &UserAppealUpsertBulk{
		create: uacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uacb *UserAppealCreateBulk) OnConflictColumns(columns ...string) *UserAppealUpsertBulk {
	uacb.conflict = append(uacb.conflict, sql.ConflictColumns(columns...))
	return &UserAppealUpsertBulk{
		create: uacb,
	}
}

// UserAppealUpsertBulk is the builder for "upsert"-ing
// a bulk of UserAppeal nodes.
type UserAppealUpsertBulk struct {
	create *UserAppealCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserAppealUpsertBulk) UpdateNewValues() *UserAppealUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(userappeal.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserAppeal.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserAppealUpsertBulk) Ignore() *UserAppealUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserAppealUpsertBulk) DoNothing() *UserAppealUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserAppealCreateBulk.OnConflict
// documentation for more info.
func (u *UserAppealUpsertBulk) Update(set func(*UserAppealUpsert)) *UserAppealUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserAppealUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserAppealUpsertBulk) SetUpdateTime(v time.Time) *UserAppealUpsertBulk {
	return u.Update(func(s *UserAppealUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserAppealUpsertBulk) UpdateUpdateTime() *UserAppealUpsertBulk {
	return u.Update(func(s *UserAppealUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetContent sets the "content" field.
func (u *UserAppealUpsertBulk) SetContent(v string) *UserAppealUpsertBulk {
	return u.Update(func(s *UserAppealUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserAppealUpsertBulk) UpdateContent() *UserAppealUpsertBulk {
	return u.Update(func(s *UserAppealUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *UserAppealUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserAppealCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserAppealCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserAppealUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
