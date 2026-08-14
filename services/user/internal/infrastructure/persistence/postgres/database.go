package postgres

import "database/sql"

type PostgresDB struct {
	connection string
}

// Make method of struct
func Open(connection string) (*sql.DB, error) {
	// Open driver
	// Configure connection pool
	// Verify connectivity
}
