package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Open(ctx context.Context, input Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(input.URL)
	if err != nil {
		return nil, errors.Join(ErrParseConfig, err)
	}

	if input.MaxConns > 0 {
		poolConfig.MaxConns = input.MaxConns
	}

	if input.MinConns > 0 {
		poolConfig.MinConns = input.MinConns
	}

	poolConfig.ConnConfig.RuntimeParams["application_name"] = applicationName

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, errors.Join(ErrCreatePool, err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, errors.Join(ErrPingDatabase, err)
	}

	return pool, nil
}
