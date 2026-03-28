package logging_test

import (
	"testing"

	"github.com/stryker/ideal/internal/logging"
)

func TestNew_nonNil(t *testing.T) {
	for _, level := range []string{"", "debug", "info", "warn", "error", "not-a-real-level"} {
		if logging.New(level, false) == nil {
			t.Fatalf("New(%q, false): nil logger", level)
		}
		if logging.New(level, true) == nil {
			t.Fatalf("New(%q, true): nil logger", level)
		}
	}
}
