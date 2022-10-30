package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 注釈群取得
// (GET /users/{byUserId}/notes)
func (s *ServerImpl) GetUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error {
	return ctx.String(http.StatusOK, "")
}

// PostUsersByUserIdNotes 注釈新規登録
// 注釈新規登録
// (POST /users/{byUserId}/notes)
func (s *ServerImpl) PostUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error {
	var userNote UserNote
	if err := ctx.Bind(&userNote); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userNote.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	var entUserNote *ent.UserNote
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		entUserNote, err = ToEntUserNoteCreate(userNote, byUserId, tx.UserNote.Create()).Save(rCtx)
		if err != nil {
			return err
		}

		if userNote.Items != nil {
			var noteItemCreates []*ent.UserNoteItemCreate
			for _, noteItem := range *userNote.Items {
				noteItemCreates = append(noteItemCreates, ToEntUserNoteItemCreate(noteItem, entUserNote.ID, tx.UserNoteItem.Create()))
			}
			entUserNoteItems, err := tx.UserNoteItem.CreateBulk(noteItemCreates...).Save(rCtx)
			if err != nil {
				return err
			}
			entUserNote.Edges.NoteItems = entUserNoteItems
		}

		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerUserNote(entUserNote))
}

// 注釈削除
// (DELETE /users/{byUserId}/notes/{byNoteId})
func (s *ServerImpl) DeleteUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈更新
// (PUT /users/{byUserId}/notes/{byNoteId})
func (s *ServerImpl) PutUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈内要素群最新化
// (PUT /users/{byUserId}/notes/{byNoteId}/items)
func (s *ServerImpl) PutUsersByUserIdNotesByNoteIdItems(ctx echo.Context, byUserId UserId, byNoteId NoteId) error {
	return ctx.String(http.StatusOK, "")
}
