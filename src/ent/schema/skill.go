package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const skillEdgeName = "skill"
const skillsRef = "skills"

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)).Unique(),
		field.String("url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
	}
}

func (Skill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(skillTagEdgeName, SkillTag.Type).Ref(skillsRef).Unique().Required(),

		edge.To(careerSkillsRef, CareerSkill.Type).StorageKey(edge.Column("skill_id")),
	}
}
