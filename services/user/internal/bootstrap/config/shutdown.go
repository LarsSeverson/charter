package config

import (
	"errors"
	"os"
	"strings"
	"time"
)

const (
	shutdownTimeoutEnvKey = "USER_SHUTDOWN_TIMEOUT"

	defaultShutdownTimeout = 10 * time.Second
)

func LoadShutdownTimeout() (time.Duration, error) {
	value := strings.TrimSpace(os.Getenv(shutdownTimeoutEnvKey))
	if value == "" {
		return defaultShutdownTimeout, nil
	}

	timeout, err := time.ParseDuration(value)
	if err != nil {
		return 0, invalidEnvironmentValue(
			shutdownTimeoutEnvKey,
			err,
		)
	}

	if timeout <= 0 {
		return 0, invalidEnvironmentValue(
			shutdownTimeoutEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	return timeout, nil
}
