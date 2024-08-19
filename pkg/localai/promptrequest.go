package localai

import (
	"gptbot/pkg/model"

	openai "github.com/gopenai/openai-client"
)
func GetMessageByPrompt(e *model.SingleEvent) string {
	prompt, err := e.GetResolvedPrompt()
	if err != nil {
		return e.Prompt.AltText
	}
	resp := SendRequest(0, prompt)
	if len(resp.Choices) > 0 {
		return GetResponseContent(resp)
	} else {
		return e.Prompt.AltText
	}
}

func GetResponseContent(resp *openai.CreateChatCompletionResponse) string {
	return resp.Choices[0].Message.Value.Content
}
