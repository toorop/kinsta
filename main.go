package main

import (
	"fmt"
	"kinsta/handlers"
	"kinsta/services/config"
	"kinsta/services/log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
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

	// log
	log.InitLogger(os.Stdout)

	// echo
	e := echo.New()

	// recover
	e.Use(middleware.Recover())

	// httpauth
	e.Use(middleware.BasicAuth(func(user, password string, c echo.Context) (bool, error) {
		if user == viper.GetString("user") && password == viper.GetString("password") {
			return true, nil
		}
		log.Infof("%s - bad password |%s| for user |%s|", c.RealIP(), password, user)
		return false, nil
	}))

	// routes

	// GET usenname
	e.GET("user/:username", handlers.GetUser)

	// go go go !!
	e.Logger.Fatal(e.Start(":1323"))
}
