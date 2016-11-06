package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VojtechVitek/go-trello"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
)

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

func showBoards(boards []trello.Board) error {
	for _, board := range boards {
		if board.Closed {
			continue
		}

		fmt.Printf("# %v (%v)\n\n", board.Name, board.ShortUrl)

		lists, err := board.Lists()
		if err != nil {
			return err
		}

		showLists(lists)

		fmt.Println()
	}

	return nil
}

func showLists(lists []trello.List) {
	for _, list := range lists {
		cards, _ := list.Cards()
		if len(cards) > 0 {
			fmt.Printf("## %s\n\n", list.Name)

			showCards(cards)
		}

		fmt.Println()
	}
}

func showCards(cards []trello.Card) {
	for _, card := range cards {
		for _, label := range card.Labels {
			setColor(label.Color)
			break
		}

		if card.Due != "" {
			t, _ := time.Parse(time.RFC3339Nano, card.Due)
			fmt.Printf("* %s %s\n", t.Format("[2006-01-02]"), card.Name)
		} else {
			fmt.Printf("* %s\n", card.Name)
		}

		color.Unset()
	}
}

func getBoards(appKey string, token string, username string) []trello.Board {
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

	return boards
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

		config, err := toml.LoadFile(home + "/.trelloclirc")
		if err != nil {
			log.Fatal(err)
		}

		appKey = config.Get("app_key").(string)
		token = config.Get("token").(string)
		username = config.Get("username").(string)
	}

	boards := getBoards(appKey, token, username)

	err := showBoards(boards)
	if err != nil {
		log.Fatal(err)
	}
}
