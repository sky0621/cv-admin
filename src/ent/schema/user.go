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
		field.String("nickname").Optional().Nillable(),
		field.String("avatar_url").Optional().Nillable(),
		field.Int("birthday_year").Min(1900),
		field.Int("birthday_month").Range(1, 12),
		field.Int("birthday_day").Range(1, 31),
		field.String("job").Optional().Nillable(),
		field.String("belong_to").Optional().Nillable(),
		field.String("pr").Optional().Nillable(),
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
