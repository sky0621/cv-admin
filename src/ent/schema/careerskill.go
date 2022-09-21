package schema

import "entgo.io/ent"

// CareerSkill holds the schema definition for the CareerSkill entity.
type CareerSkill struct {
	ent.Schema
}

// Fields of the CareerSkill.
func (CareerSkill) Fields() []ent.Field {
	return nil
}

func (CareerSkill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the CareerSkill.
func (CareerSkill) Edges() []ent.Edge {
	return nil
}
