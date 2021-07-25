package entity

import "time"

type QrisEntity struct {
	ID                int64      `db:"id"`
	IDPPCPQris        string     `db:"id_pp_cp_master_qris"`
	Title             string     `db:"title"`
	Description       string     `db:"description"`
	ThumbnailImageURL string     `db:"thumbnail_image_url"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
}

type QrisQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []QrisFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
}

type QrisFilterQueryEntity struct {
	Field   string
	Keyword string
}
