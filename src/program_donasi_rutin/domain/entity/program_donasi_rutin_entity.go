package entity

import "time"

type ProgramDonasiRutinEntity struct {
	ID                          int64      `db:"id"`
	IDPPCPProgramDonasiRutin    string     `db:"id_pp_cp_program_donasi_rutin"`
	IDPPCPProgramDonasiKategori string     `db:"id_pp_cp_master_kategori_donasi"`
	KategoriName                string     `db:"kategori_name"`
	Title                       string     `db:"title"`
	SubTitle                    string     `db:"sub_title"`
	Content                     string     `db:"content"`
	Tag                         string     `db:"tag"`
	SEOURL                      string     `db:"seo_url"`
	ThumbnailImageURL           string     `db:"thumbnail_image_url"`
	DonationCollect             float64    `db:"donation_collect"`
	Description                 string     `db:"description"`
	Status                      string     `db:"status"`
	CreatedAt                   time.Time  `db:"created_at"`
	CreatedBy                   *string    `db:"created_by"`
	UpdatedAt                   *time.Time `db:"updated_at"`
	UpdatedBy                   *string    `db:"updated_by"`
	PublishedAt                 *time.Time `db:"published_at"`
	PublishedBy                 *string    `db:"published_by"`
	IsDeleted                   bool       `db:"is_deleted"`
	IsShow                      bool       `db:"is_show"`
	IDPPCPMasterQris            *string    `db:"id_pp_cp_master_qris"`
	QrisImageURL                *string    `db:"qris_image_url"`
}

type ProgramDonasiRutinPaketEntity struct {
	ID                            int64      `db:"id"`
	IDPPCPProgramDonasiRutinPaket string     `db:"id_pp_cp_program_donasi_rutin_paket"`
	IDPPCPProgramDonasiRutin      string     `db:"id_pp_cp_program_donasi_rutin"`
	SeoURL                        string     `db:"seo_url"`
	PaketName                     string     `db:"paket_name"`
	Benefit                       string     `db:"benefit"`
	Nominal                       float64    `db:"nominal"`
	PaketImageURL                 string     `db:"paket_image_url"`
	CreatedAt                     time.Time  `db:"created_at"`
	CreatedBy                     *string    `db:"created_by"`
	UpdatedAt                     *time.Time `db:"updated_at"`
	UpdatedBy                     *string    `db:"updated_by"`
	IsDeleted                     bool       `db:"is_deleted"`
}

type ProgramDonasiRutinNewsEntity struct {
	ID                       int64      `db:"id"`
	IDPPCPProgramDonasiRutin string     `db:"id_pp_cp_program_donasi_rutin"`
	Title                    string     `db:"title"`
	SubmitAt                 time.Time  `db:"submit_at"`
	DisbursementBalance      float64    `db:"disbursement_balance"`
	DisbursementAccount      string     `db:"disbursement_account"`
	DibursementBankName      string     `db:"dibursement_bank_name"`
	DisbursementName         string     `db:"disbursement_name"`
	DisbursementDescription  string     `db:"disbursement_description"`
	IsDeleted                bool       `db:"is_deleted"`
	CreatedAt                time.Time  `db:"created_at"`
	CreatedBy                *string    `db:"created_by"`
	UpdatedAt                *time.Time `db:"updated_at"`
	UpdatedBy                *string    `db:"updated_by"`
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
