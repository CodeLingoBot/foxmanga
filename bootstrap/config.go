package bootstrap

import (
	"fmt"
	"github.com/foxmanga/config"
	"github.com/spf13/viper"
	"sync"
)

const (
	ConfigPathFoxEnv = "FOX_CONFIG"
)

var (
	loadOnce sync.Once
)

func initialize() {
	loadOnce.Do(func() {
		_config := new(config.Config)
		viper.Reset()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		//viper.AddConfigPath(os.Getenv(ConfigPathFoxEnv))
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error _config file: %s \n", err))
		}
		if err := viper.Unmarshal(_config); err != nil {
			panic(fmt.Errorf("Fatal error Unmarshal: %s \n", err))
		}
		config.SetConfig(_config)
	})
}
func init() {
	initialize()
}
