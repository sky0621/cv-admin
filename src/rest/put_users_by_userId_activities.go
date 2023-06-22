package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PutUsersByUserIdActivities アクティビティ群最新化
// アクティビティ群最新化
// (PUT /users/{byUserId}/activities)
func (s *strictServerImpl) PutUsersByUserIdActivities(ctx context.Context, request PutUsersByUserIdActivitiesRequestObject) (PutUsersByUserIdActivitiesResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithActivities()
	})
	if err != nil {
		return nil, err
	}

	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userActivities := *request.Body

	for _, userActivity := range userActivities {
		if err := userActivity.Validate(); err != nil {
			return PutUsersByUserIdActivities400JSONResponse{n400("validation failed")}, err
		}
	}

	var results []*ent.UserActivity
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		for _, activity := range entUser.Edges.Activities {
			if err := tx.UserActivity.DeleteOne(activity).Exec(ctx); err != nil {
				return err
			}
		}
		for _, activity := range userActivities {
			result, err := ToEntUserActivityCreate(activity, entUser.ID, tx.UserActivity.Create()).Save(ctx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PutUsersByUserIdActivities200JSONResponse(ToSwaggerUserActivities(results)), nil
}
