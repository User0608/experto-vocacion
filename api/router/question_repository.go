package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/api/router/paths"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/auth/roles"
	"github.com/user0608/expertos/injectors"
)

func questionUpgrader(e *echo.Echo) {
	h := injectors.GetQuestionHandler()
	g := e.Group(paths.TEST)
	g.Use(auth.JWTMiddleware)
	g.GET("/:test_id/casm", auth.RolesMiddleware(h.FindCASMQuestions, roles.ADMIN, roles.USUARIO))
}
