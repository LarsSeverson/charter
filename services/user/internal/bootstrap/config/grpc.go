package config

import (
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	grpcAddressEnvKey         = "USER_GRPC_ADDRESS"
	grpcShutdownTimeoutEnvKey = "USER_GRPC_SHUTDOWN_TIMEOUT"

	defaultGRPCAddress         = ":50051"
	defaultGRPCShutdownTimeout = 10 * time.Second
)

type GRPCConfig struct {
	Address         string
	ShutdownTimeout time.Duration
}

func loadGRPC() (GRPCConfig, error) {
	address := strings.TrimSpace(os.Getenv(grpcAddressEnvKey))
	if address == "" {
		address = defaultGRPCAddress
	}

	if err := validateGRPCAddress(address); err != nil {
		return GRPCConfig{}, invalidEnvironmentValue(
			grpcAddressEnvKey,
			err,
		)
	}

	shutdownTimeout := defaultGRPCShutdownTimeout
	value := strings.TrimSpace(os.Getenv(grpcShutdownTimeoutEnvKey))
	if value != "" {
		parsed, err := time.ParseDuration(value)
		if err != nil {
			return GRPCConfig{}, invalidEnvironmentValue(
				grpcShutdownTimeoutEnvKey,
				err,
			)
		}

		shutdownTimeout = parsed
	}

	if shutdownTimeout <= 0 {
		return GRPCConfig{}, invalidEnvironmentValue(
			grpcShutdownTimeoutEnvKey,
			errors.New("must be greater than zero"),
		)
	}

	return GRPCConfig{
		Address:         address,
		ShutdownTimeout: shutdownTimeout,
	}, nil
}

func validateGRPCAddress(address string) error {
	_, port, err := net.SplitHostPort(address)
	if err != nil {
		return errors.New("must use host:port format")
	}

	portNumber, err := strconv.ParseUint(port, 10, 16)
	if err != nil || portNumber == 0 {
		return errors.New("port must be between 1 and 65535")
	}

	return nil
}
