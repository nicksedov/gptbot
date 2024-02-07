package telegram

import (
	"testing"

	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/stretchr/testify/assert"
)

func TestMessaging(t *testing.T) {
	
	initTestConfiguration()

	bot, err := GetBot()

	assert.Nil(t, err)
	assert.NotNil(t, bot)
	assert.Equal(t, bot.Token, *cli.FlagBotToken)
}