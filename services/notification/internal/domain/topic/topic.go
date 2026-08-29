package topic

// A topic is a named type of notification describing what the notification is about
// Examples - project.mention account.security_alert

type Topic struct {
	code   Code
	status Status
}

func New(code Code) (*Topic, error) {
	if !code.IsValid() {
		return nil, ErrInvalidCode
	}

	return &Topic{
		code:   code,
		status: Active,
	}, nil
}

func Reconstitute(code Code, status Status) (*Topic, error) {
	if !code.IsValid() {
		return nil, ErrInvalidCode
	}

	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}

	return &Topic{
		code:   code,
		status: status,
	}, nil
}

func (t Topic) Code() Code {
	return t.code
}

func (t Topic) Status() Status {
	return t.status
}

func (t Topic) IsAvailable() bool {
	return t.status == Active
}

func (t *Topic) Activate() {
	t.status = Active
}

func (t *Topic) Disable() {
	t.status = Disabled
}
