package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func New(cfg Config) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, errors.Join(ErrParseConfig, err)
	}

	if cfg.MaxConns > 0 {
		poolConfig.MaxConns = cfg.MaxConns
	}
	if cfg.MinConns > 0 {
		poolConfig.MinConns = cfg.MinConns
	}

	poolConfig.ConnConfig.RuntimeParams["application_name"] = applicationName

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.Join(ErrCreatePool, err)
	}

	db := &Database{
		pool: pool,
	}
	return db, nil
}

// Lifecycle

func (d *Database) Name() string {
	return "postgres"
}

func (d *Database) Start(ctx context.Context) error {
	err := d.pool.Ping(ctx)
	if err != nil {
		return errors.Join(ErrPingDatabase, err)
	}

	return nil
}

// Stop closes the pool. pool.Close waits for every checked-out
// connection to be released, so it is bounded by ctx.
func (d *Database) Stop(ctx context.Context) error {
	closed := make(chan struct{})

	go func() {
		d.pool.Close()
		close(closed)
	}()

	select {
	case <-closed:
		return nil
	case <-ctx.Done():
		err := errors.Join(ErrClosePool, ctx.Err())
		return err
	}
}

// DBTX

func (d *Database) Exec(
	ctx context.Context,
	sql string,
	arguments ...any,
) (pgconn.CommandTag, error) {
	tag, err := d.pool.Exec(ctx, sql, arguments...)
	return tag, err
}

func (d *Database) Query(
	ctx context.Context,
	sql string,
	arguments ...any,
) (pgx.Rows, error) {
	rows, err := d.pool.Query(ctx, sql, arguments...)
	return rows, err
}

var _ DBTX = (*Database)(nil)
