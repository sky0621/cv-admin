package schema

import "entgo.io/ent"

// UserActivity holds the schema definition for the UserActivity entity.
type UserActivity struct {
	ent.Schema
}

// Fields of the UserActivity.
func (UserActivity) Fields() []ent.Field {
	return nil
}

// Edges of the UserActivity.
func (UserActivity) Edges() []ent.Edge {
	return nil
}
