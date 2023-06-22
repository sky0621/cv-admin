package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PostUsersByUserIdNotes 注釈新規登録
// (POST /users/{byUserId}/notes)
func (s *strictServerImpl) PostUsersByUserIdNotes(ctx context.Context, request PostUsersByUserIdNotesRequestObject) (PostUsersByUserIdNotesResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userNote := *request.Body

	if err := userNote.Validate(); err != nil {
		return PostUsersByUserIdNotes400JSONResponse{n400("validation failed")}, err
	}

	var entUserNote *ent.UserNote
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		entUserNote, err = ToEntUserNoteCreate(userNote, request.ByUserId, tx.UserNote.Create()).Save(ctx)
		if err != nil {
			return err
		}

		if userNote.Items != nil {
			var noteItemCreates []*ent.UserNoteItemCreate
			for _, noteItem := range *userNote.Items {
				noteItemCreates = append(noteItemCreates, ToEntUserNoteItemCreate(noteItem, entUserNote.ID, tx.UserNoteItem.Create()))
			}
			entUserNoteItems, err := tx.UserNoteItem.CreateBulk(noteItemCreates...).Save(ctx)
			if err != nil {
				return err
			}
			entUserNote.Edges.NoteItems = entUserNoteItems
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return PostUsersByUserIdNotes201JSONResponse(ToSwaggerUserNote(entUserNote)), nil
}
