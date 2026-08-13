package retention

import (
	"time"

	"github.com/LarsSeverson/charter/internal/optional"
)

type Retention struct {
	expireTime optional.Option[time.Time]
	purgeTime  optional.Option[time.Time]
}

func None() Retention {
	return Retention{}
}

func New(
	expireTime optional.Option[time.Time],
	purgeTime optional.Option[time.Time],
) (Retention, error) {
	expireValue, hasExpireTime := expireTime.Get()
	purgeValue, hasPurgeTime := purgeTime.Get()

	bothExist := hasExpireTime && hasPurgeTime
	if bothExist && purgeValue.Before(expireValue) {
		return Retention{}, ErrInvalidRetention
	}

	return Retention{
		expireTime: expireTime,
		purgeTime:  purgeTime,
	}, nil
}

func (r Retention) ExpireTime() optional.Option[time.Time] {
	return r.expireTime
}

func (r Retention) PurgeTime() optional.Option[time.Time] {
	return r.purgeTime
}

func (r Retention) IsExpired(now time.Time) bool {
	expireTime, exists := r.expireTime.Get()
	if !exists {
		return false
	}

	return !now.UTC().Before(expireTime)
}

func (r Retention) IsPurgeDue(now time.Time) bool {
	purgeTime, exists := r.purgeTime.Get()
	if !exists {
		return false
	}

	return !now.UTC().Before(purgeTime)
}
