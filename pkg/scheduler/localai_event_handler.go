package scheduler

import (
	ai "gptbot/pkg/localai"
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
)

type LocalAIEventHandler struct{}

func (h *LocalAIEventHandler) handle(event *model.SingleEvent) error {
	chatId := event.Chat.ChatID
	msg := ai.GetMessageByPrompt(event)
	telegram.SendMarkdownText(msg, chatId)
	return nil
}

func (h *LocalAIEventHandler) onError(e error) {}
