package openai

import (
	"github.com/nicksedov/gptbot/pkg/cli"
	"github.com/nicksedov/gptbot/pkg/settings"
)

func initTestConfiguration() {
	cliParams := settings.GetCliParamsFromFile("../../test_cli_params.yaml")
	*cli.FlagOpenAIToken = cliParams.OpenAIToken
	*cli.Proxy = cliParams.Proxy
	*cli.ProxyUser = cliParams.ProxyUser
	*cli.ProxyPassword = cliParams.ProxyPassword
}
