package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/api/router/paths"
	"github.com/user0608/expertos/injectors"
)

func logginUpgrader(e *echo.Echo) {
	h := injectors.GetLogginHandler()
	e.POST(paths.LOGIN, h.LogginEstudiante)
	e.POST(paths.LOGIN_ADMIN, h.LogginAdmin)
}
