package config_test

import (
	"testing"

	"github.com/stryker/ideal/internal/config"
)

func TestLoad_customValues(t *testing.T) {
	t.Setenv("IDEAL_USERNAME", "Ada")
	t.Setenv("IDEAL_LOG_LEVEL", "debug")
	t.Setenv("IDEAL_LOG_JSON", "true")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Username != "Ada" || cfg.LogLevel != "debug" || !cfg.LogJSON {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestLoad_invalidBool(t *testing.T) {
	t.Setenv("IDEAL_LOG_JSON", "not-a-bool")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for invalid IDEAL_LOG_JSON")
	}
}
