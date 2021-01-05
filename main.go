package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/toorop/kinsta/handlers"
	"github.com/toorop/kinsta/services/config"
	"github.com/toorop/kinsta/services/insta"
	"github.com/toorop/kinsta/services/log"
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

	// ping
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// go go go !!
	var port int64
	port = 1323
	// port in env (docker)
	portStr := os.Getenv("PORT")
	if portStr != "" {
		port, err = strconv.ParseInt(portStr, 10, 64)
		if err != nil {
			log.Errorf("bad port defined in ENV: %s", portStr)
			os.Exit(1)
		}
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
