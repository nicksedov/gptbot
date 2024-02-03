package settings

/*
 * Reading and initializing settings from file on service startup (default name 'gptbot-settings.yaml')
 * Including:
 *  - SQLite database settings
 *  - Telegram connection settings
 */
import (
	"log"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/gptbot/pkg/cli"
)

var settings Settings = Settings{}
var initialized bool = false

func GetSettings() *Settings {
	if !initialized {
		readSettingsFile()
		initialized = true
	}
	return &settings
}

func readSettingsFile() {
	if strings.TrimSpace(*cli.FlagConfig) != "" {
		yfile, ioErr := os.ReadFile(*cli.FlagConfig)
		if ioErr == nil {
			ymlErr := yaml.Unmarshal(yfile, &settings)
			if ymlErr != nil {
				log.Fatal(ymlErr)
			}
		}
	} else {
		log.Fatal("Settings file undefined")
	}
}