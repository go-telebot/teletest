package main

import (
	"log"
	"os"

	tele "gopkg.in/telebot.v3"
)

var db = make(map[int64]bool)

func main() {
	b, err := Bot(tele.Settings{Token: os.Getenv("TOKEN")})
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}

func Bot(pref tele.Settings) (*tele.Bot, error) {
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	b.Handle("/start", func(c tele.Context) error {
		if !db[c.Sender().ID] {
			db[c.Sender().ID] = true
		}
		return c.Send("Hello, teletest!")
	})

	return b, nil
}
