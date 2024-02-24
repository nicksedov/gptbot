package cli

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type CommandLineFlags struct {
	BotToken      string `yaml:"BotToken"`
	OpenAIToken   string `yaml:"OpenAIToken"`
	Proxy         string `yaml:"Proxy"`
	ProxyUser     string `yaml:"ProxyUser"`
	ProxyPassword string `yaml:"ProxyPassword"`
	// Extension (parameters not present as command-line)
	ServiceChatID int64  `yaml:"ServiceChatID"`
}

func GetFlagsFromFile(path string) CommandLineFlags {
	cliParams := CommandLineFlags{}
	yfile, ioErr := os.ReadFile(path)
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &cliParams)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	*FlagBotToken = cliParams.BotToken
	*FlagOpenAIToken = cliParams.OpenAIToken
	*Proxy = cliParams.Proxy
	*ProxyUser = cliParams.ProxyUser
	*ProxyPassword = cliParams.ProxyPassword

	return cliParams
}
