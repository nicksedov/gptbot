package settings

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Secrets struct {
	BotToken      string `yaml:"BotToken"`
	OpenAIToken   string `yaml:"OpenAIToken"`
	ProxyHost     string `yaml:"ProxyHost"`
	ProxyUser     string `yaml:"ProxyUser"`
	ProxyPassword string `yaml:"ProxyPassword"`
}

func GetSecrets(path string) Secrets {
	secrets := Secrets{}
	yfile, ioErr := os.ReadFile(path)
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &secrets)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	return secrets
}
