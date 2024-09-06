package scheduler

import (
	"errors"
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
)

func handle(event *model.SingleEvent, composeMessage func(e *model.SingleEvent) (string, error)) error {
	chatId := event.Chat.ChatID
	msg, err := composeMessage(event)
	if msg != "" {
		_, sendErr := telegram.SendMarkdownText(chatId, msg)
		if sendErr != nil {
			return sendErr
		}
	} else if err == nil {
		return errors.New("completions API returned empty response")
	}
	return err
}
