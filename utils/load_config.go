package utils

import (
	"log"

	"github.com/spf13/viper"
)

var variable *Config

func InitConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&variable)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() (*Config) {
	return variable
}