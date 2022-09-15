package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/adapter/controller/rest"
	"github.com/sky0621/cv-admin/src/adapter/controller/swagger"
)

func main() {
	si := &rest.ServerImpl{}
	e := echo.New()
	swagger.RegisterHandlers(e, si)
	http.ListenAndServe(":3000", e)
}
