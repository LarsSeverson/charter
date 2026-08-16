package queries

import _ "embed"

var (
	//go:embed user/find_by_id.sql
	FindUserByID string

	//go:embed user/create.sql
	CreateUser string

	//go:embded user/update.sql
	UpdateUser string
)
