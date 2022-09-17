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
		field.String("key").NotEmpty().Unique().Comment("ユーザー識別キー"),
		field.String("name"),
		field.String("nickname"),
		field.String("avatar_url"),
		field.Int("birthday_year"),
		field.Int("birthday_month"),
		field.Int("birthday_day"),
		field.String("job"),
		field.String("belong_to"),
		field.String("pr"),
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
