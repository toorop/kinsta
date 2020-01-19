package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// recover
	e.Use(middleware.Recover())

	// routes

	// go go go !!
	e.Logger.Fatal(e.Start(":1323"))
}
