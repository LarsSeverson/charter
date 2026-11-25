package migration

import "errors"

var (
	ErrMigrationURL              = errors.New("migration URL is required")
	ErrMigrationStatementTimeout = errors.New("migration statement timeout must exist and be positive")
	ErrMigrationSource           = errors.New("create migration source")
	ErrMigrationPostgresConfig   = errors.New("postgresql migration configuration")
	ErrMigrationPostgresPing     = errors.New("ping postgresql migration database")
	ErrMigrationPostgresDriver   = errors.New("create postgresql migration driver")
)
