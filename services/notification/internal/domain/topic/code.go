package topic

import (
	"regexp"
	"strings"
)

const maxCodeLength = 128

var codePattern = regexp.MustCompile(
	`^[a-z][a-z0-9_]*(\.[a-z][a-z0-9_]*)*$`,
)

type Code string

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
