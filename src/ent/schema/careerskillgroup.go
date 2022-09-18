package schema

import "entgo.io/ent"

// CareerSkillGroup holds the schema definition for the CareerSkillGroup entity.
type CareerSkillGroup struct {
	ent.Schema
}

// Fields of the CareerSkillGroup.
func (CareerSkillGroup) Fields() []ent.Field {
	return nil
}

// Edges of the CareerSkillGroup.
func (CareerSkillGroup) Edges() []ent.Edge {
	return nil
}