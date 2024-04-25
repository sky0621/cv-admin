package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PutUsersByUserIdSolutions 課題解決事例群最新化
// (PUT /users/{byUserId}/solutions)
func (s *strictServerImpl) PutUsersByUserIdSolutions(ctx context.Context, request PutUsersByUserIdSolutionsRequestObject) (PutUsersByUserIdSolutionsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithSolutions()
	})
	if err != nil {
		return nil, err
	}

	var results []*ent.UserSolution
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		for _, solution := range entUser.Edges.Solutions {
			if err := tx.UserSolution.DeleteOne(solution).Exec(ctx); err != nil {
				return err
			}
		}
		for _, solution := range *request.Body {
			result, err := ToEntUserSolutionCreate(solution, entUser.ID, tx.UserSolution.Create()).Save(ctx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PutUsersByUserIdSolutions200JSONResponse(ToSwaggerUserSolutions(results)), nil
}
