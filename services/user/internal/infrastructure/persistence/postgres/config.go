package postgres

type Config struct {
	URL      string
	MaxConns int32
	MinConns int32
}

const applicationName = "user-service"
