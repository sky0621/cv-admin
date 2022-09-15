package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/adapter/controller/swagger"

	"github.com/labstack/echo/v4"
)

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
