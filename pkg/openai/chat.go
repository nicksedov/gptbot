package openai

import (
	"encoding/json"
	"fmt"
	"io"
)

var history map[int64][]Messages = make(map[int64][]Messages)
var historyDepth int = 8

func SendRequest(chatId int64, prompt string) *ChatResponse {
	reqData := prepareRequest(chatId, prompt)
	response, error := DoPost(CHAT_API, reqData)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Printf("Error processing OpenAI API response: %s", error)
		panic(error)
	}
	resp := &ChatResponse{}
	error = json.Unmarshal(body, resp)
	if error != nil {
		fmt.Println(error)
		return resp
	}
	processResponse(chatId, resp)
	return resp
}

func updateHistory(userId int64, role string, content string) {
	userHist := history[userId]
	if userHist == nil {
		userHist = []Messages{}
	} else if len(userHist) >= historyDepth {
		userHist = userHist[len(userHist)-historyDepth:]
	}
	userHist = append(userHist, Messages{Role: role, Content: content})
	history[userId] = userHist
}

func prepareRequest(chatId int64, content string) *ChatRequest {
	updateHistory(chatId, "user", content)
	req := ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: history[chatId],
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
