package entity

import "time"

type TeamFlagEntity struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	IsDeleted bool       `db:"is_deleted"`
}
