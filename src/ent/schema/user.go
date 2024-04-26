package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const userEdgeName = "user"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
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
	return []ent.Edge{
		edge.To(userActivitiesRef, UserActivity.Type).StorageKey(edge.Column("user_id")),
		edge.To(userQualificationsRef, UserQualification.Type).StorageKey(edge.Column("user_id")),
		edge.To(userCareerGroupsRef, UserCareerGroup.Type).StorageKey(edge.Column("user_id")),
		edge.To(userNotesRef, UserNote.Type).StorageKey(edge.Column("user_id")),
		edge.To(userAppealsRef, UserAppeal.Type).StorageKey(edge.Column("user_id")),
		edge.To(userSolutionsRef, UserSolution.Type).StorageKey(edge.Column("user_id")),
	}
}
