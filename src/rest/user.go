package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/rest/converter"
	"github.com/sky0621/cv-admin/src/swagger"
)

// PostUsers ユーザー登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := converter.ToEntUserCreate(userAttribute, s.dbClient.User.Create()).
		Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, converter.ToSwaggerUserAttribute(entUser))
}

// DeleteUsersByUserId 指定ユーザー削除
// (DELETE /users/{byUserId})
func (s *ServerImpl) DeleteUsersByUserId(ctx echo.Context, byUserId swagger.UserId) error {
	err := s.dbClient.User.DeleteOneID(byUserId).Exec(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetUsersByUserIdAttribute 属性取得
// (GET /users/{byUserId}/attribute)
func (s *ServerImpl) GetUsersByUserIdAttribute(ctx echo.Context, byUserId swagger.UserId) error {
	entUser, err := s.dbClient.User.Get(ctx.Request().Context(), byUserId)
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

	return ctx.JSON(http.StatusOK, converter.ToSwaggerUserAttribute(entUser))
}

// PutUsersByUserIdAttribute 属性更新
// (PUT /users/{byUserId}/attribute)
func (s *ServerImpl) PutUsersByUserIdAttribute(ctx echo.Context, byUserId swagger.UserId) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := converter.ToEntUserUpdate(userAttribute, s.dbClient.User.UpdateOneID(byUserId)).Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, converter.ToSwaggerUserAttribute(entUser))
}

// GetUsersByUserIdActivities アクティビティ群取得
// (GET /users/{byUserId}/activities)
func (s *ServerImpl) GetUsersByUserIdActivities(ctx echo.Context, byUserId swagger.UserId) error {
	rCtx := ctx.Request().Context()

	entUser, err := s.dbClient.User.Get(rCtx, byUserId)
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

	activities, err := entUser.QueryActivities().All(rCtx)
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user activities are none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, converter.ToSwaggerUserActivities(activities))
}

// PutUsersByUserIdActivities アクティビティ群最新化
// (PUT /users/{byUserId}/activities)
func (s *ServerImpl) PutUsersByUserIdActivities(ctx echo.Context, byUserId swagger.UserId) error {
	rCtx := ctx.Request().Context()

	entUser, err := s.dbClient.User.Query().Where(user.ID(byUserId)).WithActivities().Only(rCtx)
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

	var userActivities []swagger.UserActivity
	if err := ctx.Bind(&userActivities); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	var results []*ent.UserActivity
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		for _, activity := range entUser.Edges.Activities {
			tx.UserActivity.DeleteOne(activity).Exec(rCtx)
		}
		for _, activity := range userActivities {
			result, err := converter.ToEntUserActivityCreate(activity, entUser.ID, tx.UserActivity.Create()).Save(rCtx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, converter.ToSwaggerUserActivities(results))
}

// 資格情報群取得
// (GET /users/{byUserId}/qualifications)
func (s *ServerImpl) GetUsersByUserIdQualifications(ctx echo.Context, byUserId swagger.UserId) error {
	return ctx.String(http.StatusOK, "")
}

// PutUsersByUserIdQualifications 資格情報群最新化
// (PUT /users/{byUserId}/qualifications)
func (s *ServerImpl) PutUsersByUserIdQualifications(ctx echo.Context, byUserId swagger.UserId) error {
	rCtx := ctx.Request().Context()

	entUser, err := s.dbClient.User.Query().Where(user.ID(byUserId)).WithQualifications().Only(rCtx)
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

	var userQualifications []swagger.UserQualification
	if err := ctx.Bind(&userQualifications); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	var results []*ent.UserQualification
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		for _, qualification := range entUser.Edges.Qualifications {
			tx.UserQualification.DeleteOne(qualification).Exec(rCtx)
		}
		for _, qualification := range userQualifications {
			result, err := converter.ToEntUserQualificationCreate(qualification, entUser.ID, tx.UserQualification.Create()).Save(rCtx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, converter.ToSwaggerUserActivities(results))
}
