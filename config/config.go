package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	AppPort  string
	BasePath string
	Env      string
}

var AppConfig *Config

func LoadConfiguration() error {

	viper.SetDefault("APP_NAME", "go-chat-room")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("BASE_PATH", "test")
	viper.SetDefault("ENV", "production")

	err := readConfigFile("../your-path-file")

	AppConfig = &Config{
		AppName:  viper.GetString("APP_NAME"),
		AppPort:  viper.GetString("APP_PORT"),
		BasePath: viper.GetString("BASE_PATH"),
		Env:      viper.GetString("ENV"),
	}
	return err
}

func readConfigFile(path string) error {
	viper.AutomaticEnv()
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
