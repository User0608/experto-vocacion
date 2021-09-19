package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/api/router/paths"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/auth/roles"
	"github.com/user0608/expertos/injectors"
)

func estudianteUpgrader(e *echo.Echo) {
	h := injectors.GetEstudianteHandler()
	e.POST(fmt.Sprintf("%s/create", paths.ESTUDIANTE), h.Create)
	group := e.Group(paths.ESTUDIANTE)
	group.Use(auth.JWTMiddleware)
	group.PUT("", auth.RolesMiddleware(h.Update, roles.USUARIO))
	group.GET("", auth.RolesMiddleware(h.GetAll, roles.ADMIN))
	group.GET("/me", auth.RolesMiddleware(h.GetMe, roles.USUARIO))
	group.GET("/:id", auth.RolesMiddleware(h.GetByID, roles.ADMIN))
}
