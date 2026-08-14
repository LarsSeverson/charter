package postgres

import "errors"

var (
	// Database connection errors.
	ErrParseConfig  = errors.New("parse PostgreSQL configuration")
	ErrCreatePool   = errors.New("create PostgreSQL connection pool")
	ErrPingDatabase = errors.New("ping PostgreSQL database")

	// User persistence errors.
	ErrFindUser         = errors.New("find user")
	ErrCreateUser       = errors.New("create user")
	ErrUpdateUser       = errors.New("update user")
	ErrParseUserID      = errors.New("parse persisted user ID")
	ErrReconstituteUser = errors.New("reconstitute persisted user")
)
