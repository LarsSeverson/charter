package app

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
	logger          *slog.Logger
}

func New() (*App, error) {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logHandler).With(slog.String("service.name", config.ServiceName))
	slog.SetDefault(logger)

	shutdownTimeout, err := config.LoadShutdownTimeout()
	if err != nil {
		return nil, err
	}

	lc := lifecycle.New()

	// Outbound adapters
	pgCfg, err := config.LoadPostgres()
	if err != nil {
		return nil, err
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

	grpcCfg, err := config.LoadGRPC()
	if err != nil {
		return nil, err
	}

	server := grpcadapter.New(grpcCfg, usersHandler, lc.Fail)

	lc.Add(db)
	lc.Add(server)

	app := &App{
		lifecycle:       lc,
		shutdownTimeout: shutdownTimeout,
		logger:          logger,
	}
	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("user service starting")

	err := a.lifecycle.Run(ctx, a.shutdownTimeout)
	if err != nil {
		return err
	}

	a.logger.Info("user service stopped")
	return nil
}
