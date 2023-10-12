package config

import (
	"github.com/spf13/viper"
)

var Vauban *Config

// LoadConfig initializes the configuration.
func LoadConfig() error {

	config, err := InitConfig()
	if err != nil {
		return err
	}

	Vauban = config
	return err
}

// InitConfig reads and parses the configuration file.
func InitConfig() (*Config, error) {
	// Specify the configuration file and paths
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("toml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
