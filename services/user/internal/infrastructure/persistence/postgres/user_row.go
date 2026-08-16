package postgres

import "github.com/LarsSeverson/charter/services/user/internal/domain/user"

type userRow struct {
	ID     string `db:"id"`
	Status string `db:"status"`
}

func (row userRow) domain() (*user.User, error) {
	id, err := user.ParseID(row.ID)
	if err != nil {
		return nil, err
	}

	return user.Reconstitute(id, user.Status(row.Status))
}
