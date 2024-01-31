// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
)

// UserCareerDescriptionCreate is the builder for creating a UserCareerDescription entity.
type UserCareerDescriptionCreate struct {
	config
	mutation *UserCareerDescriptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (ucdc *UserCareerDescriptionCreate) SetDescription(s string) *UserCareerDescriptionCreate {
	ucdc.mutation.SetDescription(s)
	return ucdc
}

// SetCareerID sets the "career" edge to the UserCareer entity by ID.
func (ucdc *UserCareerDescriptionCreate) SetCareerID(id int) *UserCareerDescriptionCreate {
	ucdc.mutation.SetCareerID(id)
	return ucdc
}

// SetCareer sets the "career" edge to the UserCareer entity.
func (ucdc *UserCareerDescriptionCreate) SetCareer(u *UserCareer) *UserCareerDescriptionCreate {
	return ucdc.SetCareerID(u.ID)
}

// Mutation returns the UserCareerDescriptionMutation object of the builder.
func (ucdc *UserCareerDescriptionCreate) Mutation() *UserCareerDescriptionMutation {
	return ucdc.mutation
}

// Save creates the UserCareerDescription in the database.
func (ucdc *UserCareerDescriptionCreate) Save(ctx context.Context) (*UserCareerDescription, error) {
	return withHooks(ctx, ucdc.sqlSave, ucdc.mutation, ucdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ucdc *UserCareerDescriptionCreate) SaveX(ctx context.Context) *UserCareerDescription {
	v, err := ucdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucdc *UserCareerDescriptionCreate) Exec(ctx context.Context) error {
	_, err := ucdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdc *UserCareerDescriptionCreate) ExecX(ctx context.Context) {
	if err := ucdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucdc *UserCareerDescriptionCreate) check() error {
	if _, ok := ucdc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "UserCareerDescription.description"`)}
	}
	if v, ok := ucdc.mutation.Description(); ok {
		if err := usercareerdescription.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "UserCareerDescription.description": %w`, err)}
		}
	}
	if _, ok := ucdc.mutation.CareerID(); !ok {
		return &ValidationError{Name: "career", err: errors.New(`ent: missing required edge "UserCareerDescription.career"`)}
	}
	return nil
}

func (ucdc *UserCareerDescriptionCreate) sqlSave(ctx context.Context) (*UserCareerDescription, error) {
	if err := ucdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ucdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ucdc.mutation.id = &_node.ID
	ucdc.mutation.done = true
	return _node, nil
}

