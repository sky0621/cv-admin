package schema

import "entgo.io/ent"

// UserNoteItem holds the schema definition for the UserNoteItem entity.
type UserNoteItem struct {
	ent.Schema
}

// Fields of the UserNoteItem.
func (UserNoteItem) Fields() []ent.Field {
	return nil
}

// Edges of the UserNoteItem.
func (UserNoteItem) Edges() []ent.Edge {
	return nil
}
