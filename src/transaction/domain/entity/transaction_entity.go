package entity

import "time"

type TransactionEntity struct {
	ID                       int64      `db:"id"`
	IDPPTransaction          string     `db:"id_pp_transaction"`
	UserID                   string     `db:"user_id"`
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
	QRPaymentURL             *string    `db:"qr_payment_url"`
	ImagePaymentURL          string     `db:"image_payment_url"`
	PaidAt                   *time.Time `db:"paid_at"`
	ApprovedAt               *time.Time `db:"approved_at"`
	ApprovedBy               string     `db:"approved_by"`
	Status                   string     `db:"status"` // Unpaid, Paid, Need Approval, Decline
	CreatedAt                time.Time  `db:"created_at"`
	UpdatedAt                *time.Time `db:"updated_at"`
}

type TransactionQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []TransactionFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PaidAtFrom    string `db:"paid_at_from"`
	PaidAtTo      string `db:"paid_at_to"`
}

type TransactionFilterQueryEntity struct {
	Field   string
	Keyword string
}
