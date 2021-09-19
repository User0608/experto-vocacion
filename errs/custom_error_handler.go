package errs

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/user0608/expertos/utils"
)

func CustomErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Issue #1426
	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		message = utils.Response{
			Code:    "ERROR",
			Message: m,
		}
	} else {
		log.Println("Paso algo!!!")
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			log.Println(err.Error())
		}
	}
}
