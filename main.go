package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const expertApi = "http://127.0.0.1:81/experto"

type Request struct {
	Area     string `json:"area"`
	Habito   string `json:"habito"`
	Caracter string `json:"caracter"`
}

type ExpertResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Respuesta string `json:"respuesta"`
}

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.POST("/experto", handlerPost)
	e.Logger.Fatal(e.Start("localhost:80"))
}

func handlerPost(c echo.Context) error {
	binder := echo.DefaultBinder{}
	data := &Request{}
	if err := binder.BindBody(c, data); err != nil {
		log.Fatalln("Datos incorrecta", err.Error())
	}
	datoByteFormat, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error reMarshal", err.Error())
		return echo.ErrInternalServerError
	}
	req, err := http.NewRequest(http.MethodPost, expertApi, bytes.NewReader(datoByteFormat))
	if err != nil {
		log.Fatalln("Error new request", err.Error())
		return echo.ErrInternalServerError
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Error comunicacion con servicio experto", err.Error())
		return echo.ErrInternalServerError
	}
	rr := &ExpertResponse{}
	if err := json.NewDecoder(res.Body).Decode(rr); err != nil {
		log.Fatalln("Respuesta incorrecta, servicio experto", err.Error())
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, rr)
}
