package entity

import "time"

type ProgramDonasiEntity struct {
	ID                  int64      `db:"id"`
	IDPPCPProgramDonasi string     `db:"id_pp_cp_program_donasi"`
	Title               string     `db:"title"`
	SubTitle            string     `db:"sub_title"`
	DonasiType          string     `db:"donasi_type"`
	Tag                 string     `db:"tag"`
	ThumbnailImageURL   string     `db:"thumbnail_image_url"`
	Description         string     `db:"description"`
	Status              string     `db:"status"`
	ValidFrom           *time.Time `db:"valid_from"`
	ValidTo             *time.Time `db:"valid_to"`
	Target              *float64   `db:"target"`
	CreatedAt           time.Time  `db:"created_at"`
	CreatedBy           *string    `db:"created_by"`
	UpdatedAt           *time.Time `db:"updated_at"`
	UpdatedBy           *string    `db:"updated_by"`
	PublishedAt         *time.Time `db:"published_at"`
	PublishedBy         *string    `db:"published_by"`
	IsDeleted           bool       `db:"is_deleted"`
	IsShow              bool       `db:"is_show"`
	KitaBisaLink        *string    `db:"kitabisa_link"`
	AyoBantuLink        *string    `db:"ayobantu_link"`
	IDPPCPMasterQris    *string    `db:"id_pp_cp_master_qris"`
	QrisImageURL        *string    `db:"qris_image_url"`

	Detail ProgramDonasiDetailEntity
}

type ProgramDonasiDetailEntity struct {
	ID                        int64  `db:"id"`
	IDPPCPProgramDonasi       string `db:"id_pp_cp_program_donasi"`
	IDPPCPProgramDonasiDetail string `db:"id_pp_cp_program_donasi_detail"`
	Content                   string `db:"content"`
	Tag                       string `db:"tag"`
}

type ProgramDonasiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []ProgramDonasiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type ProgramDonasiFilterQueryEntity struct {
	Field   string
	Keyword string
}
