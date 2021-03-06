package applications

import (
	"encoding/json"
	"errors"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"

	penggalangDanaApp "pemuda-peduli/src/penggalang_dana/applications"
)

type CreateProgramDonasi struct {
	KategoriID           *string    `json:"kategori_id"`
	Title                string     `json:"title" valid:"required"`
	SubTitle             string     `json:"sub_title" valid:"required"`
	Content              string     `json:"content" valid:"required"`
	Tag                  string     `json:"tag"`
	ThumbnailImageURL    string     `json:"thumbnail_image_url" valid:"url"`
	IDPPCPPenggalangDana string     `json:"id_pp_cp_penggalang_dana"`
	SEOURL               string     `json:"seo_url"`
	ValidFrom            *time.Time `json:"valid_from" valid:"required"`
	ValidTo              *time.Time `json:"valid_to" valid:"required"`
	Nominal              []int      `json:"nominal" valid:"required"`
	Target               *float64   `json:"target" valid:"required"`
	KitaBisaLink         *string    `json:"kitabisa_link" valid:"url"`
	AyoBantuLink         *string    `json:"ayobantu_link" valid:"url"`
	IDPPCPMasterQris     string     `json:"id_pp_cp_master_qris"`
	QrisImageURL         string     `json:"qris_image_url"`
	Description          string     `json:"description"`
}

type CreateProgramDonasiNews struct {
	IDPPCPProgramDonasi     string    `json:"id_pp_cp_program_donasi" valid:"required"`
	Title                   string    `json:"title" valid:"required"`
	SubmitAt                time.Time `json:"submit_at" valid:"required"`
	DisbursementBalance     float64   `json:"disbursement_balance" valid:"required"`
	DisbursementAccount     string    `json:"disbursement_account" valid:"required"`
	DisbursementBankName    string    `json:"disbursement_bank_name" valid:"required"`
	DisbursementName        string    `json:"disbursement_name" valid:"required"`
	DisbursementDescription string    `json:"disbursement_description" valid:"required"`
}

type UpdateProgramDonasi struct {
	KategoriID           *string    `json:"kategori_id"`
	Title                string     `json:"title" valid:"required"`
	SubTitle             string     `json:"sub_title" valid:"required"`
	Content              string     `json:"content" valid:"required"`
	Tag                  string     `json:"tag"`
	IDPPCPPenggalangDana string     `json:"id_pp_cp_penggalang_dana"`
	SEOURL               string     `json:"seo_url"`
	ThumbnailImageURL    string     `json:"thumbnail_image_url" valid:"url"`
	ValidFrom            *time.Time `json:"valid_from" valid:"required"`
	ValidTo              *time.Time `json:"valid_to" valid:"required"`
	Nominal              []int      `json:"nominal" valid:"required"`
	Target               *float64   `json:"target" valid:"required"`
	KitaBisaLink         *string    `json:"kitabisa_link" valid:"url"`
	AyoBantuLink         *string    `json:"ayobantu_link" valid:"url"`
	IDPPCPMasterQris     string     `json:"id_pp_cp_master_qris"`
	QrisImageURL         string     `json:"qris_image_url"`
	Description          string     `json:"description"`
	IsShow               *bool      `json:"is_show"`
}

type UpdateProgramDonasiNews struct {
	Title                   string    `json:"title" valid:"required"`
	SubmitAt                time.Time `json:"submit_at" valid:"required"`
	DisbursementBalance     float64   `json:"disbursement_balance" valid:"required"`
	DisbursementAccount     string    `json:"disbursement_account" valid:"required"`
	DisbursementBankName    string    `json:"disbursement_bank_name" valid:"required"`
	DisbursementName        string    `json:"disbursement_name" valid:"required"`
	DisbursementDescription string    `json:"disbursement_description" valid:"required"`
}

type ProgramDonasiQuery struct {
	Limit         string                     `json:"limit" valid:"required"`
	Offset        string                     `json:"offset" valid:"required"`
	Filter        []ProgramDonasiFilterQuery `json:"filters"`
	Order         string                     `json:"order"`
	Sort          string                     `json:"sort"`
	CreatedAtFrom string                     `json:"created_at_from"`
	CreatedAtTo   string                     `json:"created_at_to"`
	PublishAtFrom string                     `json:"publish_at_from"`
	PublishAtTo   string                     `json:"publish_at_to"`
}

type ProgramDonasiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramDonasi struct {
	IDPPCPProgramDonasi  string     `json:"id"`
	KategoriID           *string    `json:"kategori_id"`
	KategoriName         *string    `json:"kategori_name"`
	Title                string     `json:"title"`
	SubTitle             string     `json:"sub_title"`
	Content              string     `json:"content"`
	Tag                  string     `json:"tag"`
	ThumbnailImageURL    string     `json:"thumbnail_image_url"`
	ValidFrom            *time.Time `json:"valid_from"`
	IDPPCPPenggalangDana string     `json:"id_pp_cp_penggalang_dana"`
	ValidTo              *time.Time `json:"valid_to"`
	Nominal              []int      `json:"nominal" valid:"required"`
	Target               *float64   `json:"target"`
	DonationCollect      float64    `json:"donation_collect"`
	Description          string     `json:"description"`
	Status               string     `json:"status"`
	CreatedAt            time.Time  `json:"created_at"`
	CreatedBy            *string    `json:"created_by"`
	UpdatedAt            *time.Time `json:"updated_at"`
	UpdatedBy            *string    `json:"updated_by"`
	PublishedAt          *time.Time `json:"published_at"`
	PublishedBy          *string    `json:"published_by"`
	IsDeleted            bool       `json:"is_deleted"`
	IsShow               bool       `json:"is_show"`
	KitaBisaLink         *string    `json:"kitabisa_link"`
	AyoBantuLink         *string    `json:"ayobantu_link"`
	IDPPCPMasterQris     *string    `json:"id_pp_cp_master_qris"`
	QrisImageURL         *string    `json:"qris_image_url"`
	SEOURL               string     `json:"seo_url"`

	PenggalangDana penggalangDanaApp.ReadPenggalangDana `json:"penggalang_dana"`
}

type ReadProgramDonasiNews struct {
	ID                      int64      `json:"id"`
	IDPPCPProgramDonasi     string     `json:"id_pp_cp_program_donasi"`
	Title                   string     `json:"title"`
	SubmitAt                time.Time  `json:"submit_at"`
	DisbursementBalance     float64    `json:"disbursement_balance"`
	DisbursementAccount     string     `json:"disbursement_account"`
	DisbursementBankName    string     `json:"disbursement_bank_name"`
	DisbursementName        string     `json:"disbursement_name"`
	DisbursementDescription string     `json:"disbursement_description"`
	IsDeleted               bool       `json:"is_deleted"`
	CreatedAt               time.Time  `json:"created_at"`
	CreatedBy               *string    `json:"created_by"`
	UpdatedAt               *time.Time `json:"updated_at"`
	UpdatedBy               *string    `json:"updated_by"`
}

func GetCreatePayload(body []byte) (payload CreateProgramDonasi, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramDonasi, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramDonasiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetCreateNewsPayload(body []byte) (payload CreateProgramDonasiNews, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdateNewsPayload(body []byte) (payload UpdateProgramDonasiNews, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramDonasi) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.ValidFrom.Before(time.Now().UTC()) {
		err = errors.New("Failed : valid from smallest than current time")
		return
	}

	if r.ValidTo.Before(time.Now().UTC()) {
		err = errors.New("Failed : valid to smallest than current time")
		return
	}

	if r.ValidFrom.After(*r.ValidTo) {
		err = errors.New("Failed : valid from bigger than valid to value")
		return
	}
	return
}

func (r UpdateProgramDonasi) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.IsShow == nil {
		err = errors.New("is_show: non zero value required")
		return
	}

	return
}

func (r ProgramDonasiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiNews) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateProgramDonasiNews) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasi) ToEntity() (data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity) {
	validFrom := r.ValidFrom.UTC()
	validTo := r.ValidTo.UTC()

	var nominalStr strings.Builder
	for _, nominal := range r.Nominal {
		nominalString := strconv.Itoa(nominal)
		nominalStr.WriteString(nominalString + "|")
	}
	nominalValue := strings.TrimSuffix(nominalStr.String(), "|")

	data = entity.ProgramDonasiEntity{
		KategoriID:           r.KategoriID,
		Title:                r.Title,
		SubTitle:             r.SubTitle,
		Tag:                  r.Tag,
		SEOURL:               r.SEOURL,
		ThumbnailImageURL:    r.ThumbnailImageURL,
		IDPPCPPenggalangDana: r.IDPPCPPenggalangDana,
		ValidFrom:            &validFrom,
		ValidTo:              &validTo,
		Nominal:              nominalValue,
		Target:               r.Target,
		KitaBisaLink:         r.KitaBisaLink,
		AyoBantuLink:         r.AyoBantuLink,
		IDPPCPMasterQris:     &r.IDPPCPMasterQris,
		QrisImageURL:         &r.QrisImageURL,
		Description:          r.Description,
		CreatedAt:            time.Now().UTC(),
	}

	dataDetail = entity.ProgramDonasiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r CreateProgramDonasiNews) ToEntity() (data entity.ProgramDonasiNewsEntity) {
	data = entity.ProgramDonasiNewsEntity{

		IDPPCPProgramDonasi:     r.IDPPCPProgramDonasi,
		Title:                   r.Title,
		SubmitAt:                r.SubmitAt,
		DisbursementBalance:     r.DisbursementBalance,
		DisbursementAccount:     r.DisbursementAccount,
		DisbursementBankName:    r.DisbursementBankName,
		DisbursementName:        r.DisbursementName,
		DisbursementDescription: r.DisbursementDescription,
		IsDeleted:               false,
		CreatedAt:               time.Now().UTC(),
	}

	return
}

