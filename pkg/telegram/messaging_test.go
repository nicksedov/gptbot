package telegram

import (
	"fmt"
	"testing"

	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/stretchr/testify/assert"
)

const TEST_CHAT_ID int64 = -1001787389818

func TestBotInit(t *testing.T) {

	initTestConfiguration()

	bot, err := GetBot()

	assert.Nil(t, err)
	assert.NotNil(t, bot)
	assert.Equal(t, bot.Token, *cli.FlagBotToken)
}
func TestSendMarkdownMessage(t *testing.T) {

	initTestConfiguration()

	text := fmt.Sprintf("Это сообщение отправлено из unit-теста.\n"+
		"Проект: ```%s```\nМетод: ```%s```\nРасположение: ```%s```",
		"gptbot", "NtTestMessaging()", "gptbot/pkg/telegram/messaging_test.go")
	msg, err := SendMarkdownText(text, TEST_CHAT_ID)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, TEST_CHAT_ID, msg.Chat.ID)
}
