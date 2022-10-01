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
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// UserCareerCreate is the builder for creating a UserCareer entity.
type UserCareerCreate struct {
	config
	mutation *UserCareerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ucc *UserCareerCreate) SetCreateTime(t time.Time) *UserCareerCreate {
	ucc.mutation.SetCreateTime(t)
	return ucc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ucc *UserCareerCreate) SetNillableCreateTime(t *time.Time) *UserCareerCreate {
	if t != nil {
		ucc.SetCreateTime(*t)
	}
	return ucc
}

// SetUpdateTime sets the "update_time" field.
func (ucc *UserCareerCreate) SetUpdateTime(t time.Time) *UserCareerCreate {
	ucc.mutation.SetUpdateTime(t)
	return ucc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ucc *UserCareerCreate) SetNillableUpdateTime(t *time.Time) *UserCareerCreate {
	if t != nil {
		ucc.SetUpdateTime(*t)
	}
	return ucc
}

// SetName sets the "name" field.
func (ucc *UserCareerCreate) SetName(s string) *UserCareerCreate {
	ucc.mutation.SetName(s)
	return ucc
}

// SetFrom sets the "from" field.
func (ucc *UserCareerCreate) SetFrom(s string) *UserCareerCreate {
	ucc.mutation.SetFrom(s)
	return ucc
}

// SetTo sets the "to" field.
func (ucc *UserCareerCreate) SetTo(s string) *UserCareerCreate {
	ucc.mutation.SetTo(s)
	return ucc
}

// SetCareergroupID sets the "careergroup" edge to the UserCareerGroup entity by ID.
func (ucc *UserCareerCreate) SetCareergroupID(id int) *UserCareerCreate {
	ucc.mutation.SetCareergroupID(id)
	return ucc
}

// SetCareergroup sets the "careergroup" edge to the UserCareerGroup entity.
func (ucc *UserCareerCreate) SetCareergroup(u *UserCareerGroup) *UserCareerCreate {
	return ucc.SetCareergroupID(u.ID)
}

// AddCareerdescriptionIDs adds the "careerdescriptions" edge to the UserCareerDescription entity by IDs.
func (ucc *UserCareerCreate) AddCareerdescriptionIDs(ids ...int) *UserCareerCreate {
	ucc.mutation.AddCareerdescriptionIDs(ids...)
	return ucc
}

// AddCareerdescriptions adds the "careerdescriptions" edges to the UserCareerDescription entity.
func (ucc *UserCareerCreate) AddCareerdescriptions(u ...*UserCareerDescription) *UserCareerCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucc.AddCareerdescriptionIDs(ids...)
}

// AddCareertaskIDs adds the "careertasks" edge to the CareerTask entity by IDs.
func (ucc *UserCareerCreate) AddCareertaskIDs(ids ...int) *UserCareerCreate {
	ucc.mutation.AddCareertaskIDs(ids...)
	return ucc
}

// AddCareertasks adds the "careertasks" edges to the CareerTask entity.
func (ucc *UserCareerCreate) AddCareertasks(c ...*CareerTask) *UserCareerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucc.AddCareertaskIDs(ids...)
}

// AddCareerskillgroupIDs adds the "careerskillgroups" edge to the CareerSkillGroup entity by IDs.
func (ucc *UserCareerCreate) AddCareerskillgroupIDs(ids ...int) *UserCareerCreate {
	ucc.mutation.AddCareerskillgroupIDs(ids...)
	return ucc
}

// AddCareerskillgroups adds the "careerskillgroups" edges to the CareerSkillGroup entity.
func (ucc *UserCareerCreate) AddCareerskillgroups(c ...*CareerSkillGroup) *UserCareerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucc.AddCareerskillgroupIDs(ids...)
}

// Mutation returns the UserCareerMutation object of the builder.
func (ucc *UserCareerCreate) Mutation() *UserCareerMutation {
	return ucc.mutation
}

