package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/generated/swagger"
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

func main() {
	si := &ServerImpl{}
	e := echo.New()
	swagger.RegisterHandlers(e, si)
	http.ListenAndServe(":3000", e)
}
