package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	careerSkillGroupEdgeName = "careerSkillGroup"
	careerSkillGroupsRef     = "careerSkillGroups"
)

// CareerSkillGroup holds the schema definition for the CareerSkillGroup entity.
type CareerSkillGroup struct {
	ent.Schema
}

// Fields of the CareerSkillGroup.
func (CareerSkillGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").NotEmpty().Validate(maxRuneCount(100)),
	}
}

func (CareerSkillGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the CareerSkillGroup.
func (CareerSkillGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userCareerEdgeName, UserCareer.Type).Ref(careerSkillGroupsRef).Unique().Required(),

		edge.To(careerSkillsRef, CareerSkill.Type).StorageKey(edge.Column("career_skill_group_id")),
	}
}
