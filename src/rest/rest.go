package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent"
)

type ServerImpl struct {
	dbClient *ent.Client
}

func NewRESTService(dbClient *ent.Client) ServerInterface {
	return &ServerImpl{dbClient}
}

func sendClientError(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, &ClientError{Message: &message})
}
