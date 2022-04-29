package configs

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Settings IdealConfig
}

type IdealConfig struct {
	Username string
	Password string
}

// ReadConfig reads and parses config.yml
func ReadConfig() *Config {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config")
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("failed to decode config")
	}

	return &config

}
