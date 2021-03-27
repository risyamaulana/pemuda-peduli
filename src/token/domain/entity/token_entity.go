package entity

import "time"

type TokenEntity struct {
	ID                  int64      `db:"id"`
	Name                string     `db:"name"`
	DeviceID            string     `db:"device_id"`
	DeviceType          string     `db:"device_type"`
	Token               string     `db:"token"`
	TokenExpired        time.Time  `db:"token_expired"`
	RefreshToken        string     `db:"refresh_token"`
	RefreshTokenExpired time.Time  `db:"refresh_token_expired"`
	IsLogin             bool       `db:"is_login"`
	LoginID             string     `db:"login_id"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           *time.Time `db:"updated_at"`
}
