package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/gommon/log"
	"github.com/cobaku/tg-vrec-bot/src/utils"
)

type VoiceMessageHandler struct {
	getFileUrl func(fileId string) (string, error)
}

func NewVoiceMessageHandler(getFileUrl func(fileId string) (string, error)) *VoiceMessageHandler {
	return &VoiceMessageHandler{
		getFileUrl: getFileUrl,
	}
}

func (handler *VoiceMessageHandler) Handle(message tgbotapi.Update) string {
	fileUrl, _ := handler.getFileUrl(message.Message.Voice.FileID)
	err := utils.DownloadFile(message.Message.From.UserName+"-"+message.Message.Voice.FileID+".oga", fileUrl)

	if err != nil {
		log.Print(err)
	}

	return ""
}
