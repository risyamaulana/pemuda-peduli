package applications

import (
	"encoding/json"
	"errors"

	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/transaction/common/constants"
	"pemuda-peduli/src/transaction/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTransaction struct {
	IsRutin                  bool    `json:"is_rutin"`
	IDPPCPProgramDonasi      *string `json:"id_pp_cp_program_donasi"`
	IDPPCPProgramDonasiRutin *string `json:"id_pp_cp_program_donasi_rutin"`
	Amount                   float64 `json:"amount" valid:"required"`
	PaymentMethod            string  `json:"payment_method" valid:"required"`
}

type UploadReceiptTransaction struct {
	TransactionID string `json:"transaction_id" valid:"required"`
	ImageURL      string `json:"image_url" valid:"required,url"`
}

type TransactionQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []TransactionFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PaidAtFrom    string                   `json:"paid_at_from"`
	PaidAtTo      string                   `json:"paid_at_to"`
}

type TransactionFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadTransaction struct {
	IDPPTransaction          string     `json:"id"`
	UserID                   string     `json:"user_id"`
	Username                 string     `json:"username"`
	Email                    string     `json:"email"`
	NamaLengkap              string     `json:"nama_lengkap"`
	NamaPanggilan            string     `json:"nama_panggilan"`
	IsRutin                  bool       `json:"is_rutin"`
	IDPPCPProgramDonasi      *string    `json:"id_pp_cp_program_donasi"`
	IDPPCPProgramDonasiRutin *string    `json:"id_pp_cp_program_donasi_rutin"`
	DonasiTitle              string     `json:"donasi_title"`
	Amount                   float64    `json:"amount"`
	PaymentMethod            string     `json:"payment_method"` // enum : qris & transfer manual
	QRPaymentURL             *string    `json:"qr_payment_url"`
	ImagePaymentURL          string     `json:"image_payment_url"`
	PaidAt                   *time.Time `json:"paid_at"`
	ApprovedAt               *time.Time `json:"approved_at"`
	ApprovedBy               string     `json:"approved_by"`
	Status                   string     `json:"status"` // Unpaid, Paid, Need Approval, Decline
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                *time.Time `json:"updated_at"`
}

func GetCreatePayload(body []byte) (payload CreateTransaction, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUploadReceiptPayload(body []byte) (payload UploadReceiptTransaction, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload TransactionQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateTransaction) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.PaymentMethod != constants.PaymentQris && r.PaymentMethod != constants.PaymentManual {
		err = errors.New("Failed, payment method not found")
		return
	}

	return
}

func (r UploadReceiptTransaction) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r TransactionQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateTransaction) ToEntity() (data entity.TransactionEntity) {
	data = entity.TransactionEntity{
		IDPPTransaction:          utility.GetUUID(),
		IsRutin:                  data.IsRutin,
		IDPPCPProgramDonasi:      data.IDPPCPProgramDonasi,
		IDPPCPProgramDonasiRutin: data.IDPPCPProgramDonasiRutin,
		Amount:                   data.Amount,
		PaymentMethod:            data.PaymentMethod,
		Status:                   constants.StatusUnpaid,
		CreatedAt:                time.Now(),
	}
	return
}

func (r TransactionQuery) ToEntity() (data entity.TransactionQueryEntity) {
	filters := []entity.TransactionFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.TransactionFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.TransactionQueryEntity{
		Limit:         r.Limit,
		Offset:        r.Offset,
		Filter:        filters,
		Order:         r.Order,
		Sort:          r.Sort,
		CreatedAtFrom: r.CreatedAtFrom,
		CreatedAtTo:   r.CreatedAtTo,
		PaidAtFrom:    r.PaidAtFrom,
		PaidAtTo:      r.PaidAtTo,
	}
	return
}

func ToPayload(data entity.TransactionEntity) (response ReadTransaction) {
	response = ReadTransaction{
		IDPPTransaction:          data.IDPPTransaction,
		UserID:                   data.UserID,
		Username:                 data.Username,
		Email:                    data.Email,
		NamaLengkap:              data.NamaLengkap,
		NamaPanggilan:            data.NamaPanggilan,
		IsRutin:                  data.IsRutin,
		IDPPCPProgramDonasi:      data.IDPPCPProgramDonasi,
		IDPPCPProgramDonasiRutin: data.IDPPCPProgramDonasiRutin,
		DonasiTitle:              data.DonasiTitle,
		Amount:                   data.Amount,
		PaymentMethod:            data.PaymentMethod,
		QRPaymentURL:             data.QRPaymentURL,
		ImagePaymentURL:          data.ImagePaymentURL,
		PaidAt:                   data.PaidAt,
		ApprovedAt:               data.ApprovedAt,
		ApprovedBy:               data.ApprovedBy,
		Status:                   data.Status,
		CreatedAt:                data.CreatedAt,
		UpdatedAt:                data.UpdatedAt,
	}
	return
}
