package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicksedov/gptbot/pkg/cli"
)

const (
	BASE_URL = "https://api.openai.com/v1"
	CHAT_API = "/chat/completions"
	IMAGE_API = "/images/generations"
)

func DoPost(api string, jsonStruct any) *http.Response {

	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	
	url := BASE_URL + api
	token := *cli.FlagOpenAIToken
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer " + token)
	client := http.DefaultClient
	response, error := client.Do(request)
	if error != nil {
		fmt.Printf("Error calling OpenAI API: %s", error)
		return nil
	}
	return response
}