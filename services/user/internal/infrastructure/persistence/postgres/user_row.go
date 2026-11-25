package postgres

import (
	"time"

	"github.com/LarsSeverson/charter/internal/lifecycle"
	"github.com/LarsSeverson/charter/internal/optional"
	"github.com/LarsSeverson/charter/internal/retention"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type userRow struct {
	ID     string `db:"id"`
	Status string `db:"status"`

	CreateTime time.Time  `db:"create_time"`
	UpdateTime time.Time  `db:"update_time"`
	DeleteTime *time.Time `db:"delete_time"`

	ExpireTime *time.Time `db:"expire_time"`
	PurgeTime  *time.Time `db:"purge_time"`
}

func (row userRow) domain() (*user.User, error) {
	id, err := user.ParseID(row.ID)
	if err != nil {
		return nil, err
	}

	lc, err := lifecycle.Reconstitute(
		row.CreateTime,
		row.UpdateTime,
		optional.FromPointer(row.DeleteTime),
	)
	if err != nil {
		return nil, err
	}

	ret, err := retention.New(
		optional.FromPointer(row.ExpireTime),
		optional.FromPointer(row.PurgeTime),
	)
	if err != nil {
		return nil, err
	}

	return user.Reconstitute(id, user.Status(row.Status), lc, ret)
}
