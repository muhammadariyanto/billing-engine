package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	InternalApiKey string `mapstructure:"INTERNAL_API_KEY"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	viper.Reset()

	return &config, nil
}
