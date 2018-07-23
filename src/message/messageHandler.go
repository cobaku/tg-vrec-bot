package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramMessageHandler struct {
	inputChannel  tgbotapi.UpdatesChannel
	outputChannel chan tgbotapi.Chattable
}

var voiceMessageHandler VoiceMessageHandler

func NewTelegramMessageHandler(inputChannel tgbotapi.UpdatesChannel, getFileUrl func(fileId string) (string, error)) *TelegramMessageHandler {
	voiceMessageHandler = *NewVoiceMessageHandler(getFileUrl)
	return &TelegramMessageHandler{inputChannel: inputChannel, outputChannel: make(chan tgbotapi.Chattable, 50)}
}

func (handler *TelegramMessageHandler) Run() chan tgbotapi.Chattable {
	handler.startHandler()
	return handler.outputChannel
}

func (handler *TelegramMessageHandler) startHandler() {
	go func() {
		for update := range handler.inputChannel {
			handler.handleMessage(update)
		}
	}()
}

func (handler *TelegramMessageHandler) handleMessage(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Text != "" {
		go handler.handleTextMessage(update)
	}

	if update.Message.Voice != nil {
		go handler.handleVoiceMessage(update)
	}
}

func (handler *TelegramMessageHandler) handleVoiceMessage(message tgbotapi.Update) {
	voiceMessageText := voiceMessageHandler.Handle(message)
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, voiceMessageText)
	msg.ReplyToMessageID = message.Message.MessageID
	handler.outputChannel <- msg
}

func (handler *TelegramMessageHandler) handleTextMessage(message tgbotapi.Update) {
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, "К сожалению, пока не могу ничем помочь")
	msg.ReplyToMessageID = message.Message.MessageID
	handler.outputChannel <- msg
}
