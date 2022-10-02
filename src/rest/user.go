package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/cv-admin/src/ent/user"
)

// PostUsers ユーザー登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	var userAttribute UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := ToEntUserCreate(userAttribute, s.dbClient.User.Create()).
		Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerUserAttribute(entUser))
}

// DeleteUsersByUserId 指定ユーザー削除
// (DELETE /users/{byUserId})
func (s *ServerImpl) DeleteUsersByUserId(ctx echo.Context, byUserId UserId) error {
	err := s.dbClient.User.DeleteOneID(byUserId).Exec(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s *ServerImpl) getUserByUserId(ctx echo.Context, byUserId UserId) (*ent.User, error) {
	entUser, err := s.dbClient.User.Get(ctx.Request().Context(), byUserId)
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			sendClientError(ctx, http.StatusNotFound, "user is none")
			return nil, err
		}
		sendClientError(ctx, http.StatusInternalServerError, err.Error())
		return nil, err
	}

	if entUser == nil {
		sendClientError(ctx, http.StatusNotFound, "user is none")
		return nil, errors.New("user is none")
	}

	return entUser, nil
}

// GetUsersByUserIdAttribute 属性取得
// (GET /users/{byUserId}/attribute)
func (s *ServerImpl) GetUsersByUserIdAttribute(ctx echo.Context, byUserId UserId) error {
	entUser, err := s.getUserByUserId(ctx, byUserId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ToSwaggerUserAttribute(entUser))
}

// PutUsersByUserIdAttribute 属性更新
// (PUT /users/{byUserId}/attribute)
func (s *ServerImpl) PutUsersByUserIdAttribute(ctx echo.Context, byUserId UserId) error {
	var userAttribute UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userAttribute.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entUser, err := ToEntUserUpdate(userAttribute, s.dbClient.User.UpdateOneID(byUserId)).Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserAttribute(entUser))
}

// GetUsersByUserIdActivities アクティビティ群取得
// (GET /users/{byUserId}/activities)
func (s *ServerImpl) GetUsersByUserIdActivities(ctx echo.Context, byUserId UserId) error {
	entUser, err := s.getUserByUserId(ctx, byUserId)
	if err != nil {
		return err
	}

	activities, err := entUser.QueryActivities().All(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user activities are none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserActivities(activities))
}

// PutUsersByUserIdActivities アクティビティ群最新化
// (PUT /users/{byUserId}/activities)
func (s *ServerImpl) PutUsersByUserIdActivities(ctx echo.Context, byUserId UserId) error {
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

	var userActivities UserActivities
	if err := ctx.Bind(&userActivities); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	for _, userActivity := range userActivities {
		if err := userActivity.Validate(); err != nil {
			return sendClientError(ctx, http.StatusBadRequest, err.Error())
		}
	}

	var results []*ent.UserActivity
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		for _, activity := range entUser.Edges.Activities {
			tx.UserActivity.DeleteOne(activity).Exec(rCtx)
		}
		for _, activity := range userActivities {
			result, err := ToEntUserActivityCreate(activity, entUser.ID, tx.UserActivity.Create()).Save(rCtx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserActivities(results))
}

// GetUsersByUserIdQualifications 資格情報群取得
// (GET /users/{byUserId}/qualifications)
func (s *ServerImpl) GetUsersByUserIdQualifications(ctx echo.Context, byUserId UserId) error {
	entUser, err := s.getUserByUserId(ctx, byUserId)
	if err != nil {
		return err
	}

	qualifications, err := entUser.QueryQualifications().All(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return sendClientError(ctx, http.StatusNotFound, "user qualifications are none")
		}
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserQualifications(qualifications))
}

// PutUsersByUserIdQualifications 資格情報群最新化
// (PUT /users/{byUserId}/qualifications)
func (s *ServerImpl) PutUsersByUserIdQualifications(ctx echo.Context, byUserId UserId) error {
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

	var userQualifications []UserQualification
	if err := ctx.Bind(&userQualifications); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	var results []*ent.UserQualification
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		for _, qualification := range entUser.Edges.Qualifications {
			tx.UserQualification.DeleteOne(qualification).Exec(rCtx)
		}
		for _, qualification := range userQualifications {
			result, err := ToEntUserQualificationCreate(qualification, entUser.ID, tx.UserQualification.Create()).Save(rCtx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserQualifications(results))
}
