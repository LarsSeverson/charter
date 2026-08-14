package postgres

type Users struct {
	db PostgresDB
}

// Maybe make db pointer
func NewUsers(db PostgresDB) *Users {
	return &Users{
		db: db,
	}
}
