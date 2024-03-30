package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gptbot/pkg/settings"
)

const (
	BASE_URL  = "https://api.openai.com/v1"
	CHAT_API  = "/chat/completions"
	IMAGE_API = "/images/generations"
)

var httpClient *http.Client

func GetClient() *http.Client {
	settings := settings.GetSettings()
	if httpClient == nil {
		transport := &http.Transport{}
		if settings.Proxy.Host != "" {
			proxyUrl := &url.URL{
				Scheme: "http",
				User:   url.UserPassword(settings.Proxy.User, settings.Proxy.Password),
				Host:   settings.Proxy.Host + ":" + fmt.Sprint(settings.Proxy.Port),
			}
			transport.Proxy = http.ProxyURL(proxyUrl) // set proxy
		}
		httpClient = &http.Client{Transport: transport}
	}
	return httpClient
}

func DoGet(api string) (*http.Response, error) {

	url, err := url.JoinPath(BASE_URL, api)
	if err != nil {
		return handleError("Wrong API name format", err)
	}

	settings := settings.GetSettings()
	token := settings.OpenAI.APIToken
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return handleError("Error building OpenAI HTTP request", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
	client := GetClient()
	response, err := client.Do(request)
	if err != nil {
		return handleError("Error calling OpenAI API", err)
	}

	return response, nil
}

func DoPost(api string, jsonData any) (*http.Response, error) {

	jsonBody, err := json.Marshal(jsonData)
	if err != nil {
		return handleError("Error serializing object to JSON", err)
	}

	url, err := url.JoinPath(BASE_URL, api)
	if err != nil {
		return handleError("Wrong API name format", err)
	}

	settings := settings.GetSettings()
	token := settings.OpenAI.APIToken
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return handleError("Error building OpenAI HTTP request", err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)
	client := GetClient()
	response, err := client.Do(request)
	if err != nil {
		return handleError("Error calling OpenAI API", err)
	}

	return response, nil
}

func handleError(summary string, err error) (*http.Response, error) {
	fmt.Printf("%s: %v", summary, err)
	return nil, err
}
