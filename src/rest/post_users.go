package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/converter"
)

// PostUsers ユーザー登録
// ユーザーアカウント登録
// (POST /users)
func (s *strictServerImpl) PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userAttribute := *request.Body

	postUsers400JSONResponse := func(msg string) PostUsers400JSONResponse {
		return PostUsers400JSONResponse{N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}}
	}

	if err := userAttribute.Validate(); err != nil {
		return postUsers400JSONResponse("validation failed"), err
	}

	entUser, err := ToEntUserCreate(userAttribute, s.dbClient.User.Create()).Save(ctx)
	if err != nil {
		return nil, err
	}

	return PostUsers201JSONResponse{N201CreatedUserAttributeJSONResponse(ToSwaggerUserAttribute(entUser))}, nil
}
