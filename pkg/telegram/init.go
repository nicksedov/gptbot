package telegram

import (
	"errors"
	"fmt"
	"log"

	"gptbot/pkg/settings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func GetBot() (*tgbotapi.BotAPI, error) {
	if bot == nil {
		err := initBot()
		if err != nil {
			return nil, fmt.Errorf("cannot create bot API: %w", err)
		}
	}
	return bot, nil
}

func initBot() error {
	var err error
	settings := settings.GetSettings()
	bot, err = tgbotapi.NewBotAPI(settings.Telegram.BotToken)
	if err != nil {
		return fmt.Errorf("cannot create bot API: %w", err)
	}
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	go updatesListener(bot.GetUpdatesChan(upd), errorHandler)
	return nil
}

func errorHandler(err error) {
	settings := settings.GetSettings()
	msg := tgbotapi.NewMessage(settings.Telegram.ServiceChat,  fmt.Sprintf("Exception occured on message processing\n%v", err))
	_, reportErr := bot.Send(msg)
	if reportErr != nil {
		log.Printf("Exception occured on message processing\n%v", errors.Join(err, reportErr))
	}
}