package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramMessageHandler struct {
	inputChannel  tgbotapi.UpdatesChannel
	outputChannel chan tgbotapi.Chattable
}

func NewTelegramMessageHandler(inputChannel tgbotapi.UpdatesChannel) *TelegramMessageHandler {
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
		handler.handleTextMessage(update)
	}

	if update.Message.Voice != nil {
		handler.handleVoiceMessage(update)
	}
}

func (handler *TelegramMessageHandler) handleVoiceMessage(message tgbotapi.Update) {
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, "Hi")
	msg.ReplyToMessageID = message.Message.MessageID
	handler.outputChannel <- msg
}

func (handler *TelegramMessageHandler) handleTextMessage(message tgbotapi.Update) {
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, "К сожалению, пока не могу ничем помочь")
	msg.ReplyToMessageID = message.Message.MessageID
	handler.outputChannel <- msg
}
