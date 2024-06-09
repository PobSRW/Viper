package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

const (
	configPath = "./config"
	configName = "config"
	configType = "yaml"
)

func LoadConfiguration() (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// READ CONFIG
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("could not read config file", "msg", err)
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		slog.Error("could not unmarshal", "msg", err)
		return nil, err
	}

	return config, nil
}
