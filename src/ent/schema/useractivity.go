package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const userActivitiesRef = "activities"

// UserActivity holds the schema definition for the UserActivity entity.
type UserActivity struct {
	ent.Schema
}

// Fields of the UserActivity.
func (UserActivity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)),
		field.String("url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
		field.String("icon").Validate(maxRuneCount(40)).Optional().Nillable(),
	}
}

func (UserActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserActivity.
func (UserActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userActivitiesRef).Unique().Required(),
	}
}
