package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	userCareerGroupEdgeName = "careergroup"
	userCareerGroupsRef     = "careergroups"
)

// UserCareerGroup holds the schema definition for the UserCareerGroup entity.
type UserCareerGroup struct {
	ent.Schema
}

// Fields of the UserCareerGroup.
func (UserCareerGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").NotEmpty().Validate(maxRuneCount(100)),
	}
}

func (UserCareerGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserCareerGroup.
func (UserCareerGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userCareerGroupsRef).Unique().Required(),

		edge.To(userCareersRef, UserCareer.Type).StorageKey(edge.Column("careergroup_id")),
	}
}
