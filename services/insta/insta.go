package insta

import (
	"os"

	"github.com/tcnksm/go-input"
	"github.com/toorop/goinsta/v2"
)

var Client *goinsta.Instagram

func InitInsta() error {
	// test insta
	Client = goinsta.New("peerpx", "w3R2FaM55")

	if err := Client.Login(); err != nil {
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
		}
		return err
	}
	return nil
}
