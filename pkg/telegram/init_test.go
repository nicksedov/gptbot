package telegram

import (
	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/nicksedov/gptbot/pkg/settings"
)

func initTestConfiguration() int64 {
	secrets := settings.GetSecrets("../../secrets.yaml")
	*cli.FlagBotToken = secrets.BotToken
	return secrets.TgChatID
}
