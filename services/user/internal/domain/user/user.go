package user

import (
	"time"

	"github.com/LarsSeverson/charter/internal/lifecycle"
	"github.com/LarsSeverson/charter/internal/retention"
)

type User struct {
	id     ID
	status Status

	lifecycle lifecycle.Lifecycle
	retention retention.Retention
}

func New(id ID, now time.Time) (*User, error) {
	if id.IsZero() {
		return nil, ErrInvalidID
	}

	lifecycle := lifecycle.New(now)
	retention := retention.None()

	return &User{
		id:        id,
		status:    Pending,
		lifecycle: lifecycle,
		retention: retention,
	}, nil
}

func Reconstitute(
	id ID,
	status Status,
	lifecycle lifecycle.Lifecycle,
	retention retention.Retention,
) (*User, error) {
	if id.IsZero() {
		return nil, ErrInvalidID
	}

	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}

	return &User{
		id:        id,
		status:    status,
		lifecycle: lifecycle,
		retention: retention,
	}, nil
}

func (u User) ID() ID {
	return u.id
}

func (u User) Status() Status {
	return u.status
}

func (u *User) Suspend() error {
	if u.status == Closed {
		return ErrInvalidStatusChange
	}

	u.status = Suspended
	return nil
}
