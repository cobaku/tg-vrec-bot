package telegram

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/cobaku/tg-vrec-bot/src/utils"
	"github.com/cobaku/tg-vrec-bot/src/config"
)

func InitBot() (*tgbotapi.BotAPI, error) {

	if config.CurrentConfig.Token == "" {
		log.Panic("Access token is required")
	}

	if config.CurrentConfig.IsProxyRequired {
		return tgbotapi.NewBotAPIWithClient(config.CurrentConfig.Token, utils.BuildClientWithProxy())
	} else {
		return tgbotapi.NewBotAPI(config.CurrentConfig.Token)
	}
}
