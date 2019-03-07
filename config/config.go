package config

var Configuration *Config

type Config struct {
	Version  string `json:"version"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password int    `json:"password"`
	} `yaml:"database"`
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `yaml:"redis"`
	Domain struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `yaml:"domain"`
}

func SetConfig(config *Config){
	Configuration = config
}