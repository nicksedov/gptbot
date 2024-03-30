package logger

import (
	"io"
	"log"
	"os"
	"strings"

	"gptbot/pkg/settings"
)

func InitLog() {
	settings := settings.GetSettings()
	if strings.TrimSpace(settings.Logger.Filename) != "" {
		lumberjackLogger := &settings.Logger
		multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)
		log.SetFlags(log.LstdFlags)
		log.SetOutput(multiWriter)
	}
}
