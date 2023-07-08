package rest

import "context"

func (s *strictServerImpl) GetUsers(ctx context.Context, _ GetUsersRequestObject) (GetUsersResponseObject, error) {
	users, err := s.dbClient.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return GetUsers200JSONResponse(ToSwaggerUserAttributes(users)), nil
}
