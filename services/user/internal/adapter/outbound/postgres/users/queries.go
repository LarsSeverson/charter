package users

import _ "embed"

var (
	//go:embed queries/find_by_id.sql
	FindUserByID string

	//go:embed queries/create.sql
	CreateUser string

	//go:embded queries/update.sql
	UpdateUser string
)
