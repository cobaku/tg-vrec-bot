package utils

import (
	"golang.org/x/net/proxy"
	"net/http"
	"log"
	"github.com/cobaku/tg-vrec-bot/src/config"
	"os"
	"fmt"
	"io"
)

func BuildClientWithProxy() *http.Client {
	credentials := proxy.Auth{
		User:     config.CurrentConfig.ProxyUsername,
		Password: config.CurrentConfig.ProxyPassword,
	}

	dialer, err := proxy.SOCKS5("tcp", config.CurrentConfig.ProxyUrl, &credentials, proxy.Direct)

	if err != nil {
		log.Fatal(err)
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	return httpClient
}

func DownloadFile(filePath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	var httpClient *http.Client

	if config.CurrentConfig.IsProxyRequired {
		httpClient = BuildClientWithProxy()
	} else {
		httpClient = &http.Client{}
	}

	// Get the data
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
