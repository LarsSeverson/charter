package migration

import (
	"context"
	"fmt"
	"io/fs"
	"net/url"
	"strconv"
	"time"

	migratelib "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const defaultMigrationsTable = "schema_migrations"

type SchemaVersion struct {
	Number uint
	Dirty  bool
}

type PostgresConfig struct {
	URL             string
	MigrationsTable string
	StateTimeout    time.Duration
}

type Postgres struct {
	engine *migratelib.Migrate
}

func NewPostgres(
	cfg PostgresConfig,
	files fs.FS,
) (*Postgres, error) {
	if cfg.URL == "" {
		return nil, fmt.Errorf("migration URL is required")
	}

	if cfg.StateTimeout <= time.Millisecond {
		return nil, fmt.Errorf("migration statement timeout must be at least one millisecond")
	}

	if files == nil {
		return nil, fmt.Errorf("migration filesystem is required")
	}

	if cfg.MigrationsTable == "" {
		cfg.MigrationsTable = defaultMigrationsTable
	}

	databaseURL, err := postgresMigrationURL(cfg)
	if err != nil {
		return nil, fmt.Errorf("configure PostgreSQL migrations: %w", err)
	}

	sourceDriver, err := iofs.New(files, ".")
	if err != nil {
		return nil, fmt.Errorf("create migration source: %w", err)
	}

	engine, err := migratelib.NewWithSourceInstance(
		"iofs",
		sourceDriver,
		databaseURL,
	)
	if err != nil {
		_ = sourceDriver.Close()
		return nil, fmt.Errorf("create PostgreSQL migration engine: %w", err)
	}

	return &Postgres{
		engine: engine,
	}, nil
}

func (p *Postgres) Up(ctx context.Context) error {
	stop := context.AfterFunc(ctx, func() {
		select {
		case p.engine.GracefulStop <- true:
		default:
		}
	})
	defer stop()

	err := p.engine.Up()
	if err == migratelib.ErrNoChange {
		return nil
	}
	if err != nil {
		return fmt.Errorf("apply PostgreSQL migrations error: %w", err)
	}

	if ctxErr := ctx.Err(); ctxErr != nil {
		return fmt.Errorf("apply PostgreSQL migrations interrupted: %w", ctxErr)
	}

	return nil
}

func (p *Postgres) Version() (SchemaVersion, error) {
	number, dirty, err := p.engine.Version()
	if err != nil {
		return SchemaVersion{}, fmt.Errorf("read PostgreSQL migration version: %w", err)
	}
	return SchemaVersion{Number: number, Dirty: dirty}, nil
}

func (p *Postgres) Close() error {
	if p == nil || p.engine == nil {
		return nil
	}

	sourceErr, databaseErr := p.engine.Close()
	switch {
	case sourceErr != nil && databaseErr != nil:
		return fmt.Errorf(
			"close migration source: %w; close migration database: %w",
			sourceErr,
			databaseErr,
		)

	case sourceErr != nil:
		return fmt.Errorf(
			"close migration source: %w",
			sourceErr,
		)

	case databaseErr != nil:
		return fmt.Errorf(
			"close migration database: %w",
			databaseErr,
		)

	default:
		return nil
	}
}

func postgresMigrationURL(cfg PostgresConfig) (string, error) {
	value, err := url.Parse(cfg.URL)
	if err != nil {
		return "", fmt.Errorf("parse database URL: %w", err)
	}

	switch value.Scheme {
	case "postgres", "postgresql", "pgx5":
		value.Scheme = "pgx5"
	default:
		return "", fmt.Errorf(
			"unsupported PostgreSQL URL scheme %q",
			value.Scheme,
		)
	}

	query := value.Query()
	query.Set("x-migrations-table", cfg.MigrationsTable)
	query.Set("x-statement-timeout", strconv.FormatInt(cfg.StateTimeout.Milliseconds(), 10))
	query.Set("x-multi-statement", "true")

	value.RawQuery = query.Encode()

	return value.String(), nil
}
