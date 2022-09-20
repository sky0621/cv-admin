package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/rest/helper"
	"github.com/sky0621/cv-admin/src/swagger"
)

// ユーザー新規登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := helper.ConvertUserAttribute(userAttribute, s.dbClient.User.Create()).
		Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, helper.ConvertEntUserAttribute(entUser))
}

// 指定ユーザー削除
// (DELETE /users/{byUserKey})
func (s *ServerImpl) DeleteUsersByUserKey(ctx echo.Context, byUserKey swagger.UserKey) error {
	cnt, err := s.dbClient.User.Delete().Where(user.Key(byUserKey)).Exec(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}
	if cnt != 1 {
		return sendClientError(ctx, http.StatusNotFound, "user is none")
	}
	return ctx.NoContent(http.StatusNoContent)
}

// 属性取得
// (GET /users/{byUserKey}/attributes)
func (s *ServerImpl) GetUsersByUserKeyAttributes(ctx echo.Context, byUserKey swagger.UserKey) error {
	u, err := s.dbClient.User.Query().Where(user.Key(byUserKey)).Only(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user is none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		return sendClientError(ctx, http.StatusNotFound, "user is none")
	}
	return ctx.JSON(http.StatusOK, helper.ConvertEntUserAttribute(u))
}

// 属性更新
// (PUT /users/{byUserKey}/attributes)
func (s *ServerImpl) PutUsersByUserKeyAttributes(ctx echo.Context, byUserKey swagger.UserKey) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := helper.ConvertUserAttribute(userAttribute, s.dbClient.User.Create()).OnConflict().UpdateNewValues().ID(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	entUser, err := s.dbClient.User.Get(ctx.Request().Context(), id)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, helper.ConvertEntUserAttribute(entUser))
}

// アクティビティ群取得
// (GET /users/{byUserKey}/activities)
func (s *ServerImpl) GetUsersByUserKeyActivities(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// アクティビティ群最新化
// (PUT /users/{byUserKey}/activities)
func (s *ServerImpl) PutUsersByUserKeyActivities(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 資格情報群取得
// (GET /users/{byUserKey}/qualifications)
func (s *ServerImpl) GetUsersByUserKeyQualifications(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 資格情報群最新化
// (PUT /users/{byUserKey}/qualifications)
func (s *ServerImpl) PutUsersByUserKeyQualifications(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}
