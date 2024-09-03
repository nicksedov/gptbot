package scheduler

import (
	"gptbot/pkg/model"
	//"gptbot/pkg/settings"
	"gptbot/pkg/telegram"
)

func handle(event *model.SingleEvent, getMessageByPrompt func(e *model.SingleEvent) (string, error)) error {
	chatId := event.Chat.ChatID
	msg, err := getMessageByPrompt(event)
	if err == nil {
		_, err = telegram.SendMarkdownText(chatId, msg)
	}
	return err
}

func onError(err error) {
	//serviceChatId := settings.GetSettings().Telegram.ServiceChat
	//telegram.SendMarkdownText("Error occured while processing event with LocalAI handler")
}
