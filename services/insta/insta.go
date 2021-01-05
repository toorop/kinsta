package insta

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"github.com/toorop/goinsta"
)

var Client *goinsta.Instagram

func InitInsta() error {
	instaUser := os.Getenv("INSTA_LOGIN")
	if instaUser == "" {
		instaUser = viper.GetString("instaUser")
	}
	instaPassword := os.Getenv("INSTA_PASSWORD")
	if instaPassword == "" {
		instaPassword = viper.GetString("instaPassword")
	}

	//log.Infof("user: %s, password: %s\n", instaUser, instaPassword)

	Client = goinsta.New(instaUser, instaPassword)
	if err := Client.Login(); err != nil {
		fmt.Printf("Error login: %v", err)
		switch v := err.(type) {
		case goinsta.ChallengeError:
			err := Client.Challenge.Process(v.Challenge.APIPath)
			if err != nil {
				return err
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
				return err
			}

			err = Client.Challenge.SendSecurityCode(code)
			if err != nil {
				return err
			}
			Client.Account = Client.Challenge.LoggedInUser
		}
	}
	return nil
}
