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
	"github.com/sky0621/cv-admin/src/ent/usersolution"
)

// UserSolutionCreate is the builder for creating a UserSolution entity.
type UserSolutionCreate struct {
	config
	mutation *UserSolutionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (usc *UserSolutionCreate) SetCreateTime(t time.Time) *UserSolutionCreate {
	usc.mutation.SetCreateTime(t)
	return usc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (usc *UserSolutionCreate) SetNillableCreateTime(t *time.Time) *UserSolutionCreate {
	if t != nil {
		usc.SetCreateTime(*t)
	}
	return usc
}

// SetUpdateTime sets the "update_time" field.
func (usc *UserSolutionCreate) SetUpdateTime(t time.Time) *UserSolutionCreate {
	usc.mutation.SetUpdateTime(t)
	return usc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (usc *UserSolutionCreate) SetNillableUpdateTime(t *time.Time) *UserSolutionCreate {
	if t != nil {
		usc.SetUpdateTime(*t)
	}
	return usc
}

// SetContent sets the "content" field.
func (usc *UserSolutionCreate) SetContent(s string) *UserSolutionCreate {
	usc.mutation.SetContent(s)
	return usc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usc *UserSolutionCreate) SetUserID(id int) *UserSolutionCreate {
	usc.mutation.SetUserID(id)
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSolutionCreate) SetUser(u *User) *UserSolutionCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSolutionMutation object of the builder.
func (usc *UserSolutionCreate) Mutation() *UserSolutionMutation {
	return usc.mutation
}

// Save creates the UserSolution in the database.
func (usc *UserSolutionCreate) Save(ctx context.Context) (*UserSolution, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSolutionCreate) SaveX(ctx context.Context) *UserSolution {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSolutionCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSolutionCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSolutionCreate) defaults() {
	if _, ok := usc.mutation.CreateTime(); !ok {
		v := usersolution.DefaultCreateTime()
		usc.mutation.SetCreateTime(v)
	}
	if _, ok := usc.mutation.UpdateTime(); !ok {
		v := usersolution.DefaultUpdateTime()
		usc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSolutionCreate) check() error {
	if _, ok := usc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserSolution.create_time"`)}
	}
	if _, ok := usc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserSolution.update_time"`)}
	}
	if _, ok := usc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "UserSolution.content"`)}
	}
	if v, ok := usc.mutation.Content(); ok {
		if err := usersolution.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "UserSolution.content": %w`, err)}
		}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserSolution.user"`)}
	}
	return nil
}

