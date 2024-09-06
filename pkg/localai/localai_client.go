package localai

import (
	"context"
	"gptbot/pkg/settings"
	"log"

	openai "github.com/gopenai/openai-client"
)

var client *openai.Client
var history map[int64][]Message = make(map[int64][]Message)
const historyDepth int = 8

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func SendRequest(chatId int64, prompt string) (*openai.CreateChatCompletionResponse, error) {
	req := prepareRequest(chatId, prompt)
	if client == nil {
		var err error
		url := settings.GetSettings().LocalAI.URL
		client, err = openai.NewClient(url)
		if err != nil {
			return nil, err
		} else {
			log.Println("LocalAI client initialized successfully")
		}
	}
	
	response, error := client.CreateChatCompletion(context.Background(), req)
	if error != nil {
		return nil, error
	}
	processResponse(chatId, response)
	return response, nil
}


func updateHistory[T openai.ChatCompletionRequestMessageRole|openai.ChatCompletionResponseMessageRole](chatId int64, role T, content string) {
	userHist := history[chatId]
	if userHist == nil {
		userHist = []Message{}
	} else if len(userHist) >= historyDepth {
		userHist = userHist[len(userHist)-historyDepth:]
	}
	userHist = append(userHist, Message{Role: string(role), Content: content})
	history[chatId] = userHist
}

func prepareRequest(chatId int64, content string) *openai.CreateChatCompletionRequest {
	updateHistory(chatId, openai.ChatCompletionRequestMessageRoleUser, content)
	var messages []openai.ChatCompletionRequestMessage
	llmCfg := settings.GetSettings().LLMConfig
	contextDescription := llmCfg.Completions.Context
	if contextDescription != "" {
		messages = make( []openai.ChatCompletionRequestMessage, 0, len(history[chatId])+1)
		messages = append(messages, openai.ChatCompletionRequestMessage{
			Role: openai.ChatCompletionRequestMessageRoleSystem, Content: contextDescription,
		})
	}
	for _, message := range history[chatId] {
		messages = append(messages, openai.ChatCompletionRequestMessage{
			Role: openai.ChatCompletionRequestMessageRole(message.Role),
			Content: message.Content,
		})
	}
	messages = append(messages, openai.ChatCompletionRequestMessage{
		Role: openai.ChatCompletionRequestMessageRoleAssistant,
		Content: "",
	})
	
	req := &openai.CreateChatCompletionRequest{
		Model:             settings.GetSettings().LocalAI.Model,
		Temperature:       openai.NewOptNilFloat64(llmCfg.Temperature),
		TopP:              openai.NewOptNilFloat64(llmCfg.TopP),
		N:                 openai.NewOptNilInt(1),
		Stream:            openai.NewOptNilBool(false),
		MaxTokens:         openai.NewOptInt(int(llmCfg.MaxTokens)),
		FrequencyPenalty:  openai.NewOptNilFloat64(llmCfg.RepetitionPenalty),
		Messages:          messages,
	}
	return req
}

func processResponse(chatId int64, resp *openai.CreateChatCompletionResponse) {
	choices := resp.Choices
	if len(choices) > 0 {
		msg := choices[0]
		updateHistory(chatId, msg.Message.Value.Role, msg.Message.Value.Content)
	}
}
