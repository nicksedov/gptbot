package gigachat

import (
	"fmt"
	"testing"

	"gptbot/pkg/cli"

	"github.com/stretchr/testify/assert"
)

const TEST_CHAT_ID int64 = 5093432423
const CHAT_PROMPT string = "Представь себя команде"

func init() {
	*cli.FlagConfig = "../../settings-test.yaml"
}

func TestSendRequest(t *testing.T) {
	client := GetClient()
	assert.NotNil(t, client)
	err := client.Auth()
	assert.Nil(t, err)
	resp, err := SendRequest(TEST_CHAT_ID, CHAT_PROMPT)
	assert.Nil(t, err)
	choices := resp.Choices
	assert.LessOrEqual(t, 1, len(choices))
	fmt.Printf("%s answered:\n - %s\n", choices[0].Message.Role, choices[0].Message.Content)

	var testHist []Message = history[TEST_CHAT_ID]
	assert.NotNil(t, testHist)
	assert.Equal(t, 2, len(testHist))
}
