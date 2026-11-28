package migrator

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres/migrate"
	"github.com/LarsSeverson/charter/services/user/internal/bootstrap/config"
	"github.com/LarsSeverson/charter/services/user/migrations"
)

type Migrator struct {
	cfg    migrate.PostgresConfig
	logger *slog.Logger
}

func New() (*Migrator, error) {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logHandler).With(slog.String("service.name", config.ServiceName))
	slog.SetDefault(logger)

	cfg, err := config.LoadMigration()
	if err != nil {
		return nil, err
	}

	migrator := &Migrator{
		cfg:    cfg,
		logger: logger,
	}
	return migrator, nil
}

func (m *Migrator) Run(ctx context.Context) (err error) {
	runner, err := migrate.NewPostgres(m.cfg, migrations.Files)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := runner.Close()
		err = errors.Join(err, closeErr)
	}()

	m.logger.Info(
		"applying PostgreSQL migrations",
		slog.String("migration.table", m.cfg.MigrationsTable),
	)

	err = runner.Up(ctx)
	if err != nil {
		return err
	}

	version, err := runner.Version()
	if err != nil {
		return err
	}

	m.logger.Info(
		"PostgreSQL migrations applied",
		slog.Uint64("migration.version", uint64(version.Number)),
		slog.Bool("migration.dirty", version.Dirty),
	)

	return nil
}
