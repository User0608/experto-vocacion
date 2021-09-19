package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/berger"
	"github.com/user0608/expertos/models/casm"
	"github.com/user0608/expertos/models/hea"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type QuestionHandler struct {
	service *services.QuestionService
	binder  echo.DefaultBinder
}

func NewQuestionHandler(s *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: s, binder: echo.DefaultBinder{}}
}
func (h *QuestionHandler) GetParams(c echo.Context) (int, int, int, error) {
	numOfItems, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return 0, 0, 0, errors.New("No se especifico el numero de items")
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return 0, 0, 0, errors.New("Indique el número de pagina")
	}
	testID, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return 0, 0, 0, errors.New("El Test ID es incorrecto")
	}
	return testID, numOfItems, page, nil
}
func (h *QuestionHandler) PrepreErrorResponse(c echo.Context, err error) error {
	if err == errs.ErrDataBaseError {
		return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
	}
	if err == errs.ErrNothingFind {
		return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("No se encontro ningun registro"))
	}
	return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
}
func (h *QuestionHandler) FindCASMQuestions(c echo.Context) error {
	testID, numOfItems, page, err := h.GetParams(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	questions, err := h.service.GetCASMQuestion(testID, numOfItems, page)
	if err != nil {
		return h.PrepreErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(questions))
}
func (h *QuestionHandler) RegisterCASMQuestionAnswer(c echo.Context) error {
	answers := []casm.TestCasm{}
	if err := h.binder.BindBody(c, &answers); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida!"))
	}
	if created, err := h.service.RegisterCASMQuestionAnswer(answers); err != nil {
		return h.PrepreErrorResponse(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(created))
	}
}

func (h *QuestionHandler) FindBergerQuestions(c echo.Context) error {
	testID, numOfItems, page, err := h.GetParams(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	questions, err := h.service.GetBergerQuestions(testID, numOfItems, page)
	if err != nil {
		return h.PrepreErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(questions))
}

func (h *QuestionHandler) RegisterBergerQuestionAnswer(c echo.Context) error {
	answers := []berger.TestBerger{}
	if err := h.binder.BindBody(c, &answers); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida!"))
	}
	if created, err := h.service.RegisterBergerQuestionAnswer(answers); err != nil {
		return h.PrepreErrorResponse(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(created))
	}
}

func (h *QuestionHandler) FindHEAQuestions(c echo.Context) error {
	testID, numOfItems, page, err := h.GetParams(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	questions, err := h.service.GetHEAQuestions(testID, numOfItems, page)
	if err != nil {
		return h.PrepreErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(questions))
}
func (h *QuestionHandler) RegisterHEAQuestionAnswer(c echo.Context) error {
	answers := []hea.TestHea{}
	if err := h.binder.BindBody(c, &answers); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Estructura JSON Invalida!"))
	}
	if created, err := h.service.RegisterHEAQuestionAnswer(answers); err != nil {
		return h.PrepreErrorResponse(c, err)
	} else {
		return c.JSON(http.StatusOK, utils.NewOkResponse(created))
	}
}
