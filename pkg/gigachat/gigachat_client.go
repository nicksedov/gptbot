package gigachat

import (
	"gptbot/pkg/settings"
	"log"
	"time"
)

var client *Client
var history map[int64][]Message = make(map[int64][]Message)

const historyDepth int = 8

func GetClient() *Client {
	var err error
	if client == nil {
		settings := settings.GetSettings()
		client, err = NewInsecureClient(settings.GigaChat.ClientID, settings.GigaChat.ClientSecret)
		if err != nil {
			log.Panicln("GigaChat connection error", err)
		}
	}
	err = authCheck()
	if err != nil {
		log.Panicln("GigaChat client authorization error", err)
	}
	return client
}

func SendRequest(chatId int64, prompt string) *ChatResponse {
	reqData := prepareRequest(chatId, prompt)
	response, error := GetClient().Chat(reqData)
	if error != nil {
		log.Panicf("GigaChat API call error. Reason: %s", error.Error())
	}
	processResponse(chatId, response)
	return response
}

func authCheck() error {
	if client == nil || client.token == nil || time.Until(client.token.expiresAt) < time.Second * 10 {
		return client.Auth()
	}
	return nil
}

func updateHistory(chatId int64, role string, content string) {
	userHist := history[chatId]
	if userHist == nil {
		userHist = []Message{}
	} else if len(userHist) >= historyDepth {
		userHist = userHist[len(userHist)-historyDepth:]
	}
	userHist = append(userHist, Message{Role: role, Content: content})
	history[chatId] = userHist
}

func prepareRequest(chatId int64, content string) *ChatRequest {
	updateHistory(chatId, "user", content)
	var messages []Message
	contextDescription := settings.GetSettings().GigaChat.Completions.Context
	if contextDescription != "" {
		messages = make([]Message, 0, len(history[chatId])+1)
		messages = append(messages, Message{Role: "system", Content: contextDescription})
		messages = append(messages, history[chatId]...)
	} else {
		messages = history[chatId]
	}
	gcCfg := settings.GetSettings().GigaChat
	req := ChatRequest{
		Model:             gcCfg.Model,
		Messages:          messages,
		Temperature:       ptr(gcCfg.Temperature),
		TopP:              ptr(gcCfg.TopP),
		N:                 ptr(int64(1)),
		Stream:            ptr(false),
		MaxTokens:         ptr(gcCfg.MaxTokens),
		RepetitionPenalty: ptr(gcCfg.RepetitionPenalty),
		UpdateInterval:    ptr(int64(0)),
	}
	return &req
}

func processResponse(chatId int64, resp *ChatResponse) {
	choices := resp.Choices
	if len(choices) > 0 {
		msg := choices[0]
		updateHistory(chatId, msg.Message.Role, msg.Message.Content)
	}
}

func ptr[T any](v T) *T {
	return &v
}
