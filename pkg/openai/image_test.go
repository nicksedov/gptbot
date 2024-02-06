package openai

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var IMAGE_PROMPT_EXAMPLES = [...]string{
	"Draw a cat in black and white stripes",
}

func TestSendImageRequest(t *testing.T) {

	initTestConfiguration()

	// Prepare test conditions
	expectedNum := 1
	prompt := IMAGE_PROMPT_EXAMPLES[0]

	// Run test method
	resp := SendImageRequest(DALLE_2, DALLE_2_MID, expectedNum, prompt)

	// Perform test checks 
	assert.NotNil(t, resp)
	actualData := resp.Data
	assert.Equal(t, "200 OK", resp.HttpStatusMessage)
	assert.Equal(t, expectedNum, len(actualData))

	// Print method execution results
	fmt.Printf("%d picture(s) for prompt '%s' are successfully created\n", len(actualData), prompt)
	for i := 0; i < len(actualData); i++ {
		fmt.Printf("Picture #%d url: [%s]\n", i+1, actualData[i].Url)
	}
}
