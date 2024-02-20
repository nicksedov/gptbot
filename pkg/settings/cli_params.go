package settings

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type CommandLineParams struct {
	BotToken      string `yaml:"BotToken"`
	OpenAIToken   string `yaml:"OpenAIToken"`
	Proxy         string `yaml:"Proxy"`
	ProxyUser     string `yaml:"ProxyUser"`
	ProxyPassword string `yaml:"ProxyPassword"`
	TgChatID      int64  `yaml:"TelegramChatID"`
}

func GetCliParamsFromFile(path string) CommandLineParams {
	cliParams := CommandLineParams{}
	yfile, ioErr := os.ReadFile(path)
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &cliParams)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	return cliParams
}