func (ucdc *UserCareerDescriptionCreate) createSpec() (*UserCareerDescription, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCareerDescription{config: ucdc.config}
		_spec = sqlgraph.NewCreateSpec(usercareerdescription.Table, sqlgraph.NewFieldSpec(usercareerdescription.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ucdc.conflict
	if value, ok := ucdc.mutation.Description(); ok {
		_spec.SetField(usercareerdescription.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := ucdc.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareerdescription.CareerTable,
			Columns: []string{usercareerdescription.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.career_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCareerDescription.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCareerDescriptionUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (ucdc *UserCareerDescriptionCreate) OnConflict(opts ...sql.ConflictOption) *UserCareerDescriptionUpsertOne {
	ucdc.conflict = opts
	return &UserCareerDescriptionUpsertOne{
		create: ucdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucdc *UserCareerDescriptionCreate) OnConflictColumns(columns ...string) *UserCareerDescriptionUpsertOne {
	ucdc.conflict = append(ucdc.conflict, sql.ConflictColumns(columns...))
	return &UserCareerDescriptionUpsertOne{
		create: ucdc,
	}
}

type (
	// UserCareerDescriptionUpsertOne is the builder for "upsert"-ing
	//  one UserCareerDescription node.
	UserCareerDescriptionUpsertOne struct {
		create *UserCareerDescriptionCreate
	}

	// UserCareerDescriptionUpsert is the "OnConflict" setter.
	UserCareerDescriptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *UserCareerDescriptionUpsert) SetDescription(v string) *UserCareerDescriptionUpsert {
	u.Set(usercareerdescription.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserCareerDescriptionUpsert) UpdateDescription() *UserCareerDescriptionUpsert {
	u.SetExcluded(usercareerdescription.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserCareerDescriptionUpsertOne) UpdateNewValues() *UserCareerDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserCareerDescriptionUpsertOne) Ignore() *UserCareerDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCareerDescriptionUpsertOne) DoNothing() *UserCareerDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCareerDescriptionCreate.OnConflict
// documentation for more info.
func (u *UserCareerDescriptionUpsertOne) Update(set func(*UserCareerDescriptionUpsert)) *UserCareerDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCareerDescriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *UserCareerDescriptionUpsertOne) SetDescription(v string) *UserCareerDescriptionUpsertOne {
	return u.Update(func(s *UserCareerDescriptionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserCareerDescriptionUpsertOne) UpdateDescription() *UserCareerDescriptionUpsertOne {
	return u.Update(func(s *UserCareerDescriptionUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *UserCareerDescriptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCareerDescriptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCareerDescriptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserCareerDescriptionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserCareerDescriptionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCareerDescriptionCreateBulk is the builder for creating many UserCareerDescription entities in bulk.
type UserCareerDescriptionCreateBulk struct {
	config
	builders []*UserCareerDescriptionCreate
	conflict []sql.ConflictOption
}

// Save creates the UserCareerDescription entities in the database.
func (ucdcb *UserCareerDescriptionCreateBulk) Save(ctx context.Context) ([]*UserCareerDescription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucdcb.builders))
	nodes := make([]*UserCareerDescription, len(ucdcb.builders))
	mutators := make([]Mutator, len(ucdcb.builders))
	for i := range ucdcb.builders {
		func(i int, root context.Context) {
			builder := ucdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCareerDescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, ucdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucdcb *UserCareerDescriptionCreateBulk) SaveX(ctx context.Context) []*UserCareerDescription {
	v, err := ucdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucdcb *UserCareerDescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := ucdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdcb *UserCareerDescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := ucdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCareerDescription.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCareerDescriptionUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (ucdcb *UserCareerDescriptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserCareerDescriptionUpsertBulk {
	ucdcb.conflict = opts
	return &UserCareerDescriptionUpsertBulk{
		create: ucdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucdcb *UserCareerDescriptionCreateBulk) OnConflictColumns(columns ...string) *UserCareerDescriptionUpsertBulk {
	ucdcb.conflict = append(ucdcb.conflict, sql.ConflictColumns(columns...))
	return &UserCareerDescriptionUpsertBulk{
		create: ucdcb,
	}
}

// UserCareerDescriptionUpsertBulk is the builder for "upsert"-ing
// a bulk of UserCareerDescription nodes.
type UserCareerDescriptionUpsertBulk struct {
	create *UserCareerDescriptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserCareerDescriptionUpsertBulk) UpdateNewValues() *UserCareerDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserCareerDescription.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserCareerDescriptionUpsertBulk) Ignore() *UserCareerDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCareerDescriptionUpsertBulk) DoNothing() *UserCareerDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCareerDescriptionCreateBulk.OnConflict
// documentation for more info.
func (u *UserCareerDescriptionUpsertBulk) Update(set func(*UserCareerDescriptionUpsert)) *UserCareerDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCareerDescriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *UserCareerDescriptionUpsertBulk) SetDescription(v string) *UserCareerDescriptionUpsertBulk {
	return u.Update(func(s *UserCareerDescriptionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserCareerDescriptionUpsertBulk) UpdateDescription() *UserCareerDescriptionUpsertBulk {
	return u.Update(func(s *UserCareerDescriptionUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *UserCareerDescriptionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCareerDescriptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCareerDescriptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCareerDescriptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
