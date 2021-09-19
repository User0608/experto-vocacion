package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/api/router/paths"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/auth/roles"
	"github.com/user0608/expertos/injectors"
)

func testUpgrader(e *echo.Echo) {
	h := injectors.GetTestHandler()
	g := e.Group(paths.PRUEBA)
	g.Use(auth.JWTMiddleware)
	g.GET("/create", auth.RolesMiddleware(h.Create, roles.USUARIO))
	g.GET("", auth.RolesMiddleware(h.FindAll, roles.ADMIN, roles.USUARIO))
	g.GET("/:test_id", h.FindByID)
	g.DELETE("/:test_id", h.Delete)
}
