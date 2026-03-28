package greetings_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/stryker/ideal/internal/greetings"
)

func TestHelloName(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello(ctx, name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	msg, err := greetings.Hello(ctx, "")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
