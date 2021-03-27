package entity

import (
	roleMod "pemuda-peduli/src/role/domain/entity"
	"time"
)

type AdminUserEntity struct {
	ID              int64  `db:"id"`
	IDPPCPAdminUser string `db:"id_user_admin"`
	Username        string `db:"username"`
	Salt            string `db:"salt"`
	Password        string `db:"password"`
	Email           string `db:"email"`
	NamaLengkap     string `db:"nama_lengkap"`
	Alamat          string `db:"alamat"`
	Role            string `db:"role"`
	RoleData        roleMod.RoleEntity
	RoleLevel       int        `db:"role_level"`
	CreatedAt       time.Time  `db:"created_at"`
	CreatedBy       *string    `db:"created_by"`
	UpdatedAt       *time.Time `db:"updated_at"`
	UpdatedBy       *string    `db:"updated_by"`
	IsDeleted       bool       `db:"is_deleted"`
}

type AdminUserQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []AdminUserFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type AdminUserFilterQueryEntity struct {
	Field   string
	Keyword string
}
