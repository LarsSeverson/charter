package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/LarsSeverson/charter/services/user/internal/bootstrap/migrator"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	m, err := migrator.New()
	if err != nil {
		slog.Error("migration failed to start", slog.Any("error", err))
		os.Exit(1)
	}

	err = m.Run(ctx)
	if err != nil {
		slog.Error("migration failed", slog.Any("error", err))
		os.Exit(1)
	}
}
