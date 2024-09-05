package scheduler

import (
	"errors"
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
	} else if err == nil {
		return errors.New("completions API returned empty response")
	}
	return err
}
