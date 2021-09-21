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
	hr := injectors.GetResponseHandler()
	g := e.Group(paths.TEST)
	g.Use(auth.JWTMiddleware)
	g.GET("/:test_id/casm", auth.RolesMiddleware(h.FindCASMQuestions, roles.ADMIN, roles.USUARIO))
	g.GET("/:test_id/berger", auth.RolesMiddleware(h.FindBergerQuestions, roles.ADMIN, roles.USUARIO))
	g.GET("/:test_id/hea", auth.RolesMiddleware(h.FindHEAQuestions, roles.ADMIN, roles.USUARIO))
	g.GET("/:test_id/resultado", auth.RolesMiddleware(hr.Respuesta, roles.ADMIN, roles.USUARIO))
	//create
	g.POST("/:test_id/casm", auth.RolesMiddleware(h.RegisterCASMQuestionAnswer, roles.USUARIO, roles.USUARIO))
	g.POST("/:test_id/berger", auth.RolesMiddleware(h.RegisterBergerQuestionAnswer, roles.USUARIO, roles.USUARIO))
	g.POST("/:test_id/hea", auth.RolesMiddleware(h.RegisterHEAQuestionAnswer, roles.USUARIO, roles.USUARIO))

}
