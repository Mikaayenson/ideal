package logging

import (
	"log/slog"
	"os"
	"strings"
)

// New returns a slog.Logger configured for text or JSON output.
// level is parsed with [slog.Level.UnmarshalText] (debug, info, warn, error, case-insensitive); unknown values fall back to info.
func New(level string, jsonOut bool) *slog.Logger {
	lvl := slog.LevelInfo
	if s := strings.TrimSpace(level); s != "" {
		var parsed slog.Level
		if err := parsed.UnmarshalText([]byte(strings.ToLower(s))); err == nil {
			lvl = parsed
		}
	}
	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler
	if jsonOut {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	return slog.New(h)
}
