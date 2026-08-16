package postgres

import (
	"context"
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/application/port"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
	"github.com/LarsSeverson/charter/services/user/internal/infrastructure/persistence/postgres/queries"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	db DBTX
}

func NewUsers(db DBTX) *Users {
	return &Users{db: db}
}

func (r *Users) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	rows, err := r.db.Query(
		ctx,
		queries.FindUserByID,
		id.String(),
	)
	if err != nil {
		return nil, errors.Join(ErrFindUser, err)
	}

	row, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[userRow],
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, port.ErrUserNotFound
	}
	if err != nil {
		return nil, errors.Join(ErrFindUser, err)
	}

	value, err := row.domain()
	if err != nil {
		return nil, errors.Join(ErrReconstituteUser, err)
	}

	return value, nil
}

func (r *Users) Create(ctx context.Context, value *user.User) error {
	_, err := r.db.Exec(
		ctx,
		queries.CreateUser,
		value.ID().String(),
		value.Status().String(),
	)
	if err != nil {
		return errors.Join(ErrCreateUser, err)
	}

	return nil
}

func (r *Users) Update(ctx context.Context, value *user.User) error {
	result, err := r.db.Exec(
		ctx,
		queries.UpdateUser,
		value.ID().String(),
		value.Status().String(),
	)
	if err != nil {
		return errors.Join(ErrUpdateUser, err)
	}

	if result.RowsAffected() == 0 {
		return port.ErrUserNotFound
	}

	return nil
}

var _ port.UserRepository = (*Users)(nil)
