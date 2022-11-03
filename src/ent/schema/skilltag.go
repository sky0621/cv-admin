package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SkillTag holds the schema definition for the SkillTag entity.
type SkillTag struct {
	ent.Schema
}

// Fields of the SkillTag.
func (SkillTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)).Unique(),
		field.String("key").NotEmpty().Validate(maxRuneCount(40)).Unique(),
	}
}
