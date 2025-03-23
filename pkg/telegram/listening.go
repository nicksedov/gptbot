package telegram

import (
	"encoding/json"
	"log"
	"strings"
	"errors"

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
    if reqErr != nil {
        return reqErr
    }
    
    textResponse, imageURL, respErr := ai.GetResponseContent(resp)
    if respErr != nil {
        return respErr
    }
    
    if imageURL != "" {
        _, sendErr := SendPhotoWithCaption(chatId, imageURL, textResponse)
        return sendErr
    }
    
    if textResponse != "" {
        _, sendErr := SendMarkdownText(chatId, textResponse)
        return sendErr
    }
    
    return errors.New("empty response from AI")
}

func logMessage(message *tgbotapi.Message) {
	content, _ := json.MarshalIndent(message, " > ", "  ")
	log.Printf("New message received:\n > %s", string(content))
}
