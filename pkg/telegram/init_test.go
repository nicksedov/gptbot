package telegram

import (
	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/nicksedov/gptbot/pkg/settings"
)


func initTestConfiguration() {
	secrets :=settings.GetSecrets("../../secrets.yaml")
	*cli.FlagBotToken = secrets.BotToken
}