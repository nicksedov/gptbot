package telegram

import (
	"errors"
	"fmt"
	"gptbot/pkg/settings"
	"log"
)

func ErrorReporter(err error) {
	settings := settings.GetSettings()
	serviceChat := settings.Telegram.ServiceChat
	if serviceChat != 0 {
		respContent := fmt.Sprintf("Exception occured on message processing\n```Error\n%v```", err)
		_, reportErr := SendMarkdownText(serviceChat, respContent)
		if reportErr != nil {
			log.Printf("Exception occured on message processing\n%v", errors.Join(err, reportErr))
		}
	} else {
		log.Printf("Exception occured on message processing\n%v", err)
	}
}
