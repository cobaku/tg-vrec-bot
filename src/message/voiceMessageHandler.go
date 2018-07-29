package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/gommon/log"
	"github.com/cobaku/tg-vrec-bot/src/utils"
	"os/exec"
	"strings"
	"encoding/xml"
	"github.com/cobaku/tg-vrec-bot/src/dto"
	"github.com/cobaku/tg-vrec-bot/src/config"
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

	filename := message.Message.From.UserName + "-" + message.Message.Voice.FileID + ".oga"
	err := utils.DownloadFile(filename, fileUrl)

	if err != nil {
		log.Print(err)
	}

	cmd := exec.Command("ffmpeg", "-i", filename, filename+".wav")
	cmd.Run()

	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		log.Panic(err)
	}

	uuidString := strings.Replace(string(uuid), "-", "", -1)
	uuidString = strings.TrimSuffix(uuidString, "\n")

	resp, err := utils.UploadFile(filename+".wav", "https://asr.yandex.net/asr_xml?uuid="+uuidString+"&key="+config.CurrentConfig.SpeechKitApiKey+"&topic=queries", "audio/x-wav", false)

	if err != nil {
		log.Print(err)
	}

	result := dto.YandexXmlBody{}

	err = xml.Unmarshal(resp, &result)

	if err != nil {
		log.Print(err)
	}

	if len(result.Variants) == 0 {
		return "Не могу распознать текст"
	}

	return result.Variants[0].Value
}
