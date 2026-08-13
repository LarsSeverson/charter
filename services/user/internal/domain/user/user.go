package user

type User struct {
	id     UserID
	status Status
}

func New(id UserID) *User {
	status := Pending

	return &User{
		id,
		status,
	}
}

func (u *User) Suspend() (*User, error) {
	u.status = Suspended

	return u, nil
}
