package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const careerSkillsRef = "careerSkills"

// CareerSkill holds the schema definition for the CareerSkill entity.
type CareerSkill struct {
	ent.Schema
}

// Fields of the CareerSkill.
func (CareerSkill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)),
		field.String("url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
		field.String("version").Validate(rangeRuneCount(1, 40)).Optional().Nillable(),
	}
}

func (CareerSkill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the CareerSkill.
func (CareerSkill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(careerSkillGroupEdgeName, CareerSkillGroup.Type).Ref(careerSkillsRef).Unique().Required(),
	}
}
