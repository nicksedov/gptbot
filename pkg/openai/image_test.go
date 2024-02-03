package openai

import (
	"fmt"
	"testing"
)

func TestSendImageRequest(t *testing.T) {
	
	InitTestCfg()
	
	prompt := "Ugly cow is screaming"
	resp := SendImageRequest(prompt)
	choices := resp.Data
	if len(choices) > 0 {
		fmt.Printf("Picture %s available by url: [%s]", prompt, choices[0].Url)
	} else {
		fmt.Printf("Test failed; response HTTP status = %s\n", resp.HttpStatusMessage)
	}
}