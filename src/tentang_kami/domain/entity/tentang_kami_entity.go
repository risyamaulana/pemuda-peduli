package entity

import "time"

type TentangKamiEntity struct {
	ID                int64      `db:"id"`
	IDPPCPTentangKami string     `db:"id_pp_cp_tentang_kami"`
	ThumbnailImageURL string     `db:"thumbnail_image_url"`
	Description       string     `db:"description"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`
}

type TentangKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []TentangKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type TentangKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
