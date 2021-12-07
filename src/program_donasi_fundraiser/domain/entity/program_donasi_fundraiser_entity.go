package entity

import "time"

type ProgramDonasiFundraiserEntity struct {
	ID                            int64  `db:"id"`
	IDPPCPProgramDonasiFundraiser string `db:"id_pp_cp_program_donasi_fundraiser"`
	IDPPCPPenggalangDana          string `db:"id_pp_cp_penggalang_dana"`
	IDPPCPProgramDonasi           string `db:"id_pp_cp_program_donasi"`

	IDUser        string `db:"id_user"`
	Username      string `db:"username"`
	Email         string `db:"email"`
	PhoneNumber   string `db:"phone_number"`
	NamaLengkap   string `db:"nama_lengkap"`
	NamaPanggilan string `db:"nama_panggilan"`
	Alamat        string `db:"alamat"`

	Title      string `db:"title"`
	SubTitle   string `db:"sub_title"`
	DonasiType string `db:"donasi_type"`

	Tag               string `db:"tag"`
	Content           string `db:"content"`
	ThumbnailImageURL string `db:"thumbnail_image_url"`
	Description       string `db:"description"`

	Status    string     `db:"status"`
	ValidFrom *time.Time `db:"valid_from"`
	ValidTo   *time.Time `db:"valid_to"`

	Nominal         string   `db:"nominal"`
	Target          *float64 `db:"target"`
	DonationCollect float64  `db:"donation_collect"`

	CreatedAt   time.Time  `db:"created_at"`
	CreatedBy   *string    `db:"created_by"`
	UpdatedAt   *time.Time `db:"updated_at"`
	UpdatedBy   *string    `db:"updated_by"`
	PublishedAt *time.Time `db:"published_at"`
	PublishedBy *string    `db:"published_by"`

	IsDeleted bool `db:"is_deleted"`
	IsShow    bool `db:"is_show"`

	KitaBisaLink     *string `db:"kitabisa_link"`
	AyoBantuLink     *string `db:"ayobantu_link"`
	IDPPCPMasterQris *string `db:"id_pp_cp_master_qris"`
	QrisImageURL     *string `db:"qris_image_url"`

	SEOURL string `db:"seo_url"`
}

type ProgramDonasiFundraiserQueryEntity struct {
	Limit  string `db:"limit"`
	Offset string `db:"offset"`
	Filter []ProgramDonasiFundraiserFilterQueryEntity
	Order  string `db:"order"`
	Sort   string `db:"sort"`

	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type ProgramDonasiFundraiserFilterQueryEntity struct {
	Field   string
	Keyword string
}
