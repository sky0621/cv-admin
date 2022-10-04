package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/ent"

	"github.com/labstack/echo/v4"
)

// GetUsersByUserIdCareergroups キャリアグループ群取得
// (GET /users/{byUserId}/careergroups)
func (s *ServerImpl) GetUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error {
	entUser, err := s.getUserByUserIdWithXXX(ctx, byUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithCareerGroups()
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserCareerGroups(entUser.Edges.CareerGroups))
}

// キャリアグループ新規登録
// (POST /users/{byUserId}/careergroups)
func (s *ServerImpl) PostUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error {
	return ctx.String(http.StatusOK, "")
}

// キャリアグループ削除
// (DELETE /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *ServerImpl) DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	return ctx.String(http.StatusOK, "")
}

// キャリアグループ更新
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *ServerImpl) PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	return ctx.String(http.StatusOK, "")
}

// キャリアグループ内キャリア群最新化
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
func (s *ServerImpl) PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	return ctx.String(http.StatusOK, "")
}
