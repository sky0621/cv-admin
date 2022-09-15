package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sky0621/cv-admin/src/interfaces/swagger"
)

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
