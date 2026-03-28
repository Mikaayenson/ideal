package version_test

import (
	"strings"
	"testing"

	"github.com/stryker/ideal/internal/version"
)

func TestString_nonEmpty(t *testing.T) {
	t.Parallel()
	s := version.String()
	if s == "" {
		t.Fatal("empty version string")
	}
	if !strings.Contains(s, "commit") {
		t.Fatalf("expected commit in %q", s)
	}
}
