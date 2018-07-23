package config

import "flag"

type Config struct {
	IsProxyRequired bool
	Token           string
	ProxyUrl        string
	ProxyUsername   string
	ProxyPassword   string
}

var CurrentConfig Config

func InitConfig() {
	proxyUrl := flag.String("proxy", "", "provide proxy URL: IP:PORT")
	proxyUsername := flag.String("user", "", "proxy user")
	proxyPassword := flag.String("password", "", "proxy password")
	isProxyRequired := flag.Bool("socks5", false, "is socks5 proxy required")
	token := flag.String("token", "", "telegram bot access token")

	flag.Parse()

	CurrentConfig = Config{
		IsProxyRequired: *isProxyRequired,
		Token:           *token,
		ProxyUrl:        *proxyUrl,
		ProxyUsername:   *proxyUsername,
		ProxyPassword:   *proxyPassword,
	}
}
