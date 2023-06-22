package rest

import (
	"context"
	"errors"
)

// GetUsersByUserIdAttribute 属性取得
// 属性取得
// (GET /users/{byUserId}/attribute)
func (s *strictServerImpl) GetUsersByUserIdAttribute(ctx context.Context, request GetUsersByUserIdAttributeRequestObject) (GetUsersByUserIdAttributeResponseObject, error) {
	entUser, err := s.getUserByUserId(ctx, request.ByUserId)
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return GetUsersByUserIdAttribute404JSONResponse{n404("user is none")}, err
		}
		return nil, err
	}
	if entUser == nil {
		return GetUsersByUserIdAttribute404JSONResponse{n404("user is none")}, nil
	}

	return GetUsersByUserIdAttribute200JSONResponse(ToSwaggerUserAttribute(entUser)), nil
}
