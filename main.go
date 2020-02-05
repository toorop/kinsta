package main

import (
	"fmt"
	"kinsta/handlers"
	"kinsta/services/config"
	"kinsta/services/insta"
	"kinsta/services/log"
	"os"
	"path/filepath"

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
	home = filepath.Dir(home)

	// config
	if err = config.Init(home); err != nil {
		fmt.Printf("config.Init(%s) failed - %v", home, err)
		os.Exit(1)
	}

	// log
	log.InitLogger(os.Stdout)

	// init insta client
	if err = insta.InitInsta(); err != nil {
		log.Errorf("insta.InitInsta() failed: %v", err)
		os.Exit(1)
	}

	defer func() { _ = insta.Client.Logout() }()

	// echo
	e := echo.New()

	// recover
	e.Use(middleware.Recover())

	// httpauth
	/*e.Use(middleware.BasicAuth(func(user, password string, c echo.Context) (bool, error) {
		if user == viper.GetString("user") && password == viper.GetString("password") {
			return true, nil
		}
		log.Infof("%s - bad password |%s| for user |%s|", c.RealIP(), password, user)
		return false, nil
	}))*/

	// routes

	// GET usenname
	e.GET("/", handlers.GetUser)

	// go go go !!
	e.Logger.Fatal(e.Start(":1323"))
}
