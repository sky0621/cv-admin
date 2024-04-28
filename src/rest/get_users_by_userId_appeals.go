package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetUsersByUserIdAppeals アピール事項群取得
// (GET /users/{byUserId}/appeals)
func (s *strictServerImpl) GetUsersByUserIdAppeals(ctx context.Context, request GetUsersByUserIdAppealsRequestObject) (GetUsersByUserIdAppealsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithAppeals()
	})
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdAppeals200JSONResponse(ToSwaggerUserAppeals(entUser.Edges.Appeals)), nil
}
