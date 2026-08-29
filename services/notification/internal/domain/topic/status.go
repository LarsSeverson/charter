package topic

type Status string

const (
	Active   Status = "active"
	Disabled Status = "disabled"
)

func (s Status) IsValid() bool {
	switch s {
	case Active, Disabled:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	return string(s)
}
