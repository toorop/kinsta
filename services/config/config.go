package config

import (
	"path"

	"github.com/spf13/viper"
)

func Init(home string) error {
	//println(home)
	viper.SetConfigFile(path.Join(home, "conf.yaml"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.Set("home", home)
	viper.WatchConfig()
	return nil
}
