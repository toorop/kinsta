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
	return nil
	// test insta
	//Client = goinsta.New("peerpx", "w3R2FaM55522")
	Client = goinsta.New(viper.GetString("instaUser"), viper.GetString("instaPassword"))

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
		//return err
	}
	return nil
}
