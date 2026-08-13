package lifecycle

import (
	"time"

	"github.com/LarsSeverson/charter/internal/optional"
)

type Lifecycle struct {
	createTime time.Time
	updateTime time.Time
	deleteTime optional.Option[time.Time]
}

func New(now time.Time) Lifecycle {
	now = now.UTC()

	return Lifecycle{
		createTime: now,
		updateTime: now,
		deleteTime: optional.None[time.Time](), // Could omit
	}
}

func Reconstitute(
	createTime time.Time,
	updateTime time.Time,
	deleteTime optional.Option[time.Time],
) (Lifecycle, error) {
	createTime = createTime.UTC()
	updateTime = updateTime.UTC()

	oneZero := createTime.IsZero() || updateTime.IsZero()
	if oneZero {
		return Lifecycle{}, ErrInvalidLifecycle
	}

	if updateTime.Before(createTime) {
		return Lifecycle{}, ErrInvalidLifecycle
	}

	if value, exists := deleteTime.Get(); exists {
		value = value.UTC()

		if value.Before(createTime) || updateTime.Before(value) {
			return Lifecycle{}, ErrInvalidLifecycle
		}

		deleteTime = optional.Some(value)
	}

	return Lifecycle{
		createTime: createTime,
		updateTime: updateTime,
		deleteTime: deleteTime,
	}, nil
}

func (l *Lifecycle) Touch(now time.Time) *Lifecycle {
	l.updateTime = now.UTC()
	return l
}

func (l *Lifecycle) Delete(now time.Time) (*Lifecycle, error) {
	if l.deleteTime.IsSet() {
		return nil, ErrAlreadyDeleted
	}

	now = now.UTC()

	l.deleteTime = optional.Some(now)
	l.updateTime = now

	return l, nil
}

func (l Lifecycle) IsDeleted() bool {
	return l.deleteTime.IsSet()
}

func (l Lifecycle) CreateTime() time.Time {
	return l.createTime
}

func (l Lifecycle) UpdateTime() time.Time {
	return l.updateTime
}

func (l Lifecycle) DeleteTime() optional.Option[time.Time] {
	return l.deleteTime
}
