package user

type Status string

const (
	Pending   Status = "pending"
	Active    Status = "active"
	Suspended Status = "suspended"
	Closed    Status = "closed"
)

func (s Status) IsValid() bool {
	return true
}

func (s Status) String() string {
	return string(s)
}
