package schema

import "entgo.io/ent"

// SkillTag holds the schema definition for the SkillTag entity.
type SkillTag struct {
	ent.Schema
}

// Fields of the SkillTag.
func (SkillTag) Fields() []ent.Field {
	return nil
}

// Edges of the SkillTag.
func (SkillTag) Edges() []ent.Edge {
	return nil
}
