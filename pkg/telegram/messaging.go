package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessageToChat(mc tgbotapi.Chattable) error {
	bot, err := GetBot()
	if err == nil {
		_, err = bot.Send(mc)
	}
	return err
}