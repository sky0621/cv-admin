package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetUsersByUserIdSolutions 課題解決事例群取得
// (GET /users/{byUserId}/solutions)
func (s *strictServerImpl) GetUsersByUserIdSolutions(ctx context.Context, request GetUsersByUserIdSolutionsRequestObject) (GetUsersByUserIdSolutionsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithSolutions()
	})
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdSolutions200JSONResponse(ToSwaggerUserSolutions(entUser.Edges.Solutions)), nil
}
