package schema

import "entgo.io/ent"

// UserCareerGroup holds the schema definition for the UserCareerGroup entity.
type UserCareerGroup struct {
	ent.Schema
}

// Fields of the UserCareerGroup.
func (UserCareerGroup) Fields() []ent.Field {
	return nil
}

// Edges of the UserCareerGroup.
func (UserCareerGroup) Edges() []ent.Edge {
	return nil
}
