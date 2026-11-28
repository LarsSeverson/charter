package config

import (
	"errors"
	"os"
	"strings"

	"github.com/LarsSeverson/charter/internal/optional"
	"github.com/LarsSeverson/charter/internal/parse"
)

const (
	postgresURLEnvKey      = "USER_POSTGRES_URL"
	postgresMaxConnsEnvKey = "USER_POSTGRES_MAX_CONNS"
	postgresMinConnsEnvKey = "USER_POSTGRES_MIN_CONNS"

	defaultPostgresMaxConns int32 = 20
	defaultPostgresMinConns int32 = 2
)

type PostgresConfig struct {
	URL      string
	MaxConns int32
	MinConns int32
}

func loadPostgres() (PostgresConfig, error) {
	url := strings.TrimSpace(os.Getenv(postgresURLEnvKey))
	if url == "" {
		return PostgresConfig{}, ErrPostgresURLRequired
	}

	maxConns, err := parse.Int32(
		os.Getenv(postgresMaxConnsEnvKey),
		optional.Some(defaultPostgresMaxConns),
	)
	if err != nil {
		return PostgresConfig{}, invalidEnvironmentValue(
			postgresMaxConnsEnvKey,
			err,
		)
	}
	if maxConns <= 0 {
		return PostgresConfig{}, invalidEnvironmentValue(
			postgresMaxConnsEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	minConns, err := parse.Int32(
		os.Getenv(postgresMinConnsEnvKey),
		optional.Some(defaultPostgresMinConns),
	)
	if err != nil {
		return PostgresConfig{}, invalidEnvironmentValue(
			postgresMinConnsEnvKey,
			err,
		)
	}
	if minConns <= 0 {
		return PostgresConfig{}, invalidEnvironmentValue(
			postgresMinConnsEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	if minConns > maxConns {
		return PostgresConfig{}, ErrInvalidPostgresPoolSize
	}

	return PostgresConfig{
		URL:      url,
		MaxConns: maxConns,
		MinConns: minConns,
	}, nil
}
