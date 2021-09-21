package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/services"
	"github.com/user0608/expertos/utils"
)

type ResponseHandler struct {
	service *services.ResponseService
}

func NewResponseHandler(s *services.ResponseService) *ResponseHandler {
	return &ResponseHandler{service: s}
}

func (h *ResponseHandler) Respuesta(c echo.Context) error {
	testID, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewBadResponse(""))
	}
	test, err := h.service.Response(testID)
	if err != nil {
		if err == errs.ErrDataBaseError {
			return c.JSON(http.StatusInternalServerError, utils.NewInternalErrorResponse(""))
		}
		return c.JSON(http.StatusBadRequest, utils.NewBadResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewOkResponse(test))
}
