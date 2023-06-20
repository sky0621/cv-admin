package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetUsersByUserIdQualifications 資格情報群取得
// (GET /users/{byUserId}/qualifications)
func (s *strictServerImpl) GetUsersByUserIdQualifications(ctx context.Context, request GetUsersByUserIdQualificationsRequestObject) (GetUsersByUserIdQualificationsResponseObject, error) {
	entUser, err := s.getUserByUserIdWithXXX(ctx, request.ByUserId, func(q *ent.UserQuery) *ent.UserQuery {
		return q.WithQualifications()
	})
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdQualifications200JSONResponse{ToSwaggerUserQualifications(entUser.Edges.Qualifications)}, nil
}
