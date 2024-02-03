package openai

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/gptbot/pkg/cli"
)

type Secrets struct {
	BotToken string      `yaml:"BotToken"`
	OpenAIToken string   `yaml:"OpenAIToken"`
	ProxyHost string     `yaml:"ProxyHost"`
	ProxyUser string     `yaml:"ProxyUser"`
	ProxyPassword string `yaml:"ProxyPassword"`
}

func InitTestCfg() {
	secrets := getSecrets()
	*cli.FlagConfig = "../../gptbot-settings.yaml"
	*cli.FlagBotToken = secrets.BotToken
	*cli.FlagOpenAIToken = secrets.OpenAIToken
	*cli.ProxyHost = secrets.ProxyHost
	*cli.ProxyUser = secrets.ProxyUser
	*cli.ProxyPassword = secrets.ProxyPassword
}

func getSecrets() Secrets {
	secrets := Secrets{}
	yfile, ioErr := os.ReadFile("../../secrets.yaml")
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &secrets)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	return secrets
}