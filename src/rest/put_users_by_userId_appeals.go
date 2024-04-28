package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PutUsersByUserIdAppeals アピール事項群最新化
// (PUT /users/{byUserId}/appeals)
func (s *strictServerImpl) PutUsersByUserIdAppeals(ctx context.Context, request PutUsersByUserIdAppealsRequestObject) (PutUsersByUserIdAppealsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithAppeals()
	})
	if err != nil {
		return nil, err
	}

	var results []*ent.UserAppeal
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		for _, appeal := range entUser.Edges.Appeals {
			if err := tx.UserAppeal.DeleteOne(appeal).Exec(ctx); err != nil {
				return err
			}
		}
		for _, appeal := range *request.Body {
			result, err := ToEntUserAppealCreate(appeal, entUser.ID, tx.UserAppeal.Create()).Save(ctx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PutUsersByUserIdAppeals200JSONResponse(ToSwaggerUserAppeals(results)), nil
}
