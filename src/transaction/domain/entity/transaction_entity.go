package entity

import "time"

type TransactionEntity struct {
	ID                       int64      `db:"id"`
	IDPPTransaction          string     `db:"id_pp_transaction"`
	Username                 string     `db:"username"`
	Email                    string     `db:"email"`
	NamaLengkap              string     `db:"nama_lengkap"`
	NamaPanggilan            string     `db:"nama_panggilan"`
	IsRutin                  bool       `db:"is_rutin"`
	IDPPCPProgramDonasi      *string    `db:"id_pp_cp_program_donasi"`
	IDPPCPProgramDonasiRutin *string    `db:"id_pp_cp_program_donasi_rutin"`
	DonasiTitle              string     `db:"donasi_title"`
	Amount                   float64    `db:"amount"`
	PaymentMethod            string     `db:"payment_method"` // enum : qris & transfer manual
	ImagePaymentURL          string     `db:"image_payment_url"`
	PaidAt                   *time.Time `db:"paid_at"`
	ApprovedAt               *time.Time `db:"approved_at"`
	ApprovedBy               string     `db:"approved_by"`
	Status                   string     `db:"status"` // Unpaid, Paid, Need Approval
	CreatedAt                time.Time  `db:"created_at"`
	CreatedBy                *string    `db:"created_by"`
}

type ProgramDonasiQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []ProgramDonasiFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
}

type ProgramDonasiFilterQueryEntity struct {
	Field   string
	Keyword string
}
