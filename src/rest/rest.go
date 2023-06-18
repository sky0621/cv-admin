package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/sky0621/cv-admin/src/ent"
)

type ServerImpl struct {
	dbClient *ent.Client
}

func NewRESTService(dbClient *ent.Client) ServerInterface {
	return NewStrictHandler(&strictServerImpl{dbClient}, nil)
}

func sendClientError(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, message)
}

func sendStatusInternalServerError(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusInternalServerError, message)
}
