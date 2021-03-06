package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"src/config"
	"src/model"
	"src/router"
)

func main() {
	model.Connect()
	defer model.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.CreateRouters(e)

	e.Logger.Fatal(e.Start(config.C.App.Addr))
}
