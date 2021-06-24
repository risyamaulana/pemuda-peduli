package entity

import "time"

type ProgramDonasiRutinEntity struct {
	ID                       int64      `db:"id"`
	IDPPCPProgramDonasiRutin string     `db:"id_pp_cp_program_donasi_rutin"`
	Title                    string     `db:"title"`
	SubTitle                 string     `db:"sub_title"`
	DonasiType               string     `db:"donasi_type"`
	Tag                      string     `db:"tag"`
	ThumbnailImageURL        string     `db:"thumbnail_image_url"`
	Description              string     `db:"description"`
	Status                   string     `db:"status"`
	CreatedAt                time.Time  `db:"created_at"`
	CreatedBy                *string    `db:"created_by"`
	UpdatedAt                *time.Time `db:"updated_at"`
	UpdatedBy                *string    `db:"updated_by"`
	PublishedAt              *time.Time `db:"published_at"`
	PublishedBy              *string    `db:"published_by"`
	IsDeleted                bool       `db:"is_deleted"`
}

type ProgramDonasiRutinQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []ProgramDonasiRutinFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type ProgramDonasiRutinFilterQueryEntity struct {
	Field   string
	Keyword string
}
