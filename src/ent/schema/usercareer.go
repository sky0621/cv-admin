package schema

import "entgo.io/ent"

// UserCareer holds the schema definition for the UserCareer entity.
type UserCareer struct {
	ent.Schema
}

// Fields of the UserCareer.
func (UserCareer) Fields() []ent.Field {
	return nil
}

func (UserCareer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserCareer.
func (UserCareer) Edges() []ent.Edge {
	return nil
}
