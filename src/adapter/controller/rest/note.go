package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/adapter/controller/swagger"

	"github.com/labstack/echo/v4"
)

// 注釈群取得
// (GET /users/{userId}/notes)
func (s *ServerImpl) GetUsersUserIdNotes(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// 注釈新規登録
// (POST /users/{userId}/notes)
func (s *ServerImpl) PostUsersUserIdNotes(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// 注釈削除
// (DELETE /users/{userId}/notes/{noteId})
func (s *ServerImpl) DeleteUsersUserIdNotesNoteId(ctx echo.Context, userId swagger.UserId, noteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, noteId)
}

// 注釈更新
// (PUT /users/{userId}/notes/{noteId})
func (s *ServerImpl) PutUsersUserIdNotesNoteId(ctx echo.Context, userId swagger.UserId, noteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, noteId)
}

// 注釈内要素群更新
// (PUT /users/{userId}/notes/{noteId}/items)
func (s *ServerImpl) PutUsersUserIdNotesNoteIdItems(ctx echo.Context, userId swagger.UserId, noteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, noteId)
}
