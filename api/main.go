package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/user0608/expertos/api/router"
	"github.com/user0608/expertos/auth"
	"github.com/user0608/expertos/configs"
	"github.com/user0608/expertos/errs"
)

func main() {
	conf, err := configs.LoadServiceConfigs("service_config.json")
	if err != nil {
		log.Fatalln("No se cargo las configuraciones de servicio:", err.Error())
	}
	log.Println("Configuraciones de servicio cargados!")
	if err := auth.LoadFiles(conf.Certificates.Private, conf.Certificates.Public); err != nil {
		log.Fatalln("No se cargaron los certificados,", err.Error())
	}
	log.Println("Certificados cargados!")

	e := echo.New()
	e.HTTPErrorHandler = errs.CustomErrorHandler
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.Cors.AllowOrigins,
		AllowMethods: conf.Cors.AllowMethods,
	}))
	e.Static("/", "public")
	router.Upgrade(e)
	log.Fatal(e.Start(conf.Address))
	// e.POST("/experto", handlerPost)
}
