package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	userNoteEdgeName = "note"
	userNotesRef     = "notes"
)

// UserNote holds the schema definition for the UserNote entity.
type UserNote struct {
	ent.Schema
}

// Fields of the UserNote.
func (UserNote) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").NotEmpty().Validate(maxRuneCount(80)),
		field.String("memo").Validate(maxRuneCount(400)).Optional().Nillable(),
	}
}

func (UserNote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserNote.
func (UserNote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userNotesRef).Unique().Required(),

		edge.To(userNoteItemsRef, UserNoteItem.Type).StorageKey(edge.Column("user_note_id")),
	}
}
