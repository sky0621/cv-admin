package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetUsersByUserIdActivities アクティビティ群取得
// アクティビティ群取得
// (GET /users/{byUserId}/activities)
func (s *strictServerImpl) GetUsersByUserIdActivities(ctx context.Context, request GetUsersByUserIdActivitiesRequestObject) (GetUsersByUserIdActivitiesResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithActivities()
	})
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return GetUsersByUserIdActivities404JSONResponse{n404("user is none")}, err
		}
		return nil, err
	}
	if entUser == nil {
		return GetUsersByUserIdActivities404JSONResponse{n404("user is none")}, nil
	}

	return GetUsersByUserIdActivities200JSONResponse{ToSwaggerUserActivities(entUser.Edges.Activities)}, nil
}
