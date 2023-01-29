package config

import (
	"github.com/spf13/viper"
)

func InitDefaultConfig() (err error) {
	viper.SetConfigFile("main.json")
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	return
}
