package user

type User struct {
	id     ID
	status Status
}

func New(id ID) *User {
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
