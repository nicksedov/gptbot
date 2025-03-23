package localai

import (
	"errors"
	"gptbot/pkg/model"
	"strings"
    "log"

	openai "github.com/gopenai/openai-client"
)

func GetMessageByPrompt(e *model.SingleEvent) (text string, imageURL string, err error) {
    prompt, pErr := e.GetResolvedPrompt()
    if pErr != nil {
        altText, altErr := e.GetAltText(pErr)
        return altText, "", errors.Join(pErr, altErr)
    }
    
    resp, err := SendRequest(0, prompt)
    if err != nil {
        altText, _ := e.GetAltText(err)
        return altText, "", err
    }
    
    return GetResponseContent(resp)
}

func GetResponseContent(resp *openai.CreateChatCompletionResponse) (string, string, error) {
    for _, choice := range resp.Choices {
        content := choice.Message.Value.Content
        imageURL := "" // Логика извлечения URL из ответа LocalAI
        
        // Если LocalAI возвращает URL в специальном формате
        if strings.HasPrefix(content, "![image](") && strings.HasSuffix(content, ")") {
            start := strings.Index(content, "(") + 1
            end := strings.Index(content, ")")
            imageURL = content[start:end]
            content = ""
            log.Printf("Chat completion endpoint responded with image link: %s\n", imageURL)
        }
        
        normalizedContent := strings.TrimSpace(content)
        if normalizedContent != "" || imageURL != "" {
            return normalizedContent, imageURL, nil
        }
    }
    return "", "", errors.New("unexpectedly blank response obtained")
}