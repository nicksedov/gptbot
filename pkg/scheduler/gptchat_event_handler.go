package scheduler

import (
	ai "gptbot/pkg/gigachat"
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
)

type ChatEventHandler struct{}

func (h *ChatEventHandler) handle(event *model.SingleEvent) error {
	chatId := event.Chat.ChatID
	msg := ai.GetMessageByPrompt(event)
	telegram.SendMarkdownText(msg, chatId)
	return nil
}

func (h *ChatEventHandler) onError(e error) {}
