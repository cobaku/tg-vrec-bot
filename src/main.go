package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/cobaku/tg-vrec-bot/src/telegram"
	"log"
)

func main() {

	bot, err := telegram.InitBot()

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
