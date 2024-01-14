package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MessageHandler interface {
	handle(message *tgbotapi.Message)
}
