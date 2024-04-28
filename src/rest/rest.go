package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
)

type strictServerImpl struct {
	dbClient *ent.Client
}

func NewRESTService(dbClient *ent.Client) ServerInterface {
	return NewStrictHandler(&strictServerImpl{dbClient}, nil)
}

func n400(msg string) N400BadRequestJSONResponse {
	return N400BadRequestJSONResponse{Message: ToPtr(msg)}
}

func n404(msg string) N404NotFoundJSONResponse {
	return N404NotFoundJSONResponse{Message: ToPtr(msg)}
}
