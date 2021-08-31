package applications

import (
	"context"
	"encoding/json"
	"errors"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramDonasiRutin struct {
	IDPPCPProgramDonasiKategori string `json:"id_kategori" valid:"required"`

	Title    string `json:"title" valid:"required"`
	SubTitle string `json:"sub_title" valid:"required"`
	Content  string `json:"content" valid:"required"`

	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	SEOURL            string `json:"seo_url"`

	IDPPCPMasterQris string `json:"id_pp_cp_master_qris"`
	QrisImageURL     string `json:"qris_image_url"`

	Description string `json:"description"`
}

type CreateProgramDonasiRutinPaket struct {
	PaketName     string  `json:"paket_name" valid:"required"`
	Benefit       string  `json:"benefit" valid:"required"`
	Nominal       float64 `json:"nominal" valid:"required"`
	PaketImageURL string  `json:"paket_image_url" valid:"url"`
}

type UpdateProgramDonasiRutin struct {
	IDPPCPProgramDonasiKategori string `json:"id_kategori" valid:"required"`

	Title    string `json:"title" valid:"required"`
	SubTitle string `json:"sub_title" valid:"required"`
	Content  string `json:"content" valid:"required"`

	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	IsShow            *bool  `json:"is_show"`
	SEOURL            string `json:"seo_url"`

	IDPPCPMasterQris string `json:"id_pp_cp_master_qris"`
	QrisImageURL     string `json:"qris_image_url"`

	Description string `json:"description"`
}

type UpdateProgramDonasiRutinPaket struct {
	PaketName     string  `json:"paket_name" valid:"required"`
	Benefit       string  `json:"benefit" valid:"required"`
	Nominal       float64 `json:"nominal" valid:"required"`
	PaketImageURL string  `json:"paket_image_url" valid:"url"`
}

type ProgramDonasiRutinQuery struct {
	Limit         string                          `json:"limit" valid:"required"`
	Offset        string                          `json:"offset" valid:"required"`
	Filter        []ProgramDonasiRutinFilterQuery `json:"filters"`
	Order         string                          `json:"order"`
	Sort          string                          `json:"sort"`
	CreatedAtFrom string                          `json:"created_at_from"`
	CreatedAtTo   string                          `json:"created_at_to"`
	PublishAtFrom string                          `json:"publish_at_from"`
	PublishAtTo   string                          `json:"publish_at_to"`
}

type ProgramDonasiRutinFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadProgramDonasiRutin struct {
	IDPPCPProgramDonasiRutin    string     `json:"id"`
	IDPPCPProgramDonasiKategori string     `json:"id_kategori"`
	KategoriName                string     `json:"kategori_name"`
	Title                       string     `json:"title"`
	SubTitle                    string     `json:"sub_title"`
	Content                     string     `json:"content"`
	Benefit                     string     `json:"benefit"`
	IDPPCPMasterQris            *string    `json:"id_pp_cp_master_qris"`
	QrisImageURL                *string    `json:"qris_image_url"`
	Tag                         string     `json:"tag"`
	ThumbnailImageURL           string     `json:"thumbnail_image_url"`
	ValidFrom                   *time.Time `json:"valid_from"`
	ValidTo                     *time.Time `json:"valid_to"`
	Target                      *float64   `json:"target"`
	DonationCollect             float64    `json:"donation_collect"`
	Description                 string     `json:"description"`
	SEOURL                      string     `json:"seo_url"`
	Status                      string     `json:"status"`
	CreatedAt                   time.Time  `json:"created_at"`
	CreatedBy                   *string    `json:"created_by"`
	UpdatedAt                   *time.Time `json:"updated_at"`
	UpdatedBy                   *string    `json:"updated_by"`
	PublishedAt                 *time.Time `json:"published_at"`
	PublishedBy                 *string    `json:"published_by"`
	IsDeleted                   bool       `json:"is_deleted"`
	IsShow                      bool       `json:"is_show"`
}

type ReadProgramDonasiRutinPaket struct {
	IDPPCPProgramDonasiRutinPaket string     `json:"id"`
	IDPPCPProgramDonasiRutin      string     `json:"id_pp_cp_program_donasi_rutin"`
	PaketName                     string     `json:"paket_name"`
	Benefit                       string     `json:"benefit"`
	Nominal                       float64    `json:"nominal"`
	PaketImageURL                 string     `json:"paket_image_url"`
	CreatedAt                     time.Time  `json:"created_at"`
	UpdatedAt                     *time.Time `json:"updated_at"`
	IsDeleted                     bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateProgramDonasiRutin, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetCreatePaketPayload(body []byte) (payload CreateProgramDonasiRutinPaket, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateProgramDonasiRutin, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePaketPayload(body []byte) (payload UpdateProgramDonasiRutinPaket, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload ProgramDonasiRutinQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateProgramDonasiRutin) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiRutinPaket) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateProgramDonasiRutin) Validate() (err error) {
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

func (r UpdateProgramDonasiRutinPaket) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	return
}

func (r ProgramDonasiRutinQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateProgramDonasiRutin) ToEntity() (data entity.ProgramDonasiRutinEntity) {
	data = entity.ProgramDonasiRutinEntity{
		IDPPCPProgramDonasiKategori: r.IDPPCPProgramDonasiKategori,
		Title:                       r.Title,
		SubTitle:                    r.SubTitle,
		Content:                     r.Content,
		Tag:                         r.Tag,
		ThumbnailImageURL:           r.ThumbnailImageURL,
		Description:                 r.Description,
		SEOURL:                      r.SEOURL,

		IDPPCPMasterQris: &r.IDPPCPMasterQris,
		QrisImageURL:     &r.QrisImageURL,
		CreatedAt:        time.Now().UTC(),
	}
	return
}

func (r CreateProgramDonasiRutinPaket) ToEntity(ctx context.Context) (data entity.ProgramDonasiRutinPaketEntity) {
	data = entity.ProgramDonasiRutinPaketEntity{
		IDPPCPProgramDonasiRutin: ctx.Value("id").(string),
		PaketName:                r.PaketName,
		Benefit:                  r.Benefit,
		Nominal:                  r.Nominal,
		PaketImageURL:            r.PaketImageURL,
		CreatedAt:                time.Now().UTC(),
		IsDeleted:                false,
	}
	return
}

func (r UpdateProgramDonasiRutin) ToEntity() (data entity.ProgramDonasiRutinEntity) {
	data = entity.ProgramDonasiRutinEntity{
		IDPPCPProgramDonasiKategori: r.IDPPCPProgramDonasiKategori,
		Title:                       r.Title,
		SubTitle:                    r.SubTitle,
		Content:                     r.Content,
		Tag:                         r.Tag,
		ThumbnailImageURL:           r.ThumbnailImageURL,
		Description:                 r.Description,
		IDPPCPMasterQris:            &r.IDPPCPMasterQris,
		QrisImageURL:                &r.QrisImageURL,
		SEOURL:                      r.SEOURL,
		IsShow:                      *r.IsShow,
	}
	return
}

func (r UpdateProgramDonasiRutinPaket) ToEntity() (data entity.ProgramDonasiRutinPaketEntity) {
	currentTime := time.Now().UTC()
	data = entity.ProgramDonasiRutinPaketEntity{
		PaketName:     r.PaketName,
		Benefit:       r.Benefit,
		Nominal:       r.Nominal,
		PaketImageURL: r.PaketImageURL,
		UpdatedAt:     &currentTime,
		IsDeleted:     false,
	}
	return
}

func (r ProgramDonasiRutinQuery) ToEntity() (data entity.ProgramDonasiRutinQueryEntity) {
	filters := []entity.ProgramDonasiRutinFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.ProgramDonasiRutinFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.ProgramDonasiRutinQueryEntity{
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

func ToPayload(data entity.ProgramDonasiRutinEntity) (response ReadProgramDonasiRutin) {
	response = ReadProgramDonasiRutin{
		IDPPCPProgramDonasiRutin:    data.IDPPCPProgramDonasiRutin,
		IDPPCPProgramDonasiKategori: data.IDPPCPProgramDonasiKategori,
		KategoriName:                data.KategoriName,
		Title:                       data.Title,
		SubTitle:                    data.SubTitle,
		Content:                     data.Content,
		Tag:                         data.Tag,
		ThumbnailImageURL:           data.ThumbnailImageURL,
		IDPPCPMasterQris:            data.IDPPCPMasterQris,
		DonationCollect:             data.DonationCollect,
		QrisImageURL:                data.QrisImageURL,
		Description:                 data.Description,
		SEOURL:                      data.SEOURL,
		Status:                      data.Status,
		CreatedAt:                   data.CreatedAt,
		CreatedBy:                   data.CreatedBy,
		UpdatedAt:                   data.UpdatedAt,
		UpdatedBy:                   data.UpdatedBy,
		PublishedAt:                 data.PublishedAt,
		PublishedBy:                 data.PublishedBy,
		IsDeleted:                   data.IsDeleted,
		IsShow:                      data.IsShow,
	}
	return
}

func ToPayloadPaket(data entity.ProgramDonasiRutinPaketEntity) (response ReadProgramDonasiRutinPaket) {
	response = ReadProgramDonasiRutinPaket{
		IDPPCPProgramDonasiRutinPaket: data.IDPPCPProgramDonasiRutinPaket,
		IDPPCPProgramDonasiRutin:      data.IDPPCPProgramDonasiRutin,
		PaketName:                     data.PaketName,
		Benefit:                       data.Benefit,
		Nominal:                       data.Nominal,
		PaketImageURL:                 data.PaketImageURL,
		CreatedAt:                     data.CreatedAt,
		UpdatedAt:                     data.UpdatedAt,
		IsDeleted:                     data.IsDeleted,
	}

	return
}
