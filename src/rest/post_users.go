package rest

import (
	"context"
	"errors"
)

// PostUsers ユーザー登録
// ユーザーアカウント登録
// (POST /users)
func (s *strictServerImpl) PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userAttribute := *request.Body

	if err := userAttribute.Validate(); err != nil {
		return PostUsers400JSONResponse{n400("validation failed")}, err
	}

	entUser, err := ToEntUserCreate(userAttribute, s.dbClient.User.Create()).Save(ctx)
	if err != nil {
		return nil, err
	}

	return PostUsers201JSONResponse(ToSwaggerUserAttribute(entUser)), nil
}
