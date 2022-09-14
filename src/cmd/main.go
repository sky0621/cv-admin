package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/interfaces/swagger"
)

type ServerImpl struct {
}

// ユーザー新規登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "")
}

// 指定ユーザー削除
// (DELETE /users/{userId})
func (s *ServerImpl) DeleteUsersUserId(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// アクティビティ群取得
// (GET /users/{userId}/activities)
func (s *ServerImpl) GetUsersUserIdActivities(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// アクティビティ群更新
// (PUT /users/{userId}/activities)
func (s *ServerImpl) PutUsersUserIdActivities(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// 属性取得
// (GET /users/{userId}/attributes)
func (s *ServerImpl) GetUsersUserIdAttributes(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// 属性更新
// (PUT /users/{userId}/attributes)
func (s *ServerImpl) PutUsersUserIdAttributes(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// キャリアグループ群取得
// (GET /users/{userId}/careergroups)
func (s *ServerImpl) GetUsersUserIdCareergroups(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// キャリアグループ新規登録
// (POST /users/{userId}/careergroups)
func (s *ServerImpl) PostUsersUserIdCareergroups(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// キャリアグループ削除
// (DELETE /users/{userId}/careergroups/{careerGroupId})
func (s *ServerImpl) DeleteUsersUserIdCareergroupsCareerGroupId(ctx echo.Context, userId swagger.UserId, careerGroupId swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, userId)
}

// キャリアグループ更新
// (PUT /users/{userId}/careergroups/{careerGroupId})
func (s *ServerImpl) PutUsersUserIdCareergroupsCareerGroupId(ctx echo.Context, userId swagger.UserId, careerGroupId swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, userId)
}

// キャリアグループ内キャリア群最新化
// (PUT /users/{userId}/careergroups/{careerGroupId}/careers)
func (s *ServerImpl) PutUsersUserIdCareergroupsCareerGroupIdCareers(ctx echo.Context, userId swagger.UserId, careerGroupId swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, userId)
}

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

// 資格情報群取得
// (GET /users/{userId}/qualifications)
func (s *ServerImpl) GetUsersUserIdQualifications(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

// 資格情報群更新
// (PUT /users/{userId}/qualifications)
func (s *ServerImpl) PutUsersUserIdQualifications(ctx echo.Context, userId swagger.UserId) error {
	return ctx.String(http.StatusOK, userId)
}

func main() {
	si := &ServerImpl{}
	e := echo.New()
	swagger.RegisterHandlers(e, si)
	http.ListenAndServe(":3000", e)
}
