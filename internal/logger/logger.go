package logger

import (
	"log/slog"
	"os"
	"strings"
)

// SetupLogger initializes a structured JSON logger with a configurable log level.
func SetupLogger(levelStr string) *slog.Logger {
	var level slog.Level

	switch strings.ToLower(levelStr) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	// Output structured JSON to stdout
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	// Set as global default logger
	slog.SetDefault(logger)

	return logger
}
