package schema

import "entgo.io/ent"

// UserQualification holds the schema definition for the UserQualification entity.
type UserQualification struct {
	ent.Schema
}

// Fields of the UserQualification.
func (UserQualification) Fields() []ent.Field {
	return nil
}

// Edges of the UserQualification.
func (UserQualification) Edges() []ent.Edge {
	return nil
}
