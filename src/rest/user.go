package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/rest/helper"
	"github.com/sky0621/cv-admin/src/swagger"
)

// ユーザー登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := helper.ConvertSwaggerUserAttributeToEntUserCreate(userAttribute, s.dbClient.User.Create()).
		Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, helper.ConvertEntUserToSwaggerUserAttribute(entUser))
}

// 指定ユーザー削除
// (DELETE /users/{byUserId})
func (s *ServerImpl) DeleteUsersByUserId(ctx echo.Context, byUserId swagger.UserId) error {
	cnt, err := s.dbClient.User.Delete().Where(user.ID(byUserId)).Exec(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if cnt != 1 {
		return sendClientError(ctx, http.StatusNotFound, "user is none")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// 属性取得
// (GET /users/{byUserId}/attribute)
func (s *ServerImpl) GetUsersByUserIdAttribute(ctx echo.Context, byUserId swagger.UserId) error {
	entUser, err := s.dbClient.User.Query().Where(user.ID(byUserId)).Only(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user is none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if entUser == nil {
		return sendClientError(ctx, http.StatusNotFound, "user is none")
	}

	return ctx.JSON(http.StatusOK, helper.ConvertEntUserToSwaggerUserAttribute(entUser))
}

// 属性更新
// (PUT /users/{byUserId}/attribute)
func (s *ServerImpl) PutUsersByUserIdAttribute(ctx echo.Context, byUserId swagger.UserId) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := helper.ConvertSwaggerUserAttributeToEntUserUpdate(userAttribute, s.dbClient.User.UpdateOneID(byUserId)).Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, helper.ConvertEntUserToSwaggerUserAttribute(entUser))
}

// アクティビティ群取得
// (GET /users/{byUserId}/activities)
func (s *ServerImpl) GetUsersByUserIdActivities(ctx echo.Context, byUserId swagger.UserId) error {
	entUser, err := s.dbClient.User.Query().Where(user.ID(byUserId)).Only(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user is none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if entUser == nil {
		return sendClientError(ctx, http.StatusNotFound, "user is none")
	}

	activities, err := entUser.QueryActivities().Only(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user activities are none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(activities)

	return ctx.String(http.StatusOK, "")
}

// アクティビティ群最新化
// (PUT /users/{byUserId}/activities)
func (s *ServerImpl) PutUsersByUserIdActivities(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}

// 資格情報群取得
// (GET /users/{byUserId}/qualifications)
func (s *ServerImpl) GetUsersByUserIdQualifications(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}

// 資格情報群最新化
// (PUT /users/{byUserId}/qualifications)
func (s *ServerImpl) PutUsersByUserIdQualifications(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}