// Save creates the UserCareer in the database.
func (ucc *UserCareerCreate) Save(ctx context.Context) (*UserCareer, error) {
	var (
		err  error
		node *UserCareer
	)
	ucc.defaults()
	if len(ucc.hooks) == 0 {
		if err = ucc.check(); err != nil {
			return nil, err
		}
		node, err = ucc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCareerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucc.check(); err != nil {
				return nil, err
			}
			ucc.mutation = mutation
			if node, err = ucc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucc.hooks) - 1; i >= 0; i-- {
			if ucc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ucc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserCareer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserCareerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UserCareerCreate) SaveX(ctx context.Context) *UserCareer {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UserCareerCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UserCareerCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucc *UserCareerCreate) defaults() {
	if _, ok := ucc.mutation.CreateTime(); !ok {
		v := usercareer.DefaultCreateTime()
		ucc.mutation.SetCreateTime(v)
	}
	if _, ok := ucc.mutation.UpdateTime(); !ok {
		v := usercareer.DefaultUpdateTime()
		ucc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UserCareerCreate) check() error {
	if _, ok := ucc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserCareer.create_time"`)}
	}
	if _, ok := ucc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserCareer.update_time"`)}
	}
	if _, ok := ucc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserCareer.name"`)}
	}
	if v, ok := ucc.mutation.Name(); ok {
		if err := usercareer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserCareer.name": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "UserCareer.from"`)}
	}
	if v, ok := ucc.mutation.From(); ok {
		if err := usercareer.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "UserCareer.from": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "UserCareer.to"`)}
	}
	if v, ok := ucc.mutation.To(); ok {
		if err := usercareer.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "UserCareer.to": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.CareergroupID(); !ok {
		return &ValidationError{Name: "careergroup", err: errors.New(`ent: missing required edge "UserCareer.careergroup"`)}
	}
	return nil
}

func (ucc *UserCareerCreate) sqlSave(ctx context.Context) (*UserCareer, error) {
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ucc *UserCareerCreate) createSpec() (*UserCareer, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCareer{config: ucc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usercareer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usercareer.FieldID,
			},
		}
	)
	_spec.OnConflict = ucc.conflict
	if value, ok := ucc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercareer.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ucc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercareer.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ucc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercareer.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ucc.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercareer.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := ucc.mutation.To(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercareer.FieldTo,
		})
		_node.To = value
	}
	if nodes := ucc.mutation.CareergroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercareer.CareergroupTable,
			Columns: []string{usercareer.CareergroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareergroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.careergroup_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ucc.mutation.CareerdescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareer.CareerdescriptionsTable,
			Columns: []string{usercareer.CareerdescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercareerdescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ucc.mutation.CareertasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareer.CareertasksTable,
			Columns: []string{usercareer.CareertasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: careertask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ucc.mutation.CareerskillgroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercareer.CareerskillgroupsTable,
			Columns: []string{usercareer.CareerskillgroupsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCareer.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCareerUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ucc *UserCareerCreate) OnConflict(opts ...sql.ConflictOption) *UserCareerUpsertOne {
	ucc.conflict = opts
	return &UserCareerUpsertOne{
		create: ucc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucc *UserCareerCreate) OnConflictColumns(columns ...string) *UserCareerUpsertOne {
	ucc.conflict = append(ucc.conflict, sql.ConflictColumns(columns...))
	return &UserCareerUpsertOne{
		create: ucc,
	}
}

type (
	// UserCareerUpsertOne is the builder for "upsert"-ing
	//  one UserCareer node.
	UserCareerUpsertOne struct {
		create *UserCareerCreate
	}

	// UserCareerUpsert is the "OnConflict" setter.
	UserCareerUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *UserCareerUpsert) SetCreateTime(v time.Time) *UserCareerUpsert {
	u.Set(usercareer.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserCareerUpsert) UpdateCreateTime() *UserCareerUpsert {
	u.SetExcluded(usercareer.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserCareerUpsert) SetUpdateTime(v time.Time) *UserCareerUpsert {
	u.Set(usercareer.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserCareerUpsert) UpdateUpdateTime() *UserCareerUpsert {
	u.SetExcluded(usercareer.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *UserCareerUpsert) SetName(v string) *UserCareerUpsert {
	u.Set(usercareer.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserCareerUpsert) UpdateName() *UserCareerUpsert {
	u.SetExcluded(usercareer.FieldName)
	return u
}

// SetFrom sets the "from" field.
func (u *UserCareerUpsert) SetFrom(v string) *UserCareerUpsert {
	u.Set(usercareer.FieldFrom, v)
	return u
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *UserCareerUpsert) UpdateFrom() *UserCareerUpsert {
	u.SetExcluded(usercareer.FieldFrom)
	return u
}

// SetTo sets the "to" field.
func (u *UserCareerUpsert) SetTo(v string) *UserCareerUpsert {
	u.Set(usercareer.FieldTo, v)
	return u
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *UserCareerUpsert) UpdateTo() *UserCareerUpsert {
	u.SetExcluded(usercareer.FieldTo)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserCareerUpsertOne) UpdateNewValues() *UserCareerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(usercareer.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserCareerUpsertOne) Ignore() *UserCareerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCareerUpsertOne) DoNothing() *UserCareerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCareerCreate.OnConflict
// documentation for more info.
func (u *UserCareerUpsertOne) Update(set func(*UserCareerUpsert)) *UserCareerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCareerUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *UserCareerUpsertOne) SetCreateTime(v time.Time) *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserCareerUpsertOne) UpdateCreateTime() *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *UserCareerUpsertOne) SetUpdateTime(v time.Time) *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserCareerUpsertOne) UpdateUpdateTime() *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserCareerUpsertOne) SetName(v string) *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserCareerUpsertOne) UpdateName() *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateName()
	})
}

// SetFrom sets the "from" field.
func (u *UserCareerUpsertOne) SetFrom(v string) *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *UserCareerUpsertOne) UpdateFrom() *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *UserCareerUpsertOne) SetTo(v string) *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *UserCareerUpsertOne) UpdateTo() *UserCareerUpsertOne {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateTo()
	})
}

// Exec executes the query.
func (u *UserCareerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCareerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCareerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserCareerUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserCareerUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCareerCreateBulk is the builder for creating many UserCareer entities in bulk.
type UserCareerCreateBulk struct {
	config
	builders []*UserCareerCreate
	conflict []sql.ConflictOption
}

// Save creates the UserCareer entities in the database.
func (uccb *UserCareerCreateBulk) Save(ctx context.Context) ([]*UserCareer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UserCareer, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCareerMutation)
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
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UserCareerCreateBulk) SaveX(ctx context.Context) []*UserCareer {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UserCareerCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UserCareerCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCareer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCareerUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uccb *UserCareerCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserCareerUpsertBulk {
	uccb.conflict = opts
	return &UserCareerUpsertBulk{
		create: uccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uccb *UserCareerCreateBulk) OnConflictColumns(columns ...string) *UserCareerUpsertBulk {
	uccb.conflict = append(uccb.conflict, sql.ConflictColumns(columns...))
	return &UserCareerUpsertBulk{
		create: uccb,
	}
}

// UserCareerUpsertBulk is the builder for "upsert"-ing
// a bulk of UserCareer nodes.
type UserCareerUpsertBulk struct {
	create *UserCareerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserCareerUpsertBulk) UpdateNewValues() *UserCareerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(usercareer.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserCareer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserCareerUpsertBulk) Ignore() *UserCareerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCareerUpsertBulk) DoNothing() *UserCareerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCareerCreateBulk.OnConflict
// documentation for more info.
func (u *UserCareerUpsertBulk) Update(set func(*UserCareerUpsert)) *UserCareerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCareerUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *UserCareerUpsertBulk) SetCreateTime(v time.Time) *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *UserCareerUpsertBulk) UpdateCreateTime() *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *UserCareerUpsertBulk) SetUpdateTime(v time.Time) *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserCareerUpsertBulk) UpdateUpdateTime() *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserCareerUpsertBulk) SetName(v string) *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserCareerUpsertBulk) UpdateName() *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateName()
	})
}

// SetFrom sets the "from" field.
func (u *UserCareerUpsertBulk) SetFrom(v string) *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *UserCareerUpsertBulk) UpdateFrom() *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *UserCareerUpsertBulk) SetTo(v string) *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *UserCareerUpsertBulk) UpdateTo() *UserCareerUpsertBulk {
	return u.Update(func(s *UserCareerUpsert) {
		s.UpdateTo()
	})
}

// Exec executes the query.
func (u *UserCareerUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCareerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCareerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCareerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
