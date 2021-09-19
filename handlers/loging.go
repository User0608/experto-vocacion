package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/auth/roles"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type LogginHandler struct {
	binder  echo.DefaultBinder
	service *services.LoginService
}

func NewLogginHandler(s *services.LoginService) *LogginHandler {
	return &LogginHandler{
		service: s,
		binder:  echo.DefaultBinder{},
	}
}

func (h *LogginHandler) LogginEstudiante(c echo.Context) error {
	data := &models.Session{}
	if err := h.binder.BindBody(c, data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(""))
	}
	estudiante, err := h.service.LogginEstudiante(data)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidData) || errors.Is(err, errs.ErrUsernameOrPassword) {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}
	usuario := models.Usuario{
		Username: estudiante.Dni,
		Role:     roles.USUARIO,
	}
	token, err := auth.GenerageTokenAdmin(usuario)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	response := utils.LogginResponse{
		Code:    "OK",
		Usuario: estudiante,
		Token:   token,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *LogginHandler) LogginAdmin(c echo.Context) error {
	data := &models.Session{}
	if err := h.binder.BindBody(c, data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(""))
	}
	if err := h.service.LogginAdmin(&models.Session{
		Username: data.Username,
		Password: data.Password,
	}); err != nil {
		if errors.Is(err, errs.ErrInvalidData) || errors.Is(err, errs.ErrUsernameOrPassword) {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(err.Error()))
	}

	token, err := auth.GenerageTokenAdmin(models.Usuario{Username: data.Username, Role: roles.ADMIN})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	response := utils.LogginResponse{
		Code:  "OK",
		Token: token,
	}
	return c.JSON(http.StatusOK, response)
}
