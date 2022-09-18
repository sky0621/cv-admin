package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Validate(maxRuneCount(20)).Unique().Comment("ユーザーを一意に識別するキー。\n各URLのパスパラメーターに使う。"),
		field.String("name").NotEmpty().Validate(maxRuneCount(100)),
		field.String("nickname").Validate(rangeRuneCount(1, 100)).Optional().Nillable(),
		field.String("avatar_url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
		field.Int("birthday_year").Range(1900, 3000),
		field.Int("birthday_month").Range(1, 12),
		field.Int("birthday_day").Range(1, 31),
		field.String("job").Validate(rangeRuneCount(1, 400)).Optional().Nillable(),
		field.String("belong_to").Validate(rangeRuneCount(1, 400)).Optional().Nillable(),
		field.String("pr").Validate(rangeRuneCount(1, 4000)).Optional().Nillable(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}