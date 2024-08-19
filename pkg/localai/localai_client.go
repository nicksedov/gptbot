package localai

import (
	"context"
	"gptbot/pkg/settings"
	"log"

	openai "github.com/gopenai/openai-client"
)

var client *openai.Client
var history map[int64][]openai.ChatCompletionRequestMessage = make(map[int64][]openai.ChatCompletionRequestMessage)
const historyDepth int = 8

func SendRequest(chatId int64, prompt string) *openai.CreateChatCompletionResponse {
	req := prepareRequest(chatId, prompt)
	if client == nil {
		var err error
		client, err = openai.NewClient("http://localhost:5555/")
		if err != nil {
			log.Panicln("LocalAI connection error", err)
		} else {
			log.Println("LocalAI client initialized successfully")
		}
	}
	
	response, error := client.CreateChatCompletion(context.Background(), req)

	if error != nil {
		log.Panicf("LocalAI API call error. Reason: %s", error.Error())
	}
	processResponse(chatId, response)
	return response
}


func updateHistory(chatId int64, role openai.ChatCompletionRequestMessageRole, content string) {
	userHist := history[chatId]
	if userHist == nil {
		userHist = []openai.ChatCompletionRequestMessage{}
	} else if len(userHist) >= historyDepth {
		userHist = userHist[len(userHist)-historyDepth:]
	}
	userHist = append(userHist, openai.ChatCompletionRequestMessage{Role: role, Content: content})
	history[chatId] = userHist
}

func prepareRequest(chatId int64, content string) *openai.CreateChatCompletionRequest {
	updateHistory(chatId, openai.ChatCompletionRequestMessageRoleUser, content)
	var messages []openai.ChatCompletionRequestMessage
	contextDescription := settings.GetSettings().GigaChat.Completions.Context
	if contextDescription != "" {
		messages = make( []openai.ChatCompletionRequestMessage, 0, len(history[chatId])+1)
		messages = append(messages, openai.ChatCompletionRequestMessage{
			Role: openai.ChatCompletionRequestMessageRoleSystem, Content: contextDescription,
		})
		messages = append(messages, history[chatId]...)
	} else {
		messages = history[chatId]
	}
	llmCfg := settings.GetSettings().GigaChat.LLMConfig
	req := &openai.CreateChatCompletionRequest{
		//Model:             llmCfg.Model,
		Model:             "saiga_gemma2_10b-full",
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
		updateHistory(chatId, openai.ChatCompletionRequestMessageRole(msg.Message.Value.Role), msg.Message.Value.Content)
	}
}
