package rest

import (
	"context"
	"errors"
)

// PutUsersByUserIdAttribute 属性更新
// 属性更新
// (PUT /users/{byUserId}/attribute)
func (s *strictServerImpl) PutUsersByUserIdAttribute(ctx context.Context, request PutUsersByUserIdAttributeRequestObject) (PutUsersByUserIdAttributeResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userAttribute := *request.Body

	if err := userAttribute.Validate(); err != nil {
		return PutUsersByUserIdAttribute400JSONResponse{n400("validation failed")}, err
	}

	entUser, err := ToEntUserUpdate(userAttribute, s.dbClient.User.UpdateOneID(request.ByUserId)).Save(ctx)
	if err != nil {
		return nil, err
	}

	return PutUsersByUserIdAttribute200JSONResponse{N200OKUserAttributeJSONResponse(ToSwaggerUserAttribute(entUser))}, nil
}
