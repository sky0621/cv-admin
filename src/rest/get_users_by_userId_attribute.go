package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetUsersByUserIdAttribute 属性取得
// 属性取得
// (GET /users/{byUserId}/attribute)
func (s *strictServerImpl) GetUsersByUserIdAttribute(ctx context.Context, request GetUsersByUserIdAttributeRequestObject) (GetUsersByUserIdAttributeResponseObject, error) {
	entUser, err := s.getUserByUserId(ctx, request.ByUserId)
	if err != nil {
		if ent.IsNotFound(err) {
			return GetUsersByUserIdAttribute404JSONResponse{n404("user is none")}, nil
		}
		return nil, err
	}
	if entUser == nil {
		return GetUsersByUserIdAttribute404JSONResponse{n404("user is none")}, nil
	}

	return GetUsersByUserIdAttribute200JSONResponse(ToSwaggerUserAttribute(entUser)), nil
}
