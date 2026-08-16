package bootstrap

import (
	"context"
	"log/slog"
	"net"

	"github.com/LarsSeverson/charter/services/user/internal/application/command"
	"github.com/LarsSeverson/charter/services/user/internal/config"
	"github.com/LarsSeverson/charter/services/user/internal/infrastructure/persistence/postgres"
	grpcserver "github.com/LarsSeverson/charter/services/user/internal/transport/grpc/server"
	grpcuser "github.com/LarsSeverson/charter/services/user/internal/transport/grpc/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	pool       *pgxpool.Pool
	grpcServer *grpcserver.Server
}

func New(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
) (*App, error) {
	pool, err := postgres.Open(ctx, postgres.Config{
		URL:      cfg.Postgres.URL,
		MaxConns: cfg.Postgres.MaxConns,
		MinConns: cfg.Postgres.MinConns,
	})
	if err != nil {
		return nil, err
	}

	users := postgres.NewUsers(pool)

	createUser := command.NewCreateUser(users)

	userService := grpcuser.NewService(createUser)

	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		pool.Close()
		return nil, err
	}

	grpcServer := grpcserver.New(
		listener,
		userService,
		logger,
		grpcserver.Config{
			ShutdownTimeout: cfg.GRPC.ShutdownTimeout,
		},
	)

	return &App{
		pool:       pool,
		grpcServer: grpcServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.grpcServer.Run(ctx)
	return nil
}

func (a *App) Close() {
	a.pool.Close()
}
