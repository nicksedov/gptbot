package telegram

import (
	"fmt"
	"testing"

	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/nicksedov/gptbot/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func init() {
	*cli.FlagConfig = "../../settings-test.yaml"
}

func TestBotInit(t *testing.T) {
	settings := settings.GetSettings()
	testChatId := settings.Telegram.ServiceChat
	assert.NotEqual(t, 0, testChatId)
	bot, err := GetBot()
	assert.Nil(t, err)	
	assert.NotNil(t, bot)
	assert.Equal(t, bot.Token, settings.Telegram.BotToken)
}
func TestSendMarkdownMessage(t *testing.T) {
	settings := settings.GetSettings()
	testChatId := settings.Telegram.ServiceChat
	assert.NotEqual(t, 0, testChatId)
	text := fmt.Sprintf("Это сообщение отправлено из unit-теста.\n"+
		"Проект: ```%s```\nМетод: ```%s```\nРасположение: ```%s```",
		"gptbot", "TestSendMarkdownMessage()", "gptbot/pkg/telegram/messaging_test.go")
	msg, err := SendMarkdownText(text, testChatId)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, testChatId, msg.Chat.ID)
}
