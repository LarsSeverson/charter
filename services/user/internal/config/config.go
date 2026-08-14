package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/LarsSeverson/charter/internal/optional"
	"github.com/LarsSeverson/charter/internal/parse"
)

const (
	postgresURLEnvKey      = "USER_POSTGRES_URL"
	postgresMaxConnsEnvKey = "USER_POSTGRES_MAX_CONNS"
	postgresMinConnsEnvKey = "USER_POSTGRES_MIN_CONNS"

	defaultMaxConns int32 = 20
	defaultMinConns int32 = 2
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	URL      string
	MaxConns int32
	MinConns int32
}

func Load() (Config, error) {
	url := strings.TrimSpace(os.Getenv(postgresURLEnvKey))
	if url == "" {
		return Config{}, ErrPostgresURLRequired
	}

	maxConns, err := parse.Int32(
		os.Getenv(postgresMaxConnsEnvKey),
		optional.Some(defaultMaxConns),
	)
	if err != nil {
		return Config{}, errors.Join(
			ErrInvalidEnvironmentValue,
			fmt.Errorf("%s: %w", postgresMaxConnsEnvKey, err),
		)
	}
	if maxConns <= 0 {
		return Config{}, errors.Join(
			ErrInvalidEnvironmentValue,
			fmt.Errorf("%s must be greater than zero", postgresMaxConnsEnvKey),
		)
	}

	minConns, err := parse.Int32(
		os.Getenv(postgresMinConnsEnvKey),
		optional.Some(defaultMinConns),
	)
	if err != nil {
		return Config{}, errors.Join(
			ErrInvalidEnvironmentValue,
			fmt.Errorf("%s: %w", postgresMinConnsEnvKey, err),
		)
	}
	if minConns <= 0 {
		return Config{}, errors.Join(
			ErrInvalidEnvironmentValue,
			fmt.Errorf("%s must be greater than zero", postgresMinConnsEnvKey),
		)
	}

	if minConns > maxConns {
		return Config{}, ErrInvalidPostgresPoolSize
	}

	postgresConfig := PostgresConfig{
		URL:      url,
		MaxConns: maxConns,
		MinConns: minConns,
	}

	return Config{
		Postgres: postgresConfig,
	}, nil
}
