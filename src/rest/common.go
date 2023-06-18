package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent/user"

	"github.com/sky0621/cv-admin/src/ent"
)

func (s *strictServerImpl) getUserByUserIdWithXXX(ctx context.Context, byUserId UserId, fn func(q *ent.UserQuery) *ent.UserQuery) (*ent.User, error) {
	entUser, err := fn(s.dbClient.User.Query().Where(user.ID(byUserId))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return entUser, nil
}

func (s *strictServerImpl) getUserByUserId(ctx context.Context, byUserId UserId) (*ent.User, error) {
	entUser, err := s.dbClient.User.Get(ctx, byUserId)
	if err != nil {
		return nil, err
	}
	return entUser, nil
}
