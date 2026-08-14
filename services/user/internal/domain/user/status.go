package user

type Status string

const (
	Pending   Status = "pending"
	Active    Status = "active"
	Suspended Status = "suspended"
	Closed    Status = "closed"
)

func (s Status) IsValid() bool {
	switch s {
	case Pending, Active, Suspended, Closed:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	return string(s)
}
