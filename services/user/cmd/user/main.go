package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/LarsSeverson/charter/services/user/internal/bootstrap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, err := bootstrap.New()
	if err != nil {
		slog.Error("user service failed to start", slog.Any("error", err))
		os.Exit(1)
	}

	err = app.Run(ctx)
	if err != nil {
		slog.Error("user service stopped", slog.Any("error", err))
		os.Exit(1)
	}
}
