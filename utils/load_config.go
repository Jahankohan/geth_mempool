package utils

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func InitConfig() (Config, error) {
	config, err := LoadConfig(".")
	if err != nil {
		return Config{}, err
	}

	return config, nil
}