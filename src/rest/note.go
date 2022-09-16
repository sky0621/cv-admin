package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/swagger"

	"github.com/labstack/echo/v4"
)

// 注釈群取得
// (GET /users/{byUserKey}/notes)
func (s *ServerImpl) GetUsersByUserKeyNotes(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 注釈新規登録
// (POST /users/{byUserKey}/notes)
func (s *ServerImpl) PostUsersByUserKeyNotes(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 注釈削除
// (DELETE /users/{byUserKey}/notes/{byNid})
func (s *ServerImpl) DeleteUsersByUserKeyNotesByNid(ctx echo.Context, byUserKey swagger.UserKey, byNid swagger.NoteId) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 注釈更新
// (PUT /users/{byUserKey}/notes/{byNid})
func (s *ServerImpl) PutUsersByUserKeyNotesByNid(ctx echo.Context, byUserKey swagger.UserKey, byNid swagger.NoteId) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 注釈内要素群最新化
// (PUT /users/{byUserKey}/notes/{byNid}/items)
func (s *ServerImpl) PutUsersByUserKeyNotesByNidItems(ctx echo.Context, byUserKey swagger.UserKey, byNid swagger.NoteId) error {
	return ctx.String(http.StatusOK, byUserKey)
}
