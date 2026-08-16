package bootstrap

import (
	"context"
	"log/slog"

	"github.com/LarsSeverson/charter/services/user/internal/config"
	"github.com/LarsSeverson/charter/services/user/internal/migration"
	"github.com/LarsSeverson/charter/services/user/migrations"
)

func Migrate(
	ctx context.Context,
	cfg config.MigrationConfig,
	logger *slog.Logger,
) (runErr error) {
	runner, err := migration.NewPostgres(
		migration.PostgresConfig{
			URL:             cfg.PostgresURL,
			MigrationsTable: cfg.MigrationsTable,
			StateTimeout:    cfg.StatementTimeout,
		},
		migrations.Files,
	)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := runner.Close()
		if closeErr == nil {
			return
		}
		if runErr != nil {
			return
		}

		runErr = closeErr
	}()

	logger.Info(
		"applying PostgreSQL migrations",
		slog.String("migration.table", cfg.MigrationsTable),
	)

	if err := runner.Up(ctx); err != nil {
		return err
	}

	version, err := runner.Version()
	if err != nil {
		return err
	}

	logger.Info(
		"PostgreSQL migrations applied",
		slog.Uint64("migration.version", uint64(version.Number)),
		slog.Bool("migration.dirty", version.Dirty),
	)

	return nil
}
