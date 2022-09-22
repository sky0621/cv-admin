package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/swagger"

	"github.com/labstack/echo/v4"
)

// 注釈群取得
// (GET /users/{byUserId}/notes)
func (s *ServerImpl) GetUsersByUserIdNotes(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈新規登録
// (POST /users/{byUserId}/notes)
func (s *ServerImpl) PostUsersByUserIdNotes(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈削除
// (DELETE /users/{byUserId}/notes/{byNoteId})
func (s *ServerImpl) DeleteUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId swagger.UserId, byNoteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈更新
// (PUT /users/{byUserId}/notes/{byNoteId})
func (s *ServerImpl) PutUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId swagger.UserId, byNoteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, "")
}

// 注釈内要素群最新化
// (PUT /users/{byUserId}/notes/{byNoteId}/items)
func (s *ServerImpl) PutUsersByUserIdNotesByNoteIdItems(ctx echo.Context, byUserId swagger.UserId, byNoteId swagger.NoteId) error {
	return ctx.String(http.StatusOK, "")
}
