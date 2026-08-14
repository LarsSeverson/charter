package parse

import (
	"errors"
	"strconv"
	"strings"

	"github.com/LarsSeverson/charter/internal/optional"
)

func Int32(raw string, fallback optional.Option[int32]) (int32, error) {
	raw = strings.TrimSpace(raw)

	if raw == "" {
		value, ok := fallback.Get()
		if !ok {
			return 0, ErrValueRequired
		}

		return value, nil
	}

	value, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return 0, errors.Join(ErrParseInt32, err)
	}

	return int32(value), nil
}
