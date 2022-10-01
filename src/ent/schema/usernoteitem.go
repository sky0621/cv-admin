package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const userNoteItemsRef = "noteItems"

// UserNoteItem holds the schema definition for the UserNoteItem entity.
type UserNoteItem struct {
	ent.Schema
}

// Fields of the UserNoteItem.
func (UserNoteItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty().Validate(maxRuneCount(400)),
	}
}

func (UserNoteItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserNoteItem.
func (UserNoteItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userNoteEdgeName, UserNote.Type).Ref(userNoteItemsRef).Unique().Required(),
	}
}
