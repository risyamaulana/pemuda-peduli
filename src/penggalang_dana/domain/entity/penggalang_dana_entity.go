package entity

import "time"

type PenggalangDanaEntity struct {
	ID                   int64      `db:"id"`
	IDPPCPPenggalangDana string     `db:"id_pp_cp_penggalang_dana"`
	Name                 string     `db:"name"`
	Description          string     `db:"description"`
	ThumbnailImageURL    string     `db:"thumbnail_image_url"`
	IsVerified           bool       `db:"is_verified"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at"`
	IsDeleted            bool       `db:"is_deleted"`
}

type PenggalangDanaQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []PenggalangDanaFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
}

type PenggalangDanaFilterQueryEntity struct {
	Field   string
	Keyword string
}
