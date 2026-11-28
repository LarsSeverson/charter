package postgres

import "errors"

var (
	ErrParseConfig  = errors.New("parse PostgreSQL configuration")
	ErrCreatePool   = errors.New("create PostgreSQL connection pool")
	ErrPingDatabase = errors.New("ping PostgreSQL database")
	ErrClosePool    = errors.New("close PostgreSQL connection pool")
)
