package channel

import (
	"regexp"
	"strings"
)

const maxCodeLength = 64

var codePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

type Code string

const (
	Email   Code = "email"
	WebPush Code = "web_push"
	InApp   Code = "in_app"
)

func ParseCode(input string) (Code, error) {
	value := strings.TrimSpace(input)
	code := Code(value)

	if !code.IsValid() {
		return "", ErrInvalidCode
	}

	return code, nil
}

func (c Code) IsValid() bool {
	value := string(c)
	length := len(value)
	return length > 0 && length <= maxCodeLength && codePattern.MatchString(value)
}

func (c Code) IsZero() bool {
	return c == ""
}

func (c Code) String() string {
	return string(c)
}
