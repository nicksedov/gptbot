package openai

import (
	"encoding/json"
	"fmt"
	"io"
)

func SendImageRequest(dalle Model, size Size, n int, prompt string) *ImageResponse {
	reqData := prepareImageRequest(dalle, size, n, prompt)
	response, error := DoPost(IMAGE_API, reqData)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	//response.Status
	resp := &ImageResponse{HttpStatusCode: response.StatusCode, HttpStatusMessage: response.Status}
	error = json.Unmarshal(body, resp)
	if error != nil {
		fmt.Println(error)
		return resp
	}
	return resp
}



func prepareImageRequest(dalle Model, size Size, n int, prompt string) *ImageRequest {
	return &ImageRequest{
		Model:          string(dalle),
		Prompt:         prompt,
		N:              n,
		Size:           string(size),
		ResponseFormat: "url",
	}
}
