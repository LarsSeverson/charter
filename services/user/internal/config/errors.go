package config

import "errors"

var (
	ErrPostgresURLRequired     = errors.New("USER_POSTGRES_URL is required")
	ErrInvalidEnvironmentValue = errors.New("invalid environment value")
	ErrInvalidPostgresPoolSize = errors.New("PostgreSQL minimum connections cannot exceed maximum connections")
)
