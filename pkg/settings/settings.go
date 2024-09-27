package settings

import "gopkg.in/natefinch/lumberjack.v2"

type Settings struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`

	DbConfig struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		DbName   string `yaml:"db_name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`

	GigaChat struct {
		Auth struct {
			ClientID          string `yaml:"client_id"`
			ClientSecret      string `yaml:"client_secret"`
		} `yaml:"auth"`
		TLS               []string `yaml:"tls"`
		Model             string   `yaml:"model"`
	}  `yaml:"gigachat"`
	
	LocalAI struct {
		URL               string  `yaml:"url"`
		Model             string   `yaml:"model"`
	}  `yaml:"localai"`

	LLMConfig struct {
		Temperature       float64  `yaml:"temperature"`
		TopP              float64  `yaml:"top_p"`
		MaxTokens         int64    `yaml:"max_tokens"`
		RepetitionPenalty float64  `yaml:"repetition_penalty"`
		Completions struct {
			Context string `yaml:"context"`
		} `yaml:"completions"`
	} `yaml:"llm_config"`

	Fallback bool `yaml:"fallback"`

	Telegram struct {
		BotToken    string `yaml:"bot_token"`
		ServiceChat int64  `yaml:"service_chat"`
	} `yaml:"telegram"`

	Logger lumberjack.Logger  `yaml:"logger"`
}
