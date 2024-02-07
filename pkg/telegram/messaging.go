package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessageToChat(mc tgbotapi.Chattable) (tgbotapi.Message, error) {
	var msg tgbotapi.Message
	bot, err := GetBot()
	if err == nil {
		msg, err = bot.Send(mc)
	}
	return msg, err
}

func SendMarkdownText(content string, chatId int64) (tgbotapi.Message, error) {
	chattable := tgbotapi.NewMessage(chatId, content)
	chattable.ParseMode = "markdown"
	return SendMessageToChat(chattable)
}
