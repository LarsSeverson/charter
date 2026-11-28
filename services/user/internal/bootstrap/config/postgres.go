package config

import (
	"errors"
	"os"
	"strings"

	"github.com/LarsSeverson/charter/internal/optional"
	"github.com/LarsSeverson/charter/internal/parse"
	"github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres"
)

const (
	postgresURLEnvKey      = "USER_POSTGRES_URL"
	postgresMaxConnsEnvKey = "USER_POSTGRES_MAX_CONNS"
	postgresMinConnsEnvKey = "USER_POSTGRES_MIN_CONNS"

	defaultPostgresMaxConns int32 = 20
	defaultPostgresMinConns int32 = 2
)

func LoadPostgres() (postgres.Config, error) {
	url := strings.TrimSpace(os.Getenv(postgresURLEnvKey))
	if url == "" {
		return postgres.Config{}, ErrPostgresURLRequired
	}

	maxConns, err := parse.Int32(
		os.Getenv(postgresMaxConnsEnvKey),
		optional.Some(defaultPostgresMaxConns),
	)
	if err != nil {
		return postgres.Config{}, invalidEnvironmentValue(
			postgresMaxConnsEnvKey,
			err,
		)
	}
	if maxConns <= 0 {
		return postgres.Config{}, invalidEnvironmentValue(
			postgresMaxConnsEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	minConns, err := parse.Int32(
		os.Getenv(postgresMinConnsEnvKey),
		optional.Some(defaultPostgresMinConns),
	)
	if err != nil {
		return postgres.Config{}, invalidEnvironmentValue(
			postgresMinConnsEnvKey,
			err,
		)
	}
	if minConns <= 0 {
		return postgres.Config{}, invalidEnvironmentValue(
			postgresMinConnsEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	if minConns > maxConns {
		return postgres.Config{}, ErrInvalidPostgresPoolSize
	}

	cfg := postgres.Config{
		URL:             url,
		MaxConns:        maxConns,
		MinConns:        minConns,
		ApplicationName: ServiceName,
	}
	return cfg, nil
}
