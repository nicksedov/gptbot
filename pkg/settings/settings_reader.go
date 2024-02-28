package settings

/*
 * Reading and initializing settings from file on service startup (default name 'settings.yaml')
 * Including:
 *  - PostgreSQL database settings
 *  - Telegram connection settings
 */
import (
	"log"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/gptbot/pkg/cli"
)

var settings *Settings

func GetSettings() *Settings {
	if settings == nil {
		readSettingsFile()
	}
	return settings
}

func readSettingsFile() {
	settings = &Settings{}
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