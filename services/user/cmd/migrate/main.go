package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/LarsSeverson/charter/services/user/internal/bootstrap"
	"github.com/LarsSeverson/charter/services/user/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := run(logger); err != nil {
		logger.Error(
			"PostgreSQL migration failed",
			slog.Any("error", err),
		)

		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := config.LoadMigration()
	if err != nil {
		return err
	}

	return bootstrap.Migrate(ctx, cfg, logger)
}
