package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Send(mc tgbotapi.Chattable) (tgbotapi.Message, error) {
	var msg tgbotapi.Message
	bot, err := GetBot()
	if err == nil {
		msg, err = bot.Send(mc)
	}
	return msg, err
}

func SendMarkdownText(chatId int64, content string) (tgbotapi.Message, error) {
	chattable := tgbotapi.NewMessage(chatId, content)
	chattable.ParseMode = "markdown"
	return Send(chattable)
}

func SendPhotoWithCaption(chatId int64, url string, caption string) (tgbotapi.Message, error) {
    photo := tgbotapi.NewPhoto(chatId, tgbotapi.FileURL(url))
    photo.Caption = caption
    photo.ParseMode = "markdown"
    return Send(photo)
}