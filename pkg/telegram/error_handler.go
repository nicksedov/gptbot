package telegram

import (
	"errors"
	"fmt"
	"gptbot/pkg/settings"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ErrorReporter(err error) {
	settings := settings.GetSettings()
	serviceChat := settings.Telegram.ServiceChat
	if serviceChat != 0 {
		msg := tgbotapi.NewMessage(serviceChat, fmt.Sprintf("Exception occured on message processing\n```Error\n%v```", err))
		msg.ParseMode = "markdown"
		if bot == nil {
			initErr := initBot()
			if initErr != nil {
				log.Printf("Cannot create bot API:\n%v", errors.Join(err, initErr))
				return
			}
		}
		_, reportErr := bot.Send(msg)
		if reportErr != nil {
			log.Printf("Exception occured on message processing\n%v", errors.Join(err, reportErr))
		}
	} else {
		log.Printf("Exception occured on message processing\n%v", err)
	}
}
