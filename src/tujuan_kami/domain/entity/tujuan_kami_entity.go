package entity

import "time"

type TujuanKamiEntity struct {
	ID               int64      `db:"id"`
	IDPPCPTujuanKami string     `db:"id_pp_cp_tujuan_kami"`
	Title            string     `db:"title"`
	Description      string     `db:"description"`
	Icon             string     `db:"icon"`
	Status           string     `db:"status"`
	CreatedAt        time.Time  `db:"created_at"`
	CreatedBy        *string    `db:"created_by"`
	UpdatedAt        *time.Time `db:"updated_at"`
	UpdatedBy        *string    `db:"updated_by"`
	PublishedAt      *time.Time `db:"published_at"`
	PublishedBy      *string    `db:"published_by"`
	IsDeleted        bool       `db:"is_deleted"`
}

type TujuanKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []TujuanKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type TujuanKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
