package openai

type ImageResponse struct {
	HttpStatusCode int `json:"status"`
	HttpStatusMessage string `json:"status_msg"`
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}