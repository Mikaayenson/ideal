package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/stryker/ideal/internal/config"
	"github.com/stryker/ideal/internal/greetings"
	"github.com/stryker/ideal/internal/logging"
	"github.com/stryker/ideal/internal/version"
)

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Println(version.String())
		return
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ideal: %v\n", err)
		os.Exit(1)
	}
	slog.SetDefault(logging.New(cfg.LogLevel, cfg.LogJSON))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	msg, err := greetings.Hello(ctx, cfg.Username)
	if err != nil {
		slog.Error("greeting", "err", err)
		os.Exit(1)
	}
	fmt.Println(msg)
}
