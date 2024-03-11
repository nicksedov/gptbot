package openai

import (
	"fmt"
	"testing"

	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/stretchr/testify/assert"
)

const TEST_CHAT_ID int64 = 5093432423
const CHAT_PROMPT string = "Представь себя команде"

func init() {
	*cli.FlagConfig = "../../settings-test.yaml"
}

func TestSendRequest(t *testing.T) {
	resp := SendRequest(TEST_CHAT_ID, CHAT_PROMPT)
	fmt.Printf("Response ID is %s\n", resp.ID)
	choices := resp.Choices
	assert.LessOrEqual(t, 1, len(choices))
	fmt.Printf("%s answered:\n - %s", choices[0].Message.Role, choices[0].Message.Content)
	
	var testHist []Messages = history[TEST_CHAT_ID]
	assert.NotNil(t, testHist)
	assert.Equal(t, 2, len(testHist))
}
