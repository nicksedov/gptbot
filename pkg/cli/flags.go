package cli

import (
	"flag"
)

// Command line parameters
var (
	FlagConfig = flag.String("config", "gptbot-settings.yaml", "configuration YAML file")
	FlagBotToken       = flag.String("bot", "", "telegram bot token")
	FlagOpenAIToken    = flag.String("openai", "", "OpenAI token")
)
