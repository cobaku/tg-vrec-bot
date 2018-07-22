package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/cobaku/tg-vrec-bot/src/telegram"
	"github.com/cobaku/tg-vrec-bot/src/telegram/message"
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

	channel, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for outgoingMessage := range message.NewTelegramMessageHandler(channel).Run() {
		bot.Send(outgoingMessage)
	}
}