func (r UpdateProgramDonasi) ToEntity() (data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity) {
	validFrom := r.ValidFrom.UTC()
	validTo := r.ValidTo.UTC()

	var nominalStr strings.Builder
	for _, nominal := range r.Nominal {
		nominalString := strconv.Itoa(nominal)
		nominalStr.WriteString(nominalString + "|")
	}
	nominalValue := strings.TrimSuffix(nominalStr.String(), "|")

	data = entity.ProgramDonasiEntity{
		KategoriID:           r.KategoriID,
		Title:                r.Title,
		SubTitle:             r.SubTitle,
		Tag:                  r.Tag,
		SEOURL:               r.SEOURL,
		ThumbnailImageURL:    r.ThumbnailImageURL,
		IDPPCPPenggalangDana: r.IDPPCPPenggalangDana,
		ValidFrom:            &validFrom,
		ValidTo:              &validTo,
		Nominal:              nominalValue,
		Target:               r.Target,
		KitaBisaLink:         r.KitaBisaLink,
		AyoBantuLink:         r.AyoBantuLink,
		Description:          r.Description,
		IDPPCPMasterQris:     &r.IDPPCPMasterQris,
		QrisImageURL:         &r.QrisImageURL,
		IsShow:               *r.IsShow,
	}

	dataDetail = entity.ProgramDonasiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r UpdateProgramDonasiNews) ToEntity() (data entity.ProgramDonasiNewsEntity) {
	data = entity.ProgramDonasiNewsEntity{
		Title:                   r.Title,
		SubmitAt:                r.SubmitAt,
		DisbursementBalance:     r.DisbursementBalance,
		DisbursementBankName:    r.DisbursementBankName,
		DisbursementAccount:     r.DisbursementAccount,
		DisbursementName:        r.DisbursementName,
		DisbursementDescription: r.DisbursementDescription,
	}

	return
}

func (r ProgramDonasiQuery) ToEntity() (data entity.ProgramDonasiQueryEntity) {
	filters := []entity.ProgramDonasiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramDonasiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramDonasiQueryEntity{
		Limit:         r.Limit,
		Offset:        r.Offset,
		Filter:        filters,
		Order:         r.Order,
		Sort:          r.Sort,
		CreatedAtFrom: r.CreatedAtFrom,
		CreatedAtTo:   r.CreatedAtTo,
		PublishAtFrom: r.PublishAtFrom,
		PublishAtTo:   r.PublishAtTo,
	}
	return
}

func ToPayload(data entity.ProgramDonasiEntity) (response ReadProgramDonasi) {
	nominalSplit := strings.Split(data.Nominal, "|")

	var nominalValue []int
	for _, nominalString := range nominalSplit {
		nominal, _ := strconv.Atoi(nominalString)
		nominalValue = append(nominalValue, nominal)
	}

	response = ReadProgramDonasi{
		IDPPCPProgramDonasi:  data.IDPPCPProgramDonasi,
		KategoriID:           data.KategoriID,
		KategoriName:         data.KategoriName,
		Title:                data.Title,
		SubTitle:             data.SubTitle,
		Content:              data.Detail.Content,
		Tag:                  data.Tag,
		ThumbnailImageURL:    data.ThumbnailImageURL,
		IDPPCPPenggalangDana: data.IDPPCPPenggalangDana,
		SEOURL:               data.SEOURL,
		ValidFrom:            data.ValidFrom,
		ValidTo:              data.ValidTo,
		Nominal:              nominalValue,
		Target:               data.Target,
		DonationCollect:      data.DonationCollect,
		KitaBisaLink:         data.KitaBisaLink,
		AyoBantuLink:         data.AyoBantuLink,
		IDPPCPMasterQris:     data.IDPPCPMasterQris,
		QrisImageURL:         data.QrisImageURL,
		Description:          data.Description,
		Status:               data.Status,
		CreatedAt:            data.CreatedAt,
		CreatedBy:            data.CreatedBy,
		UpdatedAt:            data.UpdatedAt,
		UpdatedBy:            data.UpdatedBy,
		PublishedAt:          data.PublishedAt,
		PublishedBy:          data.PublishedBy,
		IsDeleted:            data.IsDeleted,
		IsShow:               data.IsShow,

		PenggalangDana: penggalangDanaApp.ToPayload(data.PenggalangDana),
	}
	return
}

func ToPayloadNews(data entity.ProgramDonasiNewsEntity) (response ReadProgramDonasiNews) {
	response = ReadProgramDonasiNews{
		ID:                      data.ID,
		IDPPCPProgramDonasi:     data.IDPPCPProgramDonasi,
		Title:                   data.Title,
		SubmitAt:                data.SubmitAt,
		DisbursementBalance:     data.DisbursementBalance,
		DisbursementAccount:     data.DisbursementAccount,
		DisbursementBankName:    data.DisbursementBankName,
		DisbursementName:        data.DisbursementName,
		DisbursementDescription: data.DisbursementDescription,
		IsDeleted:               data.IsDeleted,
		CreatedAt:               data.CreatedAt,
		CreatedBy:               data.CreatedBy,
		UpdatedAt:               data.UpdatedAt,
		UpdatedBy:               data.UpdatedBy,
	}
	return
}
