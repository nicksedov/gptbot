package openai

import (
	"fmt"
	"testing"
)

func TestSendRequest(t *testing.T) {

	InitTestCfg()

	resp := SendRequest(5093432423, "Hello buddy!")
	fmt.Printf("Response ID is %s\n", resp.ID)
	choices := resp.Choices
	if len(choices) > 0 {
		fmt.Printf("%s answered:\n - %s", choices[0].Message.Role, choices[0].Message.Content)
	} else {
		fmt.Println("Test failed")
	}
}

