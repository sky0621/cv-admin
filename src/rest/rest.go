package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

type ServerImpl struct {
	dbClient *ent.Client
}

func NewRESTService(dbClient *ent.Client) swagger.ServerInterface {
	return &ServerImpl{dbClient}
}

func sendClientError(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, &swagger.ClientError{Message: &message})
}
