package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ifrunruhin12/money-manager/internal/config"
	"github.com/ifrunruhin12/money-manager/internal/db"
	"github.com/ifrunruhin12/money-manager/pkg/logger"
)

func main() {
	// Step 1: Load configuration from environment variables.
	// Exits with a descriptive error if required vars (DATABASE_URL, JWT_SECRET) are missing.
	cfg, err := config.Load()
	if err != nil {
		// Logger not yet available; fall back to slog default for this one fatal message.
		slog.Error("failed to load configuration", "err", err)
		os.Exit(1)
	}

	// Step 2: Initialize structured logger.
	log := logger.New(cfg.LogLevel)
	log.Info("configuration loaded", "port", cfg.Port, "log_level", cfg.LogLevel)

	// Step 3: Connect to PostgreSQL.
	ctx := context.Background()
	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	defer pool.Close()
	log.Info("database connection established")

	// Step 4: Run pending migrations.
	// The migrations directory path is relative to the working directory (project root).
	if err := db.RunMigrations(cfg.DatabaseURL, "migrations", log); err != nil {
		log.Error("failed to run migrations", "err", err)
		os.Exit(1)
	}

	// Step 5: Startup complete.
	log.Info("ready")
}
