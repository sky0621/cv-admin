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
	"github.com/sky0621/cv-admin/src/ent/predicate"
	"github.com/sky0621/cv-admin/src/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetKey sets the "key" field.
func (uu *UserUpdate) SetKey(s string) *UserUpdate {
	uu.mutation.SetKey(s)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *UserUpdate) SetAvatarURL(s string) *UserUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarURL(*s)
	}
	return uu
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uu *UserUpdate) ClearAvatarURL() *UserUpdate {
	uu.mutation.ClearAvatarURL()
	return uu
}

// SetBirthdayYear sets the "birthday_year" field.
func (uu *UserUpdate) SetBirthdayYear(i int) *UserUpdate {
	uu.mutation.ResetBirthdayYear()
	uu.mutation.SetBirthdayYear(i)
	return uu
}

// AddBirthdayYear adds i to the "birthday_year" field.
func (uu *UserUpdate) AddBirthdayYear(i int) *UserUpdate {
	uu.mutation.AddBirthdayYear(i)
	return uu
}

// SetBirthdayMonth sets the "birthday_month" field.
func (uu *UserUpdate) SetBirthdayMonth(i int) *UserUpdate {
	uu.mutation.ResetBirthdayMonth()
	uu.mutation.SetBirthdayMonth(i)
	return uu
}

// AddBirthdayMonth adds i to the "birthday_month" field.
func (uu *UserUpdate) AddBirthdayMonth(i int) *UserUpdate {
	uu.mutation.AddBirthdayMonth(i)
	return uu
}

// SetBirthdayDay sets the "birthday_day" field.
func (uu *UserUpdate) SetBirthdayDay(i int) *UserUpdate {
	uu.mutation.ResetBirthdayDay()
	uu.mutation.SetBirthdayDay(i)
	return uu
}

// AddBirthdayDay adds i to the "birthday_day" field.
func (uu *UserUpdate) AddBirthdayDay(i int) *UserUpdate {
	uu.mutation.AddBirthdayDay(i)
	return uu
}

// SetJob sets the "job" field.
func (uu *UserUpdate) SetJob(s string) *UserUpdate {
	uu.mutation.SetJob(s)
	return uu
}

// SetNillableJob sets the "job" field if the given value is not nil.
func (uu *UserUpdate) SetNillableJob(s *string) *UserUpdate {
	if s != nil {
		uu.SetJob(*s)
	}
	return uu
}

// ClearJob clears the value of the "job" field.
func (uu *UserUpdate) ClearJob() *UserUpdate {
	uu.mutation.ClearJob()
	return uu
}

// SetBelongTo sets the "belong_to" field.
func (uu *UserUpdate) SetBelongTo(s string) *UserUpdate {
	uu.mutation.SetBelongTo(s)
	return uu
}

// SetNillableBelongTo sets the "belong_to" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBelongTo(s *string) *UserUpdate {
	if s != nil {
		uu.SetBelongTo(*s)
	}
	return uu
}

// ClearBelongTo clears the value of the "belong_to" field.
func (uu *UserUpdate) ClearBelongTo() *UserUpdate {
	uu.mutation.ClearBelongTo()
	return uu
}

// SetPr sets the "pr" field.
func (uu *UserUpdate) SetPr(s string) *UserUpdate {
	uu.mutation.SetPr(s)
	return uu
}

// SetNillablePr sets the "pr" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePr(s *string) *UserUpdate {
	if s != nil {
		uu.SetPr(*s)
	}
	return uu
}

// ClearPr clears the value of the "pr" field.
func (uu *UserUpdate) ClearPr() *UserUpdate {
	uu.mutation.ClearPr()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Key(); ok {
		if err := user.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "User.key": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BirthdayYear(); ok {
		if err := user.BirthdayYearValidator(v); err != nil {
			return &ValidationError{Name: "birthday_year", err: fmt.Errorf(`ent: validator failed for field "User.birthday_year": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BirthdayMonth(); ok {
		if err := user.BirthdayMonthValidator(v); err != nil {
			return &ValidationError{Name: "birthday_month", err: fmt.Errorf(`ent: validator failed for field "User.birthday_month": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BirthdayDay(); ok {
		if err := user.BirthdayDayValidator(v); err != nil {
			return &ValidationError{Name: "birthday_day", err: fmt.Errorf(`ent: validator failed for field "User.birthday_day": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Job(); ok {
		if err := user.JobValidator(v); err != nil {
			return &ValidationError{Name: "job", err: fmt.Errorf(`ent: validator failed for field "User.job": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BelongTo(); ok {
		if err := user.BelongToValidator(v); err != nil {
			return &ValidationError{Name: "belong_to", err: fmt.Errorf(`ent: validator failed for field "User.belong_to": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Pr(); ok {
		if err := user.PrValidator(v); err != nil {
			return &ValidationError{Name: "pr", err: fmt.Errorf(`ent: validator failed for field "User.pr": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldKey,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if uu.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uu.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uu.mutation.BirthdayYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayYear,
		})
	}
	if value, ok := uu.mutation.AddedBirthdayYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayYear,
		})
	}
	if value, ok := uu.mutation.BirthdayMonth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayMonth,
		})
	}
	if value, ok := uu.mutation.AddedBirthdayMonth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayMonth,
		})
	}
	if value, ok := uu.mutation.BirthdayDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayDay,
		})
	}
	if value, ok := uu.mutation.AddedBirthdayDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayDay,
		})
	}
	if value, ok := uu.mutation.Job(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldJob,
		})
	}
	if uu.mutation.JobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldJob,
		})
	}
	if value, ok := uu.mutation.BelongTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBelongTo,
		})
	}
	if uu.mutation.BelongToCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBelongTo,
		})
	}
	if value, ok := uu.mutation.Pr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPr,
		})
	}
	if uu.mutation.PrCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPr,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetKey sets the "key" field.
