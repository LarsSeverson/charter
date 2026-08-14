package email

import (
	"time"

	"github.com/LarsSeverson/charter/internal/lifecycle"
	"github.com/LarsSeverson/charter/internal/optional"
	"github.com/LarsSeverson/charter/internal/retention"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type Type string

const (
	TypeHome  Type = "home"
	TypeWork  Type = "work"
	TypeOther Type = "other"
)

type Email struct {
	id         ID
	userID     user.ID
	address    Address
	emailType  Type
	primary    bool
	verifyTime optional.Option[time.Time]

	lifecycle lifecycle.Lifecycle
	retention retention.Retention
}

func New(
	id ID,
	userID user.ID,
	address Address,
	emailType Type,
	primary bool,
	now time.Time,
) (Email, error) {
	if id.IsZero() {
		return Email{}, ErrInvalidID
	}

	if userID.IsZero() {
		return Email{}, ErrInvalidUserID
	}

	if address.IsZero() {
		return Email{}, ErrInvalidEmailAddress
	}

	if !emailType.IsValid() {
		return Email{}, ErrInvalidEmailType
	}

	verifyTime := optional.None[time.Time]()
	lifecycle := lifecycle.New(now)
	retention := retention.None()

	return Email{
		id:         id,
		userID:     userID,
		address:    address,
		emailType:  emailType,
		primary:    primary,
		verifyTime: verifyTime,

		lifecycle: lifecycle,
		retention: retention,
	}, nil
}

func Reconstitute(
	id ID,
	userID user.ID,
	address Address,
	emailType Type,
	primary bool,
	verifyTime optional.Option[time.Time],
	lifecycle lifecycle.Lifecycle,
	retention retention.Retention,
) (Email, error) {
	if address.IsZero() {
		return Email{}, ErrInvalidEmailAddress
	}

	if !emailType.IsValid() {
		return Email{}, ErrInvalidEmailType
	}

	return Email{
		id:         id,
		userID:     userID,
		address:    address,
		emailType:  emailType,
		primary:    primary,
		verifyTime: verifyTime,
		lifecycle:  lifecycle,
		retention:  retention,
	}, nil
}

func (t Type) IsValid() bool {
	switch t {
	case TypeHome, TypeWork, TypeOther:
		return true
	default:
		return false
	}
}

func (t Type) String() string {
	return string(t)
}
