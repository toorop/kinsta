package main

import (
	"fmt"
	"kinsta/services/config"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// home
	home, err := os.Executable()
	if err != nil {
		fmt.Printf("os.Executable() failed - %v", err)
		os.Exit(1)
	}

	// config
	if err = config.Init(home); err != nil {
		fmt.Printf("config.Init(%s) failed - %v", home, err)
	}

	e := echo.New()

	// recover
	e.Use(middleware.Recover())

	// routes

	// go go go !!
	e.Logger.Fatal(e.Start(":1323"))
}
