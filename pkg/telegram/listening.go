package telegram

import (
	"encoding/json"
	"log"
	"strings"

	openai "gptbot/pkg/gigachat"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func updatesListener(updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		// execution thread locks until event received
		update := <-updates
		if update.Message != nil {
			handleMessage(update.Message)
		}
	}
}

func handleMessage(message *tgbotapi.Message) {
	logMessage(message)
	text := message.Text
	chatId := message.Chat.ID
	if chatId == 0 {
		return
	}
	text = strings.TrimSpace(text)
	if text != "" {
		processChat(chatId, text)
	}
}

func processChat(chatId int64, prompt string) {
	resp := openai.SendRequest(chatId, prompt)
	if len(resp.Choices) > 0 {
		msg := tgbotapi.NewMessage(chatId, resp.Choices[0].Message.Content)
		bot.Send(msg)
	}
}

func logMessage(message *tgbotapi.Message) {
	content, _ := json.MarshalIndent(message, " > ", "  ")
	log.Printf("New message received:\n > %s", string(content))
}
