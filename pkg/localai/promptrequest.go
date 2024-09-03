package localai

import (
	"errors"
	"gptbot/pkg/model"
	"strings"
	openai "github.com/gopenai/openai-client"
)

func GetMessageByPrompt(e *model.SingleEvent) (string, error) {
	prompt, pErr := e.GetResolvedPrompt()
	if pErr != nil {
		altText, altErr := e.GetAltText()
		return altText, errors.Join(pErr, altErr)
	}
	resp, err := SendRequest(0, prompt)
	if err != nil {
		return e.GetAltText()
	}
	respContent, err := GetResponseContent(resp)
	if err != nil {
		return e.GetAltText()
	}
	return respContent, nil
}

func GetResponseContent(resp *openai.CreateChatCompletionResponse) (string, error) {
	for _, choice := range resp.Choices {
		content := choice.Message.Value.Content
		normalizedContent := strings.TrimSpace(content)
		if normalizedContent != "" {
			return normalizedContent, nil
		}
	}
	return "", errors.New("unexpectedly blank response obtained")
}