func (usc *UserSolutionCreate) sqlSave(ctx context.Context) (*UserSolution, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSolutionCreate) createSpec() (*UserSolution, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSolution{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersolution.Table, sqlgraph.NewFieldSpec(usersolution.FieldID, field.TypeInt))
	)
	_spec.OnConflict = usc.conflict
	if value, ok := usc.mutation.CreateTime(); ok {
		_spec.SetField(usersolution.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := usc.mutation.UpdateTime(); ok {
		_spec.SetField(usersolution.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := usc.mutation.Content(); ok {
		_spec.SetField(usersolution.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersolution.UserTable,
			Columns: []string{usersolution.UserColumn},
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
//	client.UserSolution.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserSolutionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (usc *UserSolutionCreate) OnConflict(opts ...sql.ConflictOption) *UserSolutionUpsertOne {
	usc.conflict = opts
	return &UserSolutionUpsertOne{
		create: usc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (usc *UserSolutionCreate) OnConflictColumns(columns ...string) *UserSolutionUpsertOne {
	usc.conflict = append(usc.conflict, sql.ConflictColumns(columns...))
	return &UserSolutionUpsertOne{
		create: usc,
	}
}

type (
	// UserSolutionUpsertOne is the builder for "upsert"-ing
	//  one UserSolution node.
	UserSolutionUpsertOne struct {
		create *UserSolutionCreate
	}

	// UserSolutionUpsert is the "OnConflict" setter.
	UserSolutionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserSolutionUpsert) SetUpdateTime(v time.Time) *UserSolutionUpsert {
	u.Set(usersolution.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserSolutionUpsert) UpdateUpdateTime() *UserSolutionUpsert {
	u.SetExcluded(usersolution.FieldUpdateTime)
	return u
}

// SetContent sets the "content" field.
func (u *UserSolutionUpsert) SetContent(v string) *UserSolutionUpsert {
	u.Set(usersolution.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserSolutionUpsert) UpdateContent() *UserSolutionUpsert {
	u.SetExcluded(usersolution.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserSolutionUpsertOne) UpdateNewValues() *UserSolutionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(usersolution.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserSolutionUpsertOne) Ignore() *UserSolutionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserSolutionUpsertOne) DoNothing() *UserSolutionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserSolutionCreate.OnConflict
// documentation for more info.
func (u *UserSolutionUpsertOne) Update(set func(*UserSolutionUpsert)) *UserSolutionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserSolutionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserSolutionUpsertOne) SetUpdateTime(v time.Time) *UserSolutionUpsertOne {
	return u.Update(func(s *UserSolutionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserSolutionUpsertOne) UpdateUpdateTime() *UserSolutionUpsertOne {
	return u.Update(func(s *UserSolutionUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetContent sets the "content" field.
func (u *UserSolutionUpsertOne) SetContent(v string) *UserSolutionUpsertOne {
	return u.Update(func(s *UserSolutionUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserSolutionUpsertOne) UpdateContent() *UserSolutionUpsertOne {
	return u.Update(func(s *UserSolutionUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *UserSolutionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserSolutionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserSolutionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserSolutionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserSolutionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserSolutionCreateBulk is the builder for creating many UserSolution entities in bulk.
type UserSolutionCreateBulk struct {
	config
	err      error
	builders []*UserSolutionCreate
	conflict []sql.ConflictOption
}

// Save creates the UserSolution entities in the database.
func (uscb *UserSolutionCreateBulk) Save(ctx context.Context) ([]*UserSolution, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSolution, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSolutionMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSolutionCreateBulk) SaveX(ctx context.Context) []*UserSolution {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSolutionCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSolutionCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserSolution.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserSolutionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uscb *UserSolutionCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserSolutionUpsertBulk {
	uscb.conflict = opts
	return &UserSolutionUpsertBulk{
		create: uscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uscb *UserSolutionCreateBulk) OnConflictColumns(columns ...string) *UserSolutionUpsertBulk {
	uscb.conflict = append(uscb.conflict, sql.ConflictColumns(columns...))
	return &UserSolutionUpsertBulk{
		create: uscb,
	}
}

// UserSolutionUpsertBulk is the builder for "upsert"-ing
// a bulk of UserSolution nodes.
type UserSolutionUpsertBulk struct {
	create *UserSolutionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserSolutionUpsertBulk) UpdateNewValues() *UserSolutionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(usersolution.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserSolution.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserSolutionUpsertBulk) Ignore() *UserSolutionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserSolutionUpsertBulk) DoNothing() *UserSolutionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserSolutionCreateBulk.OnConflict
// documentation for more info.
func (u *UserSolutionUpsertBulk) Update(set func(*UserSolutionUpsert)) *UserSolutionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserSolutionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserSolutionUpsertBulk) SetUpdateTime(v time.Time) *UserSolutionUpsertBulk {
	return u.Update(func(s *UserSolutionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserSolutionUpsertBulk) UpdateUpdateTime() *UserSolutionUpsertBulk {
	return u.Update(func(s *UserSolutionUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetContent sets the "content" field.
func (u *UserSolutionUpsertBulk) SetContent(v string) *UserSolutionUpsertBulk {
	return u.Update(func(s *UserSolutionUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *UserSolutionUpsertBulk) UpdateContent() *UserSolutionUpsertBulk {
	return u.Update(func(s *UserSolutionUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *UserSolutionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserSolutionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserSolutionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserSolutionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}