package rest

import (
	"net/http"

	"github.com/sky0621/cv-admin/src/swagger"

	"github.com/labstack/echo/v4"
)

// ユーザー新規登録
// (POST /users)
func (s *ServerImpl) PostUsers(ctx echo.Context) error {
	var userAttribute swagger.UserAttribute
	if err := ctx.Bind(&userAttribute); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, "can not bind request")
	}

	// TODO: validation

	_, err := s.dbClient.User.Create().
		SetKey(*userAttribute.Key).
		SetName(*userAttribute.Name).
		SetNillableNickname(userAttribute.Nickname).
		SetNillableAvatarURL(userAttribute.AvatarUrl).
		SetBirthdayYear(int(*userAttribute.Birthday.Year)).
		SetBirthdayMonth(int(*userAttribute.Birthday.Month)).
		SetBirthdayDay(int(*userAttribute.Birthday.Day)).
		SetNillableJob(userAttribute.Job).
		SetNillableBelongTo(userAttribute.BelongTo).
		SetNillablePr(userAttribute.Pr).
		Save(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusCreated, "Created")
}

// 指定ユーザー削除
// (DELETE /users/{byUserKey})
func (s *ServerImpl) DeleteUsersByUserKey(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
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

// 属性取得
// (GET /users/{byUserKey}/attributes)
func (s *ServerImpl) GetUsersByUserKeyAttributes(ctx echo.Context, byUserKey swagger.UserKey) error {
	return ctx.String(http.StatusOK, byUserKey)
}

// 属性更新
// (PUT /users/{byUserKey}/attributes)
func (s *ServerImpl) PutUsersByUserKeyAttributes(ctx echo.Context, byUserKey swagger.UserKey) error {
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
