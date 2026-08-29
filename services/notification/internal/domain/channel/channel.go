package channel

// A channel is a delivery method for notifactions. It tells the notifaction system
// how a message can be delivered.

type Channel struct {
	code   Code
	status Status
}

func New(code Code) (*Channel, error) {
	if !code.IsValid() {
		return nil, ErrInvalidCode
	}

	return &Channel{
		code:   code,
		status: Active,
	}, nil
}

func Reconstitute(code Code, status Status) (*Channel, error) {
	if !code.IsValid() {
		return nil, ErrInvalidCode
	}

	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}

	return &Channel{
		code:   code,
		status: status,
	}, nil
}

func (c Channel) Code() Code {
	return c.code
}

func (c Channel) Status() Status {
	return c.status
}

func (c Channel) IsAvailable() bool {
	return c.status == Active
}

func (c *Channel) Activate() {
	c.status = Active
}

func (c *Channel) Disable() {
	c.status = Disabled
}
