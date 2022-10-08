package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/ent/helper"

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

// PostUsersByUserIdCareergroups キャリアグループ新規登録
// (POST /users/{byUserId}/careergroups)
func (s *ServerImpl) PostUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error {
	var userCareerGroup UserCareerGroup
	if err := ctx.Bind(&userCareerGroup); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userCareerGroup.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	var entUserCareerGroup *ent.UserCareerGroup
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		entUserCareerGroup, err = ToEntUserCareerGroupCreate(userCareerGroup, byUserId, s.dbClient.UserCareerGroup.Create()).Save(rCtx)
		if err != nil {
			return err
		}

		if userCareerGroup.Careers == nil {
			return nil
		}
		var builders []*ent.UserCareerCreate
		for _, career := range *userCareerGroup.Careers {
			builders = append(builders, ToEntUserCareerCreate(career, entUserCareerGroup.ID, tx.UserCareer.Create()))
		}
		userCareers, err := tx.UserCareer.CreateBulk(builders...).Save(rCtx)
		if err != nil {
			return err
		}
		entUserCareerGroup.Edges.Careers = userCareers

		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerUserCareerGroup(entUserCareerGroup))
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
