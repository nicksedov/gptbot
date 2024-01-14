package settings

type Settings struct {
	
	DbConfig struct {
		Path string `yaml:"path"`
	} `yaml:"database"`

	Telegram struct {
		Chats []struct {
			Alias  string `yaml:"alias"`
			ChatId int64  `yaml:"chatid"`
		} `yaml:"chats"` 
	} `yaml:"telegram"`
}
