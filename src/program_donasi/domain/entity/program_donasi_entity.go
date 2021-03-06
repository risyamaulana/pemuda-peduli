package entity

import (
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"time"
)

type ProgramDonasiEntity struct {
	ID                   int64      `db:"id"`
	IDPPCPProgramDonasi  string     `db:"id_pp_cp_program_donasi"`
	KategoriID           *string    `db:"kategori_id"`
	KategoriName         *string    `db:"kategori_name"`
	Title                string     `db:"title"`
	SubTitle             string     `db:"sub_title"`
	DonasiType           string     `db:"donasi_type"`
	Tag                  string     `db:"tag"`
	ThumbnailImageURL    string     `db:"thumbnail_image_url"`
	Description          string     `db:"description"`
	IDPPCPPenggalangDana string     `db:"id_pp_cp_penggalang_dana"`
	Status               string     `db:"status"`
	ValidFrom            *time.Time `db:"valid_from"`
	ValidTo              *time.Time `db:"valid_to"`
	Nominal              string     `db:"nominal"`
	Target               *float64   `db:"target"`
	DonationCollect      float64    `db:"donation_collect"`
	CreatedAt            time.Time  `db:"created_at"`
	CreatedBy            *string    `db:"created_by"`
	UpdatedAt            *time.Time `db:"updated_at"`
	UpdatedBy            *string    `db:"updated_by"`
	PublishedAt          *time.Time `db:"published_at"`
	PublishedBy          *string    `db:"published_by"`
	IsDeleted            bool       `db:"is_deleted"`
	IsShow               bool       `db:"is_show"`
	KitaBisaLink         *string    `db:"kitabisa_link"`
	AyoBantuLink         *string    `db:"ayobantu_link"`
	IDPPCPMasterQris     *string    `db:"id_pp_cp_master_qris"`
	QrisImageURL         *string    `db:"qris_image_url"`
	SEOURL               string     `db:"seo_url"`

	Detail         ProgramDonasiDetailEntity
	PenggalangDana entity.PenggalangDanaEntity
}

type ProgramDonasiDetailEntity struct {
	ID                        int64  `db:"id"`
	IDPPCPProgramDonasi       string `db:"id_pp_cp_program_donasi"`
	IDPPCPProgramDonasiDetail string `db:"id_pp_cp_program_donasi_detail"`
	Content                   string `db:"content"`
	Tag                       string `db:"tag"`
}

type ProgramDonasiNewsEntity struct {
	ID                      int64      `db:"id"`
	IDPPCPProgramDonasi     string     `db:"id_pp_cp_program_donasi"`
	Title                   string     `db:"title"`
	SubmitAt                time.Time  `db:"submit_at"`
	DisbursementBalance     float64    `db:"disbursement_balance"`
	DisbursementAccount     string     `db:"disbursement_account"`
	DisbursementBankName    string     `db:"disbursement_bank_name"`
	DisbursementName        string     `db:"disbursement_name"`
	DisbursementDescription string     `db:"disbursement_description"`
	IsDeleted               bool       `db:"is_deleted"`
	CreatedAt               time.Time  `db:"created_at"`
	CreatedBy               *string    `db:"created_by"`
	UpdatedAt               *time.Time `db:"updated_at"`
	UpdatedBy               *string    `db:"updated_by"`
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
