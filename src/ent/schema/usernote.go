package schema

import "entgo.io/ent"

// UserNote holds the schema definition for the UserNote entity.
type UserNote struct {
	ent.Schema
}

// Fields of the UserNote.
func (UserNote) Fields() []ent.Field {
	return nil
}

func (UserNote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserNote.
func (UserNote) Edges() []ent.Edge {
	return nil
}
