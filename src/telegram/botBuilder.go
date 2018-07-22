package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"flag"
	"golang.org/x/net/proxy"
	"net/http"
	"log"
)

var (
	isProxyRequired bool
	token           string
	proxyUrl        string
	proxyUsername   string
	proxyPassword   string
)

func InitBot() (*tgbotapi.BotAPI, error) {
	flag.BoolVar(&isProxyRequired, "socks5", false, "is socks5 proxy required")
	flag.StringVar(&token, "token", "", "telegram bot access token")
	flag.StringVar(&proxyUrl, "proxy", "", "provide proxy URL: IP:PORT")
	flag.StringVar(&proxyUsername, "user", "", "proxy user")
	flag.StringVar(&proxyPassword, "password", "", "proxy password")
	flag.Parse()

	if token == "" {
		log.Panic("Access token is required")
	}

	if isProxyRequired {

		log.Printf("Socks5 config: url %s, username: %s", proxyUrl, proxyUsername)

		credentials := proxy.Auth{
			User:     proxyUsername,
			Password: proxyPassword,
		}

		dialer, err := proxy.SOCKS5("tcp", proxyUrl, &credentials, proxy.Direct)

		if err != nil {
			log.Fatal(err)
		}

		httpTransport := &http.Transport{}
		httpClient := &http.Client{Transport: httpTransport}
		httpTransport.Dial = dialer.Dial

		return tgbotapi.NewBotAPIWithClient(token, httpClient)
	} else {
		return tgbotapi.NewBotAPI(token)
	}
}
