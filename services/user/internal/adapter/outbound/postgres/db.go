package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const componentName = "postgres"

type Database struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
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

	poolConfig.ConnConfig.RuntimeParams["application_name"] = cfg.ApplicationName

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.Join(ErrCreatePool, err)
	}

	logger := slog.Default().With(slog.String("component", componentName))

	db := &Database{
		pool:   pool,
		logger: logger,
	}
	return db, nil
}

func (d *Database) Name() string {
	return componentName
}

func (d *Database) Start(ctx context.Context) error {
	err := d.pool.Ping(ctx)
	if err != nil {
		return errors.Join(ErrPingDatabase, err)
	}

	connConfig := d.pool.Config().ConnConfig

	d.logger.InfoContext(
		ctx,
		"PostgreSQL connected",
		slog.String("db.host", connConfig.Host),
		slog.String("db.name", connConfig.Database),
	)

	return nil
}

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
