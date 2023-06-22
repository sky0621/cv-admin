package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PutUsersByUserIdQualifications 資格情報群最新化
// (PUT /users/{byUserId}/qualifications)
func (s *strictServerImpl) PutUsersByUserIdQualifications(ctx context.Context, request PutUsersByUserIdQualificationsRequestObject) (PutUsersByUserIdQualificationsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithQualifications()
	})
	if err != nil {
		return nil, err
	}

	var userQualifications []UserQualification

	var results []*ent.UserQualification
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		for _, qualification := range entUser.Edges.Qualifications {
			if err := tx.UserQualification.DeleteOne(qualification).Exec(ctx); err != nil {
				return err
			}
		}
		for _, qualification := range userQualifications {
			result, err := ToEntUserQualificationCreate(qualification, entUser.ID, tx.UserQualification.Create()).Save(ctx)
			if err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PutUsersByUserIdQualifications200JSONResponse(ToSwaggerUserQualifications(results)), nil
}
