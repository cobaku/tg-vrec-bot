package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	IsProxyRequired bool   `json:"socks5"`
	Token           string `json:"token"`
	ProxyUrl        string `json:"proxy"`
	ProxyUsername   string `json:"user"`
	ProxyPassword   string `json:"password"`
	SpeechKitApiKey string `json:"speechKitApiKey"`
	Debug           bool   `json:"debug"`
}

var CurrentConfig Config

func InitConfig() {
	file, err := ioutil.ReadFile("../src/config/config.json")

	if err != nil {
		log.Panic(err)
	}

	json.Unmarshal([]byte(file), &CurrentConfig)
}
