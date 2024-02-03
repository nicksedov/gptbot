package openai

import "github.com/nicksedov/gptbot/pkg/model"

func GetMessageByPrompt(e *model.SingleEvent) string {
	prompt, err := e.GetResolvedPrompt()
	if err != nil {
		return e.Prompt.AltText
	}
	resp := SendRequest(0, prompt)
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content
	} else {
		return e.Prompt.AltText
	}
}
