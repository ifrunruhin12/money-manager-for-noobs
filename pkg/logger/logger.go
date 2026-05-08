package logger

import (
	"log/slog"
	"os"
	"strings"
)

// New creates a new structured JSON logger at the given log level.
// Accepts "debug", "info", "warn", "error" (case-insensitive); defaults to "info".
func New(level string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl})
	return slog.New(handler)
}
