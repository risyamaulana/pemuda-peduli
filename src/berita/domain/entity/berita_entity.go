package entity

import "time"

type BeritaEntity struct {
	ID                int64      `db:"id"`
	IDPPCPBerita      string     `db:"id_pp_cp_berita"`
	Title             string     `db:"title"`
	SubTitle          string     `db:"sub_title"`
	Tag               string     `db:"tag"`
	ThumbnailImageURL string     `db:"thumbnail_image_url"`
	Description       string     `db:"description"`
	Status            string     `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	CreatedBy         *string    `db:"created_by"`
	UpdatedAt         *time.Time `db:"updated_at"`
	UpdatedBy         *string    `db:"updated_by"`
	PublishedAt       *time.Time `db:"published_at"`
	PublishedBy       *string    `db:"published_by"`
	IsDeleted         bool       `db:"is_deleted"`

	Detail BeritaDetailEntity
}

type BeritaDetailEntity struct {
	ID                 int64  `db:"id"`
	IDPPCPBerita       string `db:"id_pp_cp_berita"`
	IDPPCPBeritaDetail string `db:"id_pp_cp_berita_detail"`
	Content            string `db:"content"`
	Tag                string `db:"tag"`
}

type BeritaQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []BeritaFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type BeritaFilterQueryEntity struct {
	Field   string
	Keyword string
}
