package gigachat

import (
	"gptbot/pkg/settings"
	"log"

	"github.com/paulrzcz/go-gigachat"
)

var client *gigachat.Client

func GetClient() *gigachat.Client {
	settings := settings.GetSettings()
	var err error
	if client == nil {
		client, err = gigachat.NewClient(settings.GigaChat.ClientID, settings.GigaChat.ClientSecret)
		if err != nil {
			log.Panicln("GigaChat connection error", err)
		}
	}
	return client
}
