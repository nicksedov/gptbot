package localai

import (
	"context"
	"encoding/json"
	"fmt"
	"gptbot/pkg/cli"
	"gptbot/pkg/settings"
	"testing"
	"time"

	openai "github.com/gopenai/openai-client"
	"github.com/stretchr/testify/assert"
)

func init() {
	*cli.FlagConfig = "../../settings-test.yaml"
}

func TestBackend(t *testing.T) {
	settings := settings.GetSettings()
	client, err := openai.NewBearerAuthClient(settings.LocalAI.URL, settings.LocalAI.APIKey)
	if err != nil {
		panic(err)
	}
	resp, err := client.ListModels(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, resp.Data)
	for _, model := range resp.Data {
		var objmap map[string]string
		err := json.Unmarshal(model, &objmap)
		assert.Nil(t, err)
		modelName := objmap["id"]
		model, _ := client.RetrieveModel(context.Background(), openai.RetrieveModelParams{Model: modelName})
		assert.Nil(t, err)
		fmt.Println(string(model))
	}
}

// Warning: this is rather long lasting test, use DEBUG mode when launching from VSCode IDE
func TestSendRequest(t *testing.T) {
	settings := settings.GetSettings()
	client, err := openai.NewBearerAuthClient(settings.LocalAI.URL, settings.LocalAI.APIKey)
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Напомни команде, что завтра feature freeze релиза Release-20240914. Стоит обсудить с тимлидом задачи, которые готовы для добавления в состав релиза."

	req := &openai.CreateChatCompletionRequest{
		Model: settings.LocalAI.Model,
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    openai.ChatCompletionRequestMessageRoleSystem,
				Content: settings.LLMConfig.Completions.Context,
			},
			{
				Role:    openai.ChatCompletionRequestMessageRoleUser,
				Content: prompt,
			},
			{
				Role:    openai.ChatCompletionRequestMessageRoleAssistant,
				Content: "...",
			},
		},
	}
	startTime := time.Now()
	response, err := client.CreateChatCompletion(context.Background(), req)
	assert.Nil(t, err)
	duration := time.Since(startTime)

	// Print the completed text
	fmt.Printf("Время работы: %s\nРезультат: %s\n", duration, response.Choices[0].Message.Value.Content)
}

func TestSendRequest2(t *testing.T) {
	settings := settings.GetSettings()
	testChatId := settings.Telegram.ServiceChat
	settings.LLMConfig.Completions.Context = "Ты - специалист по истории Франции."
	assert.NotEqual(t, 0, testChatId)
	req1 := "Расскажи, кто возглавлял Францию в 1812 году"
	req2 := "В каких битвах он принимал участие в этот год?"
	res1, err := SendRequest(testChatId, req1)
	assert.Nil(t, err)
	res2, err := SendRequest(testChatId, req2)
	assert.Nil(t, err)
	fmt.Printf("Вопрос:\n%s\nОтвет:\n%s\nВопрос:\n%s\nОтвет:\n%s\n", 
	req1, res1.Choices[0].Message.Value.Content, 
	req2, res2.Choices[0].Message.Value.Content)
}
