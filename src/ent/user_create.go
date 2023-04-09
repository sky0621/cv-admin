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
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
	"github.com/sky0621/cv-admin/src/ent/usernote"
	"github.com/sky0621/cv-admin/src/ent/userqualification"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetAvatarURL sets the "avatar_url" field.
func (uc *UserCreate) SetAvatarURL(s string) *UserCreate {
	uc.mutation.SetAvatarURL(s)
	return uc
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatarURL(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatarURL(*s)
	}
	return uc
}

// SetBirthdayYear sets the "birthday_year" field.
func (uc *UserCreate) SetBirthdayYear(i int) *UserCreate {
	uc.mutation.SetBirthdayYear(i)
	return uc
}

// SetBirthdayMonth sets the "birthday_month" field.
func (uc *UserCreate) SetBirthdayMonth(i int) *UserCreate {
	uc.mutation.SetBirthdayMonth(i)
	return uc
}

// SetBirthdayDay sets the "birthday_day" field.
func (uc *UserCreate) SetBirthdayDay(i int) *UserCreate {
	uc.mutation.SetBirthdayDay(i)
	return uc
}

// SetJob sets the "job" field.
func (uc *UserCreate) SetJob(s string) *UserCreate {
	uc.mutation.SetJob(s)
	return uc
}

// SetNillableJob sets the "job" field if the given value is not nil.
func (uc *UserCreate) SetNillableJob(s *string) *UserCreate {
	if s != nil {
		uc.SetJob(*s)
	}
	return uc
}

// SetBelongTo sets the "belong_to" field.
func (uc *UserCreate) SetBelongTo(s string) *UserCreate {
	uc.mutation.SetBelongTo(s)
	return uc
}

// SetNillableBelongTo sets the "belong_to" field if the given value is not nil.
func (uc *UserCreate) SetNillableBelongTo(s *string) *UserCreate {
	if s != nil {
		uc.SetBelongTo(*s)
	}
	return uc
}

// SetPr sets the "pr" field.
func (uc *UserCreate) SetPr(s string) *UserCreate {
	uc.mutation.SetPr(s)
	return uc
}

// SetNillablePr sets the "pr" field if the given value is not nil.
func (uc *UserCreate) SetNillablePr(s *string) *UserCreate {
	if s != nil {
		uc.SetPr(*s)
	}
	return uc
}

// AddActivityIDs adds the "activities" edge to the UserActivity entity by IDs.
func (uc *UserCreate) AddActivityIDs(ids ...int) *UserCreate {
	uc.mutation.AddActivityIDs(ids...)
	return uc
}

// AddActivities adds the "activities" edges to the UserActivity entity.
func (uc *UserCreate) AddActivities(u ...*UserActivity) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddActivityIDs(ids...)
}

// AddQualificationIDs adds the "qualifications" edge to the UserQualification entity by IDs.
func (uc *UserCreate) AddQualificationIDs(ids ...int) *UserCreate {
	uc.mutation.AddQualificationIDs(ids...)
	return uc
}

// AddQualifications adds the "qualifications" edges to the UserQualification entity.
func (uc *UserCreate) AddQualifications(u ...*UserQualification) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddQualificationIDs(ids...)
}

// AddCareerGroupIDs adds the "careerGroups" edge to the UserCareerGroup entity by IDs.
func (uc *UserCreate) AddCareerGroupIDs(ids ...int) *UserCreate {
	uc.mutation.AddCareerGroupIDs(ids...)
	return uc
}

// AddCareerGroups adds the "careerGroups" edges to the UserCareerGroup entity.
func (uc *UserCreate) AddCareerGroups(u ...*UserCareerGroup) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddCareerGroupIDs(ids...)
}

// AddNoteIDs adds the "notes" edge to the UserNote entity by IDs.
func (uc *UserCreate) AddNoteIDs(ids ...int) *UserCreate {
	uc.mutation.AddNoteIDs(ids...)
	return uc
}

