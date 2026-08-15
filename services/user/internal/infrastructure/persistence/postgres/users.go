package postgres

import (
	"context"
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/application/port"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	db DBTX
}

func NewUsers(db DBTX) *Users {
	return &Users{db: db}
}

func (r *Users) FindByID(
	ctx context.Context,
	id user.ID,
) (*user.User, error) {
	const query = `
		SELECT id, status
		FROM users
		WHERE id = $1
	`

	// Change this to an actual struct for the row tailored to user
	var (
		storedID     string
		storedStatus string
	)

	err := r.db.QueryRow(ctx, query, id.String()).Scan(
		&storedID,
		&storedStatus,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, port.ErrUserNotFound
	}
	if err != nil {
		return nil, errors.Join(ErrFindUser, err)
	}

	parsedID, err := user.ParseID(storedID)
	if err != nil {
		return nil, errors.Join(ErrParseUserID, err)
	}

	value, err := user.Reconstitute(parsedID, user.Status(storedStatus))
	if err != nil {
		return nil, errors.Join(ErrReconstituteUser, err)
	}

	return value, err
}

func (r *Users) Create(
	ctx context.Context,
	value *user.User,
) error {
	const query = `
		INSERT INTO users (id, status)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		value.ID().String(),
		value.Status().String(),
	)
	// if isConstraint
	if err != nil {
		return errors.Join(ErrCreateUser, err)
	}

	return nil
}

func (r *Users) Update(
	ctx context.Context,
	value *user.User,
) error {
	const query = `
		UPDATE users
		SET status = $2,
			updated_at = now()
		WHERE id = $1
	`

	result, err := r.db.Exec(
		ctx,
		query,
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
