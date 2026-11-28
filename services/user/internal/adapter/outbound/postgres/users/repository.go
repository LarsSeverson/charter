package users

import (
	"context"
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/adapter/outbound/postgres"
	userport "github.com/LarsSeverson/charter/services/user/internal/app/user/port"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	db postgres.DBTX
}

func NewUsers(db postgres.DBTX) *Users {
	return &Users{db: db}
}

func (r *Users) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	rows, err := r.db.Query(
		ctx,
		FindUserByID,
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
		return nil, userport.ErrUserNotFound
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
		CreateUser,
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
		UpdateUser,
		value.ID().String(),
		value.Status().String(),
	)
	if err != nil {
		return errors.Join(ErrUpdateUser, err)
	}

	if result.RowsAffected() == 0 {
		return userport.ErrUserNotFound
	}

	return nil
}

var _ userport.Repository = (*Users)(nil)
