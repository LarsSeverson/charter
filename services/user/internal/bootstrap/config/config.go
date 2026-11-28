package config

type Config struct {
	Postgres PostgresConfig
	GRPC     GRPCConfig
}

func Load() (Config, error) {
	postgresConfig, err := loadPostgres()
	if err != nil {
		return Config{}, err
	}

	grpcConfig, err := loadGRPC()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Postgres: postgresConfig,
		GRPC:     grpcConfig,
	}, nil
}