// AddNotes adds the "notes" edges to the UserNote entity.
func (uc *UserCreate) AddNotes(u ...*UserNote) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddNoteIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uc.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if _, ok := uc.mutation.BirthdayYear(); !ok {
		return &ValidationError{Name: "birthday_year", err: errors.New(`ent: missing required field "User.birthday_year"`)}
	}
	if v, ok := uc.mutation.BirthdayYear(); ok {
		if err := user.BirthdayYearValidator(v); err != nil {
			return &ValidationError{Name: "birthday_year", err: fmt.Errorf(`ent: validator failed for field "User.birthday_year": %w`, err)}
		}
	}
	if _, ok := uc.mutation.BirthdayMonth(); !ok {
		return &ValidationError{Name: "birthday_month", err: errors.New(`ent: missing required field "User.birthday_month"`)}
	}
	if v, ok := uc.mutation.BirthdayMonth(); ok {
		if err := user.BirthdayMonthValidator(v); err != nil {
			return &ValidationError{Name: "birthday_month", err: fmt.Errorf(`ent: validator failed for field "User.birthday_month": %w`, err)}
		}
	}
	if _, ok := uc.mutation.BirthdayDay(); !ok {
		return &ValidationError{Name: "birthday_day", err: errors.New(`ent: missing required field "User.birthday_day"`)}
	}
	if v, ok := uc.mutation.BirthdayDay(); ok {
		if err := user.BirthdayDayValidator(v); err != nil {
			return &ValidationError{Name: "birthday_day", err: fmt.Errorf(`ent: validator failed for field "User.birthday_day": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Job(); ok {
		if err := user.JobValidator(v); err != nil {
			return &ValidationError{Name: "job", err: fmt.Errorf(`ent: validator failed for field "User.job": %w`, err)}
		}
	}
	if v, ok := uc.mutation.BelongTo(); ok {
		if err := user.BelongToValidator(v); err != nil {
			return &ValidationError{Name: "belong_to", err: fmt.Errorf(`ent: validator failed for field "User.belong_to": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Pr(); ok {
		if err := user.PrValidator(v); err != nil {
			return &ValidationError{Name: "pr", err: fmt.Errorf(`ent: validator failed for field "User.pr": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = &value
	}
	if value, ok := uc.mutation.AvatarURL(); ok {
		_spec.SetField(user.FieldAvatarURL, field.TypeString, value)
		_node.AvatarURL = &value
	}
	if value, ok := uc.mutation.BirthdayYear(); ok {
		_spec.SetField(user.FieldBirthdayYear, field.TypeInt, value)
		_node.BirthdayYear = value
	}
	if value, ok := uc.mutation.BirthdayMonth(); ok {
		_spec.SetField(user.FieldBirthdayMonth, field.TypeInt, value)
		_node.BirthdayMonth = value
	}
	if value, ok := uc.mutation.BirthdayDay(); ok {
		_spec.SetField(user.FieldBirthdayDay, field.TypeInt, value)
		_node.BirthdayDay = value
	}
	if value, ok := uc.mutation.Job(); ok {
		_spec.SetField(user.FieldJob, field.TypeString, value)
		_node.Job = &value
	}
	if value, ok := uc.mutation.BelongTo(); ok {
		_spec.SetField(user.FieldBelongTo, field.TypeString, value)
		_node.BelongTo = &value
	}
	if value, ok := uc.mutation.Pr(); ok {
		_spec.SetField(user.FieldPr, field.TypeString, value)
		_node.Pr = &value
	}
	if nodes := uc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ActivitiesTable,
			Columns: []string{user.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useractivity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.QualificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QualificationsTable,
			Columns: []string{user.QualificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userqualification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CareerGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CareerGroupsTable,
			Columns: []string{user.CareerGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercareergroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usernote.FieldID, field.TypeInt),
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
//	client.User.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsert) SetUpdateTime(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdateTime() *UserUpsert {
	u.SetExcluded(user.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// SetNickname sets the "nickname" field.
func (u *UserUpsert) SetNickname(v string) *UserUpsert {
	u.Set(user.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsert) UpdateNickname() *UserUpsert {
	u.SetExcluded(user.FieldNickname)
	return u
}

// ClearNickname clears the value of the "nickname" field.
func (u *UserUpsert) ClearNickname() *UserUpsert {
	u.SetNull(user.FieldNickname)
	return u
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsert) SetAvatarURL(v string) *UserUpsert {
	u.Set(user.FieldAvatarURL, v)
	return u
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsert) UpdateAvatarURL() *UserUpsert {
	u.SetExcluded(user.FieldAvatarURL)
	return u
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsert) ClearAvatarURL() *UserUpsert {
	u.SetNull(user.FieldAvatarURL)
	return u
}

// SetBirthdayYear sets the "birthday_year" field.
func (u *UserUpsert) SetBirthdayYear(v int) *UserUpsert {
	u.Set(user.FieldBirthdayYear, v)
	return u
}

// UpdateBirthdayYear sets the "birthday_year" field to the value that was provided on create.
func (u *UserUpsert) UpdateBirthdayYear() *UserUpsert {
	u.SetExcluded(user.FieldBirthdayYear)
	return u
}

// AddBirthdayYear adds v to the "birthday_year" field.
func (u *UserUpsert) AddBirthdayYear(v int) *UserUpsert {
	u.Add(user.FieldBirthdayYear, v)
	return u
}

// SetBirthdayMonth sets the "birthday_month" field.
func (u *UserUpsert) SetBirthdayMonth(v int) *UserUpsert {
	u.Set(user.FieldBirthdayMonth, v)
	return u
}

// UpdateBirthdayMonth sets the "birthday_month" field to the value that was provided on create.
func (u *UserUpsert) UpdateBirthdayMonth() *UserUpsert {
	u.SetExcluded(user.FieldBirthdayMonth)
	return u
}

// AddBirthdayMonth adds v to the "birthday_month" field.
func (u *UserUpsert) AddBirthdayMonth(v int) *UserUpsert {
	u.Add(user.FieldBirthdayMonth, v)
	return u
}

// SetBirthdayDay sets the "birthday_day" field.
func (u *UserUpsert) SetBirthdayDay(v int) *UserUpsert {
	u.Set(user.FieldBirthdayDay, v)
	return u
}

// UpdateBirthdayDay sets the "birthday_day" field to the value that was provided on create.
func (u *UserUpsert) UpdateBirthdayDay() *UserUpsert {
	u.SetExcluded(user.FieldBirthdayDay)
	return u
}

// AddBirthdayDay adds v to the "birthday_day" field.
func (u *UserUpsert) AddBirthdayDay(v int) *UserUpsert {
	u.Add(user.FieldBirthdayDay, v)
	return u
}

// SetJob sets the "job" field.
func (u *UserUpsert) SetJob(v string) *UserUpsert {
	u.Set(user.FieldJob, v)
	return u
}

// UpdateJob sets the "job" field to the value that was provided on create.
func (u *UserUpsert) UpdateJob() *UserUpsert {
	u.SetExcluded(user.FieldJob)
	return u
}

// ClearJob clears the value of the "job" field.
func (u *UserUpsert) ClearJob() *UserUpsert {
	u.SetNull(user.FieldJob)
	return u
}

// SetBelongTo sets the "belong_to" field.
func (u *UserUpsert) SetBelongTo(v string) *UserUpsert {
	u.Set(user.FieldBelongTo, v)
	return u
}

// UpdateBelongTo sets the "belong_to" field to the value that was provided on create.
func (u *UserUpsert) UpdateBelongTo() *UserUpsert {
	u.SetExcluded(user.FieldBelongTo)
	return u
}

// ClearBelongTo clears the value of the "belong_to" field.
func (u *UserUpsert) ClearBelongTo() *UserUpsert {
	u.SetNull(user.FieldBelongTo)
	return u
}

// SetPr sets the "pr" field.
func (u *UserUpsert) SetPr(v string) *UserUpsert {
	u.Set(user.FieldPr, v)
	return u
}

// UpdatePr sets the "pr" field to the value that was provided on create.
func (u *UserUpsert) UpdatePr() *UserUpsert {
	u.SetExcluded(user.FieldPr)
	return u
}

// ClearPr clears the value of the "pr" field.
func (u *UserUpsert) ClearPr() *UserUpsert {
	u.SetNull(user.FieldPr)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(user.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertOne) SetUpdateTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdateTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetNickname sets the "nickname" field.
func (u *UserUpsertOne) SetNickname(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateNickname() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *UserUpsertOne) ClearNickname() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearNickname()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsertOne) SetAvatarURL(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAvatarURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatarURL()
	})
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsertOne) ClearAvatarURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearAvatarURL()
	})
}

// SetBirthdayYear sets the "birthday_year" field.
func (u *UserUpsertOne) SetBirthdayYear(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayYear(v)
	})
}

// AddBirthdayYear adds v to the "birthday_year" field.
func (u *UserUpsertOne) AddBirthdayYear(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayYear(v)
	})
}

// UpdateBirthdayYear sets the "birthday_year" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBirthdayYear() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayYear()
	})
}

// SetBirthdayMonth sets the "birthday_month" field.
func (u *UserUpsertOne) SetBirthdayMonth(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayMonth(v)
	})
}

// AddBirthdayMonth adds v to the "birthday_month" field.
func (u *UserUpsertOne) AddBirthdayMonth(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayMonth(v)
	})
}

// UpdateBirthdayMonth sets the "birthday_month" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBirthdayMonth() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayMonth()
	})
}

// SetBirthdayDay sets the "birthday_day" field.
func (u *UserUpsertOne) SetBirthdayDay(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayDay(v)
	})
}

// AddBirthdayDay adds v to the "birthday_day" field.
func (u *UserUpsertOne) AddBirthdayDay(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayDay(v)
	})
}

// UpdateBirthdayDay sets the "birthday_day" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBirthdayDay() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayDay()
	})
}

// SetJob sets the "job" field.
func (u *UserUpsertOne) SetJob(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetJob(v)
	})
}

// UpdateJob sets the "job" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateJob() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateJob()
	})
}

// ClearJob clears the value of the "job" field.
func (u *UserUpsertOne) ClearJob() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearJob()
	})
}

// SetBelongTo sets the "belong_to" field.
func (u *UserUpsertOne) SetBelongTo(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBelongTo(v)
	})
}

// UpdateBelongTo sets the "belong_to" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBelongTo() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBelongTo()
	})
}

// ClearBelongTo clears the value of the "belong_to" field.
func (u *UserUpsertOne) ClearBelongTo() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearBelongTo()
	})
}

// SetPr sets the "pr" field.
func (u *UserUpsertOne) SetPr(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPr(v)
	})
}

// UpdatePr sets the "pr" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePr() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePr()
	})
}

// ClearPr clears the value of the "pr" field.
func (u *UserUpsertOne) ClearPr() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPr()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(user.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertBulk) SetUpdateTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdateTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetNickname sets the "nickname" field.
func (u *UserUpsertBulk) SetNickname(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateNickname() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *UserUpsertBulk) ClearNickname() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearNickname()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsertBulk) SetAvatarURL(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAvatarURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatarURL()
	})
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsertBulk) ClearAvatarURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearAvatarURL()
	})
}

// SetBirthdayYear sets the "birthday_year" field.
func (u *UserUpsertBulk) SetBirthdayYear(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayYear(v)
	})
}

// AddBirthdayYear adds v to the "birthday_year" field.
func (u *UserUpsertBulk) AddBirthdayYear(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayYear(v)
	})
}

// UpdateBirthdayYear sets the "birthday_year" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBirthdayYear() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayYear()
	})
}

// SetBirthdayMonth sets the "birthday_month" field.
func (u *UserUpsertBulk) SetBirthdayMonth(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayMonth(v)
	})
}

// AddBirthdayMonth adds v to the "birthday_month" field.
func (u *UserUpsertBulk) AddBirthdayMonth(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayMonth(v)
	})
}

// UpdateBirthdayMonth sets the "birthday_month" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBirthdayMonth() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayMonth()
	})
}

// SetBirthdayDay sets the "birthday_day" field.
func (u *UserUpsertBulk) SetBirthdayDay(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBirthdayDay(v)
	})
}

// AddBirthdayDay adds v to the "birthday_day" field.
func (u *UserUpsertBulk) AddBirthdayDay(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddBirthdayDay(v)
	})
}

// UpdateBirthdayDay sets the "birthday_day" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBirthdayDay() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBirthdayDay()
	})
}

// SetJob sets the "job" field.
func (u *UserUpsertBulk) SetJob(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetJob(v)
	})
}

// UpdateJob sets the "job" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateJob() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateJob()
	})
}

// ClearJob clears the value of the "job" field.
func (u *UserUpsertBulk) ClearJob() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearJob()
	})
}

// SetBelongTo sets the "belong_to" field.
func (u *UserUpsertBulk) SetBelongTo(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBelongTo(v)
	})
}

// UpdateBelongTo sets the "belong_to" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBelongTo() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBelongTo()
	})
}

// ClearBelongTo clears the value of the "belong_to" field.
func (u *UserUpsertBulk) ClearBelongTo() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearBelongTo()
	})
}

// SetPr sets the "pr" field.
func (u *UserUpsertBulk) SetPr(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPr(v)
	})
}

// UpdatePr sets the "pr" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePr() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePr()
	})
}

// ClearPr clears the value of the "pr" field.
func (u *UserUpsertBulk) ClearPr() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPr()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
