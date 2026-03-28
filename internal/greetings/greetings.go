package greetings

import (
	"context"
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(_ context.Context, name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hi, %s. Welcome!", name), nil
}
