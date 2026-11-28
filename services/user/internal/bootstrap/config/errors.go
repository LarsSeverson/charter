package config

import (
	"errors"
	"fmt"
)

var (
	ErrPostgresURLRequired     = errors.New("USER_POSTGRES_URL is required")
	ErrInvalidEnvironmentValue = errors.New("invalid environment value")
	ErrInvalidPostgresPoolSize = errors.New("PostgreSQL minimum connections cannot exceed maximum connections")
)

func invalidEnvironmentValue(key string, err error) error {
	return errors.Join(
		ErrInvalidEnvironmentValue,
		fmt.Errorf("%s: %w", key, err),
	)
}
