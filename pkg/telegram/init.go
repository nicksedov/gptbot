package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksedov/gptbot/pkg/cli"
)

var (
	bot *tgbotapi.BotAPI
)

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
	bot, err = tgbotapi.NewBotAPI(*cli.FlagBotToken)
	if err != nil {
		return fmt.Errorf("cannot create bot API: %w", err)
	}
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	handler := MessageChatGPTResponder{}
	go updatesListener(upd, handler)
	return nil
}

func updatesListener(upd tgbotapi.UpdateConfig, handler MessageHandler) {
	// `for {` means the loop is infinite until we manually stop it
	updatesChan := bot.GetUpdatesChan(upd)
	for {
		// execution thread locks until event received
		update := <-updatesChan
		if update.Message != nil {
			handler.handle(update.Message)
		}
	}
}
