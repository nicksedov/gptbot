package telegram

import (
	"fmt"
	"testing"

	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/stretchr/testify/assert"
)

func TestBotInit(t *testing.T) {

	initTestConfiguration()

	bot, err := GetBot()

	assert.Nil(t, err)
	assert.NotNil(t, bot)
	assert.Equal(t, bot.Token, *cli.FlagBotToken)
}
func TestSendMarkdownMessage(t *testing.T) {

	testChatId := initTestConfiguration()

	text := fmt.Sprintf("Это сообщение отправлено из unit-теста.\n"+
		"Проект: ```%s```\nМетод: ```%s```\nРасположение: ```%s```",
		"gptbot", "TestSendMarkdownMessage()", "gptbot/pkg/telegram/messaging_test.go")
	msg, err := SendMarkdownText(text, testChatId)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, testChatId, msg.Chat.ID)
}
