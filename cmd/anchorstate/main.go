package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/abrarul-islam/AnchorState/internal/config"
	"github.com/abrarul-islam/AnchorState/internal/logger"
	"github.com/abrarul-islam/AnchorState/internal/version"
)

func main() {
	// 1. Load application configuration from environment
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize logger with configured log level
	log := logger.SetupLogger(cfg.LogLevel)

	// 3. Set up OS signal listening for graceful shutdown (Ctrl+C / SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 4. Log application startup metadata
	log.Info("starting application",
		"app", version.Name,
		"version", version.Version,
		"log_level", cfg.LogLevel,
		"target_namespace", cfg.Namespace,
	)

	// 5. Wait for shutdown signal
	<-ctx.Done()
	log.Info("shutdown signal received, exiting gracefully...")
}
