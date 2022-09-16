package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

type ServerImpl struct {
	dbClient *ent.Client
}

func NewRESTService(dbClient *ent.Client) swagger.ServerInterface {
	return &ServerImpl{dbClient}
}
