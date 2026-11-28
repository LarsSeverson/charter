package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres/migrate"
)

const (
	postgresMigrationURLEnvKey               = "USER_POSTGRES_MIGRATION_URL"
	postgresMigrationTableEnvKey             = "USER_POSTGRES_MIGRATIONS_TABLE"
	postgresMigrationStatementTimeoutEnvKey  = "USER_POSTGRES_MIGRATION_STATEMENT_TIMEOUT"
	defaultPostgresMigrationsTable           = "schema_migrations"
	defaultPostgresMigrationStatementTimeout = 5 * time.Minute
)

func LoadMigration() (migrate.PostgresConfig, error) {
	postgresURL := strings.TrimSpace(os.Getenv(postgresMigrationURLEnvKey))
	if postgresURL == "" {
		return migrate.PostgresConfig{}, fmt.Errorf("%s is required", postgresMigrationURLEnvKey)
	}

	migrationsTable := strings.TrimSpace(os.Getenv(postgresMigrationTableEnvKey))
	if migrationsTable == "" {
		migrationsTable = defaultPostgresMigrationsTable
	}

	statementTime, err := loadMigrationStatementTimeout()
	if err != nil {
		return migrate.PostgresConfig{}, err
	}

	cfg := migrate.PostgresConfig{
		URL:              postgresURL,
		MigrationsTable:  migrationsTable,
		StatementTimeout: statementTime,
	}
	return cfg, nil
}

func loadMigrationStatementTimeout() (time.Duration, error) {
	value := strings.TrimSpace(os.Getenv(postgresMigrationStatementTimeoutEnvKey))
	if value == "" {
		return defaultPostgresMigrationStatementTimeout, nil
	}

	timeout, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf(
			"parse %s: %w",
			postgresMigrationStatementTimeoutEnvKey,
			err,
		)
	}

	if timeout < time.Millisecond {
		return 0, fmt.Errorf(
			"%s must be at least one millisecond",
			postgresMigrationStatementTimeoutEnvKey,
		)
	}

	return timeout, nil
}
