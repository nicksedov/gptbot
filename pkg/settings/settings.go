package settings

import "gopkg.in/natefinch/lumberjack.v2"

type Settings struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`

	Proxy struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"proxy"`

	DbConfig struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		DbName   string `yaml:"db_name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`

	OpenAI struct {
		APIToken string `yaml:"api_token"`
		Model    string `yaml:"model"`
		Completions struct {
			Context string `yaml:"context"`
		} `yaml:"completions"`
	} `yaml:"openai"`

	GigaChat struct {
		ClientID          string `yaml:"client_id"`
		ClientSecret      string `yaml:"client_secret"`
		Model             string `yaml:"model"`
		Temperature       float64  `yaml:"temperature"`
		TopP              float64  `yaml:"top_p"`
		MaxTokens         int64    `yaml:"max_tokens"`
		RepetitionPenalty float64  `yaml:"repetition_penalty"`
		Completions struct {
			Context string `yaml:"context"`
		} `yaml:"completions"`
	}  `yaml:"gigachat"`

	Telegram struct {
		BotToken    string `yaml:"bot_token"`
		ServiceChat int64  `yaml:"service_chat"`
	} `yaml:"telegram"`

	Logger lumberjack.Logger  `yaml:"logger"`
}
