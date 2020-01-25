package main

import (
	"fmt"
	"kinsta/services/config"
	"kinsta/services/log"
	"os"
	"path/filepath"

	"github.com/tcnksm/go-input"

	"github.com/toorop/goinsta/v2"
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
	}

	// log
	log.InitLogger(os.Stdout)

	// test insta
	insta := goinsta.New("peerpx", "w3R2FaM5")

	if err := insta.Login(); err != nil {
		switch v := err.(type) {
		case goinsta.ChallengeError:
			err := insta.Challenge.Process(v.Challenge.APIPath)
			if err != nil {
				panic(err)
			}

			ui := &input.UI{
				Writer: os.Stdout,
				Reader: os.Stdin,
			}

			query := "What is SMS code for instagram?"
			code, err := ui.Ask(query, &input.Options{
				Default:  "000000",
				Required: true,
				Loop:     true,
			})
			if err != nil {
				panic(err)
			}

			err = insta.Challenge.SendSecurityCode(code)
			if err != nil {
				panic(err)
			}
		}

		panic(err)
	}
	defer func() { _ = insta.Logout() }()

	user, err := insta.Profiles.ByName("wesbos")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Printf("followers: %d\n", user.FollowerCount)

	/*err = user.Sync()
	if err != nil {
		panic(err)
	}*/
	feed := user.Feed()
	feed.Next()

	for _, post := range feed.Items {
		fmt.Printf("Images: %v\n", post.Images)
	}

	fmt.Printf("feed: %v", feed.Items[0].Images)

	/*
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
		e.Logger.Fatal(e.Start(":1323"))*/
}
