package telegram

import (
	"encoding/json"
	"log"
	"strings"

	//ai "gptbot/pkg/gigachat"
	ai "gptbot/pkg/localai"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func updatesListener(updates tgbotapi.UpdatesChannel, errHandler func(error)) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		// execution thread locks until event received
		update := <-updates
		if update.Message != nil {
			err := handleMessage(update.Message)
			if err != nil {
				errHandler(err)
			}
		}
	}
}

func handleMessage(message *tgbotapi.Message) error {
	logMessage(message)
	text := message.Text
	chatId := message.Chat.ID
	if chatId == 0 {
		return nil
	}
	text = strings.TrimSpace(text)
	if text != "" {
		return processChat(chatId, text)
	}
	return nil
}

func processChat(chatId int64, prompt string) error {
	resp, reqErr := ai.SendRequest(chatId, prompt)
	if reqErr == nil {
		respContent, respErr := ai.GetResponseContent(resp)
		if respErr == nil {
			msg := tgbotapi.NewMessage(chatId, respContent)
			bot.Send(msg)
			return nil
		} else {
			return respErr
		}
	} else {
		return reqErr
	}
}

func logMessage(message *tgbotapi.Message) {
	content, _ := json.MarshalIndent(message, " > ", "  ")
	log.Printf("New message received:\n > %s", string(content))
}
