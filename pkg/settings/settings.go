package settings

type Settings struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`

	DbConfig struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		DbName   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SslMode  string `yaml:"sslmode"` 
	} `yaml:"database"`
}
