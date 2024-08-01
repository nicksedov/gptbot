package scheduler

import (
	"gptbot/pkg/model"
	openai "gptbot/pkg/gigachat"
	"gptbot/pkg/telegram"
)

type GptChatEventHandler struct{}

func (h *GptChatEventHandler) handle(event *model.SingleEvent) error {
	chatId := event.Chat.ChatID
	msg := openai.GetMessageByPrompt(event)
	telegram.SendMarkdownText(msg, chatId)
	return nil
}

func (h *GptChatEventHandler) onError(e error) {}
