package openai

import (
	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/nicksedov/gptbot/pkg/settings"
)


func initTestConfiguration() {
	secrets :=settings.GetSecrets("../../secrets.yaml")
	*cli.FlagOpenAIToken = secrets.OpenAIToken
	*cli.ProxyHost = secrets.ProxyHost
	*cli.ProxyUser = secrets.ProxyUser
	*cli.ProxyPassword = secrets.ProxyPassword
}

