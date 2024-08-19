package scheduler

import (
	ai "gptbot/pkg/gigachat"
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
)

type GigaChatEventHandler struct{}

func (h *GigaChatEventHandler) handle(event *model.SingleEvent) error {
	chatId := event.Chat.ChatID
	msg := ai.GetMessageByPrompt(event)
	telegram.SendMarkdownText(msg, chatId)
	return nil
}

func (h *GigaChatEventHandler) onError(e error) {}
