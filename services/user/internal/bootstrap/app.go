package bootstrap

import (
	"context"

	"github.com/LarsSeverson/charter/services/user/internal/application/command"
	"github.com/LarsSeverson/charter/services/user/internal/config"
	"github.com/LarsSeverson/charter/services/user/internal/infrastructure/persistence/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	pool *pgxpool.Pool
	// server
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
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

	// handler := transport.NewHandler(createUser)
	// server, err := transport.NewServer(handler)
	// if err != nil {
	//     pool.Close()
	//     return nil, errors.Join(ErrCreateServer, err)
	// }

	// Remove once createUser is passed to the handler.
	_ = createUser

	return &App{
		pool: pool,
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