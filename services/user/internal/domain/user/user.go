package user

type User struct {
	id     ID
	status Status
}

func New(id ID) (*User, error) {
	if id.IsZero() {
		return nil, ErrInvalidID
	}

	return &User{
		id:     id,
		status: Pending,
	}, nil
}

func Reconstitute(id ID, status Status) (*User, error) {
	if id.IsZero() {
		return nil, ErrInvalidID
	}

	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}

	return &User{
		id:     id,
		status: status,
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
