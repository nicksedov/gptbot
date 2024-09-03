package scheduler

import (
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
)

func handle(event *model.SingleEvent, getMessageByPrompt func(e *model.SingleEvent) (string, error)) error {
	chatId := event.Chat.ChatID
	msg, err := getMessageByPrompt(event)
	if msg != "" {
		_, sendErr := telegram.SendMarkdownText(chatId, msg)
		if sendErr != nil {
			return sendErr
		}
	}
	return err
}
