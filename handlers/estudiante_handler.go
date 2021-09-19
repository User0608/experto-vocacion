package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type EstudianteHandler struct {
	binder  echo.DefaultBinder
	service *services.EstudianteService
}

func NewEstudianteService(s *services.EstudianteService) *EstudianteHandler {
	return &EstudianteHandler{
		service: s,
		binder:  echo.DefaultBinder{},
	}
}
func (h *EstudianteHandler) createUpdate(c echo.Context, action string) error {
	estudiante := &models.Estudiante{}
	if err := h.binder.BindBody(c, estudiante); err != nil {
		if strings.Contains(err.Error(), "parsing time") {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("La fecha debe cumplir el estandar RFC3339, 2021-09-12T00:00:00Z"))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(""))
	}
	var err error
	if estudiante.ID != 0 {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("No puede utilizar el campo ID!"))
	}
	if action == "update" {
		dni, ok := c.Get(auth.USERNAME_KEY).(string)
		if !ok {
			return c.JSON(http.StatusForbidden, utils.NewForbiddenResponse(""))
		}
		if estudiante.Dni != "" {
			return c.JSON(http.StatusBadRequest, utils.NewBadResponse("No puede utilizar el campo DNI!"))
		}
		estudiante.Dni = dni
		err = h.service.Update(estudiante)
	} else {
		err = h.service.Create(estudiante)
	}
	if err != nil {
		if err == errs.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(estudiante))
}
func (h *EstudianteHandler) Create(c echo.Context) error {
	return h.createUpdate(c, "create")
}

func (h *EstudianteHandler) Update(c echo.Context) error {
	return h.createUpdate(c, "update")
}

func (h *EstudianteHandler) GetAll(c echo.Context) error {
	estudiantes, err := h.service.Find()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(estudiantes))
}

func (h *EstudianteHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("Invalid ID"))
	}
	estudiante, err := h.service.FindByID(id)
	if err != nil {
		if err == errs.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(estudiante))
}

func (h *EstudianteHandler) GetMe(c echo.Context) error {
	dni, ok := c.Get(auth.USERNAME_KEY).(string)
	if !ok {
		return c.JSON(http.StatusForbidden, utils.NewForbiddenResponse(""))
	}
	estudiante, err := h.service.FindByDNI(dni)
	if err != nil {
		if err == errs.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(estudiante))
}
