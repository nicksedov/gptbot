package cli

import (
	"flag"
)

// Command line parameters
var (
	FlagConfig      = flag.String("config", "settings.yaml", "configuration YAML file")
	FlagBotToken    = flag.String("bot", "", "telegram bot token")
	FlagOpenAIToken = flag.String("openai", "", "OpenAI token")
	Proxy           = flag.String("proxy", "", "HTTP Proxy host:port")
	ProxyUser       = flag.String("proxy.user", "", "HTTP Proxy user")
	ProxyPassword   = flag.String("proxy.password", "", "HTTP Proxy password")
)
