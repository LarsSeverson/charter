package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/LarsSeverson/charter/services/user/internal/bootstrap/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	a, err := app.New()
	if err != nil {
		slog.Error("user service failed to start", slog.Any("error", err))
		os.Exit(1)
	}

	err = a.Run(ctx)
	if err != nil {
		slog.Error("user service stopped", slog.Any("error", err))
		os.Exit(1)
	}
}
