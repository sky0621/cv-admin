package rest

import "github.com/sky0621/cv-admin/src/ent"

func ToSwaggerUserNoteItem(entNoteItem *ent.UserNoteItem) UserNoteItem {
	if entNoteItem == nil {
		return UserNoteItem{}
	}
	return UserNoteItem{
		Text: &entNoteItem.Text,
	}
}

func ToSwaggerUserNoteItems(entNoteItems []*ent.UserNoteItem) *[]UserNoteItem {
	if entNoteItems == nil {
		return nil
	}
	var noteItems []UserNoteItem
	for _, entNoteItem := range entNoteItems {
		noteItems = append(noteItems, ToSwaggerUserNoteItem(entNoteItem))
	}
	return &noteItems
}

func ToSwaggerUserNote(entUserNote *ent.UserNote) UserNote {
	if entUserNote == nil {
		return UserNote{}
	}
	un := UserNote{
		Label: &entUserNote.Label,
		Memo:  entUserNote.Memo,
	}
	if entUserNote.Edges.NoteItems == nil {
		return un
	}

	un.Items = ToSwaggerUserNoteItems(entUserNote.Edges.NoteItems)
	return un
}

func ToEntUserNoteCreate(u UserNote, userID int, c *ent.UserNoteCreate) *ent.UserNoteCreate {
	return c.
		SetUserID(userID).
		SetLabel(*u.Label).
		SetNillableMemo(u.Memo)
}

func ToEntUserNoteItemCreate(u UserNoteItem, noteID int, c *ent.UserNoteItemCreate) *ent.UserNoteItemCreate {
	return c.
		SetNoteID(noteID).
		SetText(*u.Text)
}
