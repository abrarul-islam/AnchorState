package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abrarul-islam/AnchorState/internal/version"
	"github.com/abrarul-islam/AnchorState/pkg/logger"
)

func main() {
	// Initialize structured JSON logging (default level: INFO)
	log := logger.SetupLogger("info")

	log.Info("starting runtime trust engine",
		slog.String("app", version.Name),
		slog.String("version", version.Version),
	)

	// Context listening for OS interrupt signals (Ctrl+C / SIGTERM)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx, log); err != nil {
		log.Error("fatal application error", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("AnchorState shut down gracefully")
}

func run(ctx context.Context, log *slog.Logger) error {
	log.Info("initializing secret drift monitoring engine...")

	// Simulate background watcher startup
	select {
	case <-time.After(200 * time.Millisecond):
		log.Info("runtime watcher active", slog.String("target", "kubernetes.io/secrets"))
	case <-ctx.Done():
		return ctx.Err()
	}

	// Block until SIGINT / SIGTERM is received
	<-ctx.Done()
	log.Info("shutdown signal received, flushing event pipeline...")

	return nil
}