func (uuo *UserUpdateOne) SetKey(s string) *UserUpdateOne {
	uuo.mutation.SetKey(s)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *UserUpdateOne) SetAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarURL(*s)
	}
	return uuo
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uuo *UserUpdateOne) ClearAvatarURL() *UserUpdateOne {
	uuo.mutation.ClearAvatarURL()
	return uuo
}

// SetBirthdayYear sets the "birthday_year" field.
func (uuo *UserUpdateOne) SetBirthdayYear(i int) *UserUpdateOne {
	uuo.mutation.ResetBirthdayYear()
	uuo.mutation.SetBirthdayYear(i)
	return uuo
}

// AddBirthdayYear adds i to the "birthday_year" field.
func (uuo *UserUpdateOne) AddBirthdayYear(i int) *UserUpdateOne {
	uuo.mutation.AddBirthdayYear(i)
	return uuo
}

// SetBirthdayMonth sets the "birthday_month" field.
func (uuo *UserUpdateOne) SetBirthdayMonth(i int) *UserUpdateOne {
	uuo.mutation.ResetBirthdayMonth()
	uuo.mutation.SetBirthdayMonth(i)
	return uuo
}

// AddBirthdayMonth adds i to the "birthday_month" field.
func (uuo *UserUpdateOne) AddBirthdayMonth(i int) *UserUpdateOne {
	uuo.mutation.AddBirthdayMonth(i)
	return uuo
}

// SetBirthdayDay sets the "birthday_day" field.
func (uuo *UserUpdateOne) SetBirthdayDay(i int) *UserUpdateOne {
	uuo.mutation.ResetBirthdayDay()
	uuo.mutation.SetBirthdayDay(i)
	return uuo
}

// AddBirthdayDay adds i to the "birthday_day" field.
func (uuo *UserUpdateOne) AddBirthdayDay(i int) *UserUpdateOne {
	uuo.mutation.AddBirthdayDay(i)
	return uuo
}

// SetJob sets the "job" field.
func (uuo *UserUpdateOne) SetJob(s string) *UserUpdateOne {
	uuo.mutation.SetJob(s)
	return uuo
}

// SetNillableJob sets the "job" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableJob(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetJob(*s)
	}
	return uuo
}

// ClearJob clears the value of the "job" field.
func (uuo *UserUpdateOne) ClearJob() *UserUpdateOne {
	uuo.mutation.ClearJob()
	return uuo
}

// SetBelongTo sets the "belong_to" field.
func (uuo *UserUpdateOne) SetBelongTo(s string) *UserUpdateOne {
	uuo.mutation.SetBelongTo(s)
	return uuo
}

// SetNillableBelongTo sets the "belong_to" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBelongTo(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBelongTo(*s)
	}
	return uuo
}

// ClearBelongTo clears the value of the "belong_to" field.
func (uuo *UserUpdateOne) ClearBelongTo() *UserUpdateOne {
	uuo.mutation.ClearBelongTo()
	return uuo
}

// SetPr sets the "pr" field.
func (uuo *UserUpdateOne) SetPr(s string) *UserUpdateOne {
	uuo.mutation.SetPr(s)
	return uuo
}

// SetNillablePr sets the "pr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePr(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPr(*s)
	}
	return uuo
}

// ClearPr clears the value of the "pr" field.
func (uuo *UserUpdateOne) ClearPr() *UserUpdateOne {
	uuo.mutation.ClearPr()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Key(); ok {
		if err := user.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "User.key": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BirthdayYear(); ok {
		if err := user.BirthdayYearValidator(v); err != nil {
			return &ValidationError{Name: "birthday_year", err: fmt.Errorf(`ent: validator failed for field "User.birthday_year": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BirthdayMonth(); ok {
		if err := user.BirthdayMonthValidator(v); err != nil {
			return &ValidationError{Name: "birthday_month", err: fmt.Errorf(`ent: validator failed for field "User.birthday_month": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BirthdayDay(); ok {
		if err := user.BirthdayDayValidator(v); err != nil {
			return &ValidationError{Name: "birthday_day", err: fmt.Errorf(`ent: validator failed for field "User.birthday_day": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Job(); ok {
		if err := user.JobValidator(v); err != nil {
			return &ValidationError{Name: "job", err: fmt.Errorf(`ent: validator failed for field "User.job": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BelongTo(); ok {
		if err := user.BelongToValidator(v); err != nil {
			return &ValidationError{Name: "belong_to", err: fmt.Errorf(`ent: validator failed for field "User.belong_to": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Pr(); ok {
		if err := user.PrValidator(v); err != nil {
			return &ValidationError{Name: "pr", err: fmt.Errorf(`ent: validator failed for field "User.pr": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldKey,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if uuo.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uuo.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uuo.mutation.BirthdayYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayYear,
		})
	}
	if value, ok := uuo.mutation.AddedBirthdayYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayYear,
		})
	}
	if value, ok := uuo.mutation.BirthdayMonth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayMonth,
		})
	}
	if value, ok := uuo.mutation.AddedBirthdayMonth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayMonth,
		})
	}
	if value, ok := uuo.mutation.BirthdayDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayDay,
		})
	}
	if value, ok := uuo.mutation.AddedBirthdayDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBirthdayDay,
		})
	}
	if value, ok := uuo.mutation.Job(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldJob,
		})
	}
	if uuo.mutation.JobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldJob,
		})
	}
	if value, ok := uuo.mutation.BelongTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBelongTo,
		})
	}
	if uuo.mutation.BelongToCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBelongTo,
		})
	}
	if value, ok := uuo.mutation.Pr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPr,
		})
	}
	if uuo.mutation.PrCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPr,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}