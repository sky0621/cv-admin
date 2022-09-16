package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/swagger"

	"github.com/labstack/echo/v4"
)

// キャリアグループ群取得
// (GET /users/{byUserKey}/careergroups)
func (s *ServerImpl) GetUsersByUserKeyCareergroups(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// キャリアグループ新規登録
// (POST /users/{byUserKey}/careergroups)
func (s *ServerImpl) PostUsersByUserKeyCareergroups(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// キャリアグループ削除
// (DELETE /users/{byUserKey}/careergroups/{byCid})
func (s *ServerImpl) DeleteUsersByUserKeyCareergroupsByCid(ctx echo.Context, byUserKey swagger.UserKey, byCid swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// キャリアグループ更新
// (PUT /users/{byUserKey}/careergroups/{byCid})
func (s *ServerImpl) PutUsersByUserKeyCareergroupsByCid(ctx echo.Context, byUserKey swagger.UserKey, byCid swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// キャリアグループ内キャリア群最新化
// (PUT /users/{byUserKey}/careergroups/{byCid}/careers)
func (s *ServerImpl) PutUsersByUserKeyCareergroupsByCidCareers(ctx echo.Context, byUserKey swagger.UserKey, byCid swagger.CareerGroupId) error {
	return ctx.String(http.StatusOK, byUserKey)
}
