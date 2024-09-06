package gigachat

import (
	"errors"
	"gptbot/pkg/model"
	"strings"
)

func GetMessageByPrompt(e *model.SingleEvent) (string, error) {
	prompt, pErr := e.GetResolvedPrompt()
	if pErr != nil {
		altText, altErr := e.GetAltText(pErr)
		return altText, errors.Join(pErr, altErr)
	}
	resp, err := SendRequest(0, prompt)
	if err != nil {
		return e.GetAltText(err)
	}
	respContent, err := GetResponseContent(resp)
	if err != nil {
		return e.GetAltText(err)
	}
	return respContent, nil
}

func GetResponseContent(resp *ChatResponse) (string, error) {
	for _, choice := range resp.Choices {
		content := choice.Message.Content
		normalizedContent := strings.TrimSpace(content)
		if normalizedContent != "" {
			return normalizedContent, nil
		}
	}
	return "", errors.New("unexpectedly blank response obtained")
}