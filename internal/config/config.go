package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config holds application settings loaded from the environment (12-factor).
type Config struct {
	Username string `env:"IDEAL_USERNAME" envDefault:"guest"`
	// LogLevel: slog level name (debug, info, warn, error); invalid values are treated as info in logging.New.
	LogLevel string `env:"IDEAL_LOG_LEVEL" envDefault:"info"`
	LogJSON  bool   `env:"IDEAL_LOG_JSON" envDefault:"false"`
}

// Load parses environment variables into Config.
func Load() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}
	return cfg, nil
}
