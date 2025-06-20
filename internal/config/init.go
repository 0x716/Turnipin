package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config error: %v", err)
	}

	var cfg AppConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("load config error: %v", err)
	}

	GlobalConfig = &cfg

	return nil
}
