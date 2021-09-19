package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type QuestionHandler struct {
	service *services.QuestionService
}

func NewQuestionHandler(s *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: s}
}

func (h *QuestionHandler) FindCASMQuestions(c echo.Context) error {
	numOfItems, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("No se especifico el numero de items"))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("Indique el número de pagina"))
	}
	testID, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse("El Test ID es incorrecto"))
	}
	questions, err := h.service.GetCASMQuestion(testID, numOfItems, page)
	if err != nil {
		if err == errs.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		if err == errs.ErrNothingFind {
			return c.JSON(http.StatusNotFound, utils.NewNoFindResponse("No se encontro ningun registro"))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(questions))
}
