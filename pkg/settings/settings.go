package settings

type Settings struct {
	
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`  
	} `yaml:"server"`
	
	DbConfig struct {
		Path string `yaml:"path"`
	} `yaml:"database"`

}
