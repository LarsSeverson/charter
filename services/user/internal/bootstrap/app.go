package bootstrap

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/LarsSeverson/charter/internal/lifecycle"
	grpcadapter "github.com/LarsSeverson/charter/services/user/internal/adapter/inbound/grpc"
	grpcusers "github.com/LarsSeverson/charter/services/user/internal/adapter/inbound/grpc/users"
	"github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres"
	pgusers "github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres/users"
	appuser "github.com/LarsSeverson/charter/services/user/internal/app/user"
	userport "github.com/LarsSeverson/charter/services/user/internal/app/user/port"
	"github.com/LarsSeverson/charter/services/user/internal/bootstrap/config"
)

type App struct {
	lifecycle       *lifecycle.Lifecycle
	shutdownTimeout time.Duration
}

// New builds the whole service: config, logger, adapters, use cases.
// It performs no I/O. Connecting and listening happen in Run.
func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	lc := lifecycle.New()

	// Outbound adapters
	pgCfg := postgres.Config{
		URL:      cfg.Postgres.URL,
		MaxConns: cfg.Postgres.MaxConns,
		MinConns: cfg.Postgres.MinConns,
	}
	db, err := postgres.New(pgCfg)
	if err != nil {
		return nil, err
	}

	userRepo := pgusers.NewUsers(db)

	// Use cases
	userDeps := userport.Deps{
		Users: userRepo,
	}
	userUseCases := appuser.New(userDeps)

	// Inbound adapters
	usersHandler := grpcusers.NewHandler(userUseCases)

	server := grpcadapter.New(
		grpcadapter.Config{
			Address: cfg.GRPC.Address,
		},
		usersHandler,
		logger,
		lc.Fail,
	)

	// Start order: db first, server last. Stop order is the reverse.
	lc.Add(db)
	lc.Add(server)

	app := &App{
		lifecycle:       lc,
		shutdownTimeout: cfg.GRPC.ShutdownTimeout,
	}
	return app, nil
}

// Run starts every component, blocks until ctx is cancelled or a
// component fails, then stops everything in reverse order.
func (a *App) Run(ctx context.Context) error {
	err := a.lifecycle.Run(ctx, a.shutdownTimeout)
	return err
}
