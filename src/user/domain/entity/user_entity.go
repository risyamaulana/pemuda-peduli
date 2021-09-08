package entity

import "time"

type UserEntity struct {
	ID            int64      `db:"id"`
	IDUser        string     `db:"id_user"`
	Username      string     `db:"username"`
	Salt          string     `db:"salt"`
	Password      string     `db:"password"`
	Email         string     `db:"email"`
	NamaLengkap   string     `db:"nama_lengkap"`
	NamaPanggilan string     `db:"nama_panggilan"`
	Alamat        string     `db:"alamat"`
	PhoneNumber   string     `db:"phone_number"`
	IsReset       bool       `db:"is_reset"`
	TokenReset    string     `db:"token_reset"`
	TokenValid    *time.Time `db:"token_valid"`
	IsDeleted     bool       `db:"is_deleted"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
}
