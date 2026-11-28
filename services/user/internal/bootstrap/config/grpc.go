package config

import (
	"errors"
	"net"
	"os"
	"strconv"
	"strings"

	grpcadapter "github.com/LarsSeverson/charter/services/user/internal/adapter/inbound/grpc"
)

const (
	grpcAddressEnvKey = "USER_GRPC_ADDRESS"

	defaultGRPCAddress = ":50051"
)

func LoadGRPC() (grpcadapter.Config, error) {
	address := strings.TrimSpace(os.Getenv(grpcAddressEnvKey))
	if address == "" {
		address = defaultGRPCAddress
	}

	err := validateGRPCAddress(address)
	if err != nil {
		return grpcadapter.Config{}, invalidEnvironmentValue(
			grpcAddressEnvKey,
			err,
		)
	}

	cfg := grpcadapter.Config{
		Address: address,
	}
	return cfg, nil
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
