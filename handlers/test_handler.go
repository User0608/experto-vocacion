package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/auth/roles"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type TestHandler struct {
	service *services.TestService
}

func NewTestHandler(s *services.TestService) *TestHandler {
	return &TestHandler{service: s}
}
func (h *TestHandler) LoadMessage(c echo.Context, err error) error {
	if err == errs.ErrDataBaseError {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
}
func (h *TestHandler) Create(c echo.Context) error {
	username, ok := c.Get(auth.USERNAME_KEY).(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	test, err := h.service.CreateTest(username)
	if err != nil {
		return h.LoadMessage(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(test))
}

func (h *TestHandler) FindAll(c echo.Context) error {
	role, ok := c.Get(auth.ROLE_KEY).(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if role == roles.ADMIN {
		tests, err := h.service.FindAll()
		if err != nil {
			return h.LoadMessage(c, err)
		}
		return c.JSON(http.StatusOK, utils.NewOkResponse(tests))
	} else if role == roles.USUARIO {
		username, ok := c.Get(auth.USERNAME_KEY).(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		tests, err := h.service.FindAllByEstudianteDNI(username)
		if err != nil {
			return h.LoadMessage(c, err)
		}
		return c.JSON(http.StatusOK, utils.NewOkResponse(tests))
	} else {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("No se puede realizar la operacion"))
	}
}

func (h *TestHandler) FindByID(c echo.Context) error {
	testID, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("parametro invalido!"))
	}
	test, err := h.service.FindID(testID)
	if err != nil {
		return h.LoadMessage(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(test))
}

func (h *TestHandler) Delete(c echo.Context) error {
	testID, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse("parametro invalido!"))
	}
	if err := h.service.DeleteTest(testID); err != nil {
		return h.LoadMessage(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkMesage("Eliminado correctamente"))
}
