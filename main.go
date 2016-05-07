package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/VojtechVitek/go-trello"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	AppKey   string `toml:"app_key"`
	Token    string `toml:"token"`
	Username string `toml:"username"`
}

func setColor(c string) {
	if c == "red" {
		color.Set(color.FgRed)
	} else if c == "yellow" {
		color.Set(color.FgYellow)
	} else if c == "green" {
		color.Set(color.FgGreen)
	} else if c == "purple" {
		color.Set(color.FgMagenta)
	}
}

func main() {
	appKey := os.Getenv("TRELLO_APP_KEY")
	token := os.Getenv("TRELLO_TOKEN")
	username := os.Getenv("TRELLO_USER")

	if appKey == "" || token == "" || username == "" {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		var config Config
		_, err = toml.DecodeFile(home+"/.trelloclirc", &config)
		if err != nil {
			log.Fatal(err)
		}

		appKey = config.AppKey
		token = config.Token
		username = config.Username
	}

	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	user, err := trello.Member(username)
	if err != nil {
		log.Fatal(err)
	}

	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			cards, _ := list.Cards()
			if len(cards) > 0 {
				fmt.Println("  *", list.Name)

				for _, card := range cards {
					for _, label := range card.Labels {
						setColor(label.Color)
						break
					}

					if card.Due != "" {
						t, _ := time.Parse(time.RFC3339Nano, card.Due)
						fmt.Printf("    * %s %s\n", t.Format("[2006-01-02]"), card.Name)
					} else {
						fmt.Printf("    * %s\n", card.Name)
					}

					color.Unset()
				}
			}
		}
	}
}
