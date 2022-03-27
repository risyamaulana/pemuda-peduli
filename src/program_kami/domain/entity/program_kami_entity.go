package entity

import "time"

type ProgramKamiEntity struct {
	ID                   int64      `db:"id"`
	IDPPCPProgramKami    string     `db:"id_pp_cp_program_kami"`
	Title                string     `db:"title"`
	SubTitle             string     `db:"sub_title"`
	Tag                  string     `db:"tag"`
	ThumbnailImageURL    string     `db:"thumbnail_image_url"`
	BeneficariesImageURL string     `db:"beneficaries_image_url"`
	Achievements         string     `db:"achievements"`
	Document             *string    `db:"document"`
	Description          string     `db:"description"`
	Status               string     `db:"status"`
	CreatedAt            time.Time  `db:"created_at"`
	CreatedBy            *string    `db:"created_by"`
	UpdatedAt            *time.Time `db:"updated_at"`
	UpdatedBy            *string    `db:"updated_by"`
	PublishedAt          *time.Time `db:"published_at"`
	PublishedBy          *string    `db:"published_by"`
	IsDeleted            bool       `db:"is_deleted"`

	Detail ProgramKamiDetailEntity
}

type ProgramKamiDetailEntity struct {
	ID                      int64  `db:"id"`
	IDPPCPProgramKami       string `db:"id_pp_cp_program_kami"`
	IDPPCPProgramKamiDetail string `db:"id_pp_cp_program_kami_detail"`
	Content                 string `db:"content"`
	Tag                     string `db:"tag"`
}

type ProgramKamiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []ProgramKamiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type ProgramKamiFilterQueryEntity struct {
	Field   string
	Keyword string
}
