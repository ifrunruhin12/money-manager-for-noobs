package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(databaseURL, migrationsDir string, logger *slog.Logger) error {
	sourceURL := "file://" + migrationsDir

	logger.Info("migrate: applying migrations", "dir", migrationsDir)
	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("migrate: init: %w", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			logger.Error("migrate: source close error", "err", srcErr)
		}
		if dbErr != nil {
			logger.Error("migrate: db close error", "err", dbErr)
		}
	}()

	m.Log = &migrateLogger{logger: logger}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("migrate: no new migrations to apply")
			return nil
		}
		return fmt.Errorf("migrate: up: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("migrate: version check: %w", err)
	}
	if dirty {
		return fmt.Errorf("migrate: database is in a dirty state at version %d; manual intervention required", version)
	}

	logger.Info("migrate: all migrations applied", "version", version)
	return nil
}

type migrateLogger struct {
	logger *slog.Logger
}

func (l *migrateLogger) Printf(format string, v ...any) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func (l *migrateLogger) Verbose() bool {
	return l.logger.Enabled(context.TODO(), slog.LevelDebug)
}
