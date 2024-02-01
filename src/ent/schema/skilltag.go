package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const skillTagEdgeName = "skillTag"

// SkillTag holds the schema definition for the SkillTag entity.
type SkillTag struct {
	ent.Schema
}

// Fields of the SkillTag.
func (SkillTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)).Unique(),
		field.String("code").NotEmpty().Validate(maxRuneCount(40)).Unique(),
	}
}

// Edges of the SkillTag.
func (SkillTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(skillsRef, Skill.Type).StorageKey(edge.Column("tag_id")),
	}
}
