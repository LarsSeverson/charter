package bootstrap

import (
	"context"
	"errors"
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
		return nil, ErrOpenDatabase
	}

	users := postgres.NewUsers(pool)

	createUser := command.NewCreateUser(users)

	userService := grpcuser.NewService(createUser)

	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		pool.Close()
		return nil, errors.Join(ErrCreateListener, err)
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
	// return a.server.Run(ctx)
	return nil
}

func (a *App) Close() {
	a.pool.Close()
}

// An option:
// package bootstrap

// import (
// 	"context"
// 	"errors"
// 	"net"
// 	"time"

// 	transportgrpc "github.com/LarsSeverson/charter/services/user/internal/transport/grpc"
// 	"github.com/LarsSeverson/charter/services/user/internal/application/command"
// 	"github.com/LarsSeverson/charter/services/user/internal/config"
// 	"github.com/LarsSeverson/charter/services/user/internal/infrastructure/persistence/postgres"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type App struct {
// 	pool   *pgxpool.Pool
// 	server *transportgrpc.Server
// }

// func New(ctx context.Context, cfg config.Config) (*App, error) {
// 	// Infrastructure adapter.
// 	pool, err := postgres.Open(ctx, postgres.Config{
// 		URL:      cfg.Postgres.URL,
// 		MaxConns: cfg.Postgres.MaxConns,
// 		MinConns: cfg.Postgres.MinConns,
// 	})
// 	if err != nil {
// 		return nil, errors.Join(ErrOpenDatabase, err)
// 	}

// 	users := postgres.NewUsers(pool)

// 	// Application use cases.
// 	createUser := command.NewCreateUser(users)

// 	// Inbound gRPC adapter.
// 	userService := transportgrpc.NewUserServiceHandler(createUser)

// 	// Network resource.
// 	listener, err := net.Listen("tcp", cfg.GRPC.Address)
// 	if err != nil {
// 		pool.Close()
// 		return nil, errors.Join(ErrCreateServer, err)
// 	}

// 	// gRPC runtime.
// 	server := transportgrpc.NewServer(
// 		listener,
// 		userService,
// 		10*time.Second,
// 	)

// 	return &App{
// 		pool:   pool,
// 		server: server,
// 	}, nil
// }

// func (a *App) Run(ctx context.Context) error {
// 	return a.server.Run(ctx)
// }

// func (a *App) Close() {
// 	a.pool.Close()
// }
