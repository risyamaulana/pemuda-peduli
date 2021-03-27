package entity

import "time"

type RoleEntity struct {
	ID               int64      `db:"id"`
	IDPPCPMasterRole string     `db:"id_pp_cp_master_role"`
	RoleType         string     `db:"role_type"`
	RoleLevel        int        `db:"role_level"`
	CreatedAt        time.Time  `db:"created_at"`
	CreatedBy        *string    `db:"created_by"`
	UpdatedAt        *time.Time `db:"updated_at"`
	UpdatedBy        *string    `db:"updated_by"`
}

type RoleQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []RoleFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
}

type RoleFilterQueryEntity struct {
	Field   string
	Keyword string
}
