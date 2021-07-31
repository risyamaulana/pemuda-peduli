package applications

import (
	"encoding/json"
	"errors"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateProgramDonasi struct {
	Title             string     `json:"title" valid:"required"`
	SubTitle          string     `json:"sub_title" valid:"required"`
	Content           string     `json:"content" valid:"required"`
	Tag               string     `json:"tag"`
	ThumbnailImageURL string     `json:"thumbnail_image_url" valid:"url"`
	ValidFrom         *time.Time `json:"valid_from" valid:"required"`
	ValidTo           *time.Time `json:"valid_to" valid:"required"`
	Target            *float64   `json:"target" valid:"required"`
	KitaBisaLink      *string    `json:"kitabisa_link" valid:"url"`
	AyoBantuLink      *string    `json:"ayobantu_link" valid:"url"`
	IDPPCPMasterQris  string     `json:"id_pp_cp_master_qris"`
	QrisImageURL      string     `json:"qris_image_url"`
	Description       string     `json:"description"`
}

type UpdateProgramDonasi struct {
	Title             string     `json:"title" valid:"required"`
	SubTitle          string     `json:"sub_title" valid:"required"`
	Content           string     `json:"content" valid:"required"`
	Tag               string     `json:"tag"`
	ThumbnailImageURL string     `json:"thumbnail_image_url" valid:"url"`
	ValidFrom         *time.Time `json:"valid_from" valid:"required"`
	ValidTo           *time.Time `json:"valid_to" valid:"required"`
	Target            *float64   `json:"target" valid:"required"`
	KitaBisaLink      *string    `json:"kitabisa_link" valid:"url"`
	AyoBantuLink      *string    `json:"ayobantu_link" valid:"url"`
	IDPPCPMasterQris  string     `json:"id_pp_cp_master_qris"`
	QrisImageURL      string     `json:"qris_image_url"`
	Description       string     `json:"description"`
	IsShow            *bool      `json:"is_show"`
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
	IDPPCPProgramDonasi string     `json:"id"`
	Title               string     `json:"title"`
	SubTitle            string     `json:"sub_title"`
	Content             string     `json:"content"`
	Tag                 string     `json:"tag"`
	ThumbnailImageURL   string     `json:"thumbnail_image_url"`
	ValidFrom           *time.Time `json:"valid_from"`
	ValidTo             *time.Time `json:"valid_to"`
	Target              *float64   `json:"target"`
	DonationCollect     float64    `json:"donation_collect"`
	Description         string     `json:"description"`
	Status              string     `json:"status"`
	CreatedAt           time.Time  `json:"created_at"`
	CreatedBy           *string    `json:"created_by"`
	UpdatedAt           *time.Time `json:"updated_at"`
	UpdatedBy           *string    `json:"updated_by"`
	PublishedAt         *time.Time `json:"published_at"`
	PublishedBy         *string    `json:"published_by"`
	IsDeleted           bool       `json:"is_deleted"`
	IsShow              bool       `json:"is_show"`
	KitaBisaLink        *string    `json:"kitabisa_link"`
	AyoBantuLink        *string    `json:"ayobantu_link"`
	IDPPCPMasterQris    *string    `json:"id_pp_cp_master_qris"`
	QrisImageURL        *string    `json:"qris_image_url"`
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

func (r CreateProgramDonasi) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}

	if r.ValidFrom.Before(time.Now()) {
		err = errors.New("Failed : valid from smallest than current time")
		return
	}

	if r.ValidTo.Before(time.Now()) {
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

func (r CreateProgramDonasi) ToEntity() (data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity) {
	data = entity.ProgramDonasiEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		ValidFrom:         r.ValidFrom,
		ValidTo:           r.ValidTo,
		Target:            r.Target,
		KitaBisaLink:      r.KitaBisaLink,
		AyoBantuLink:      r.AyoBantuLink,
		IDPPCPMasterQris:  &r.IDPPCPMasterQris,
		QrisImageURL:      &r.QrisImageURL,
		Description:       r.Description,
		CreatedAt:         time.Now(),
	}

	dataDetail = entity.ProgramDonasiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r UpdateProgramDonasi) ToEntity() (data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity) {
	data = entity.ProgramDonasiEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		ValidFrom:         r.ValidFrom,
		ValidTo:           r.ValidTo,
		Target:            r.Target,
		KitaBisaLink:      r.KitaBisaLink,
		AyoBantuLink:      r.AyoBantuLink,
		Description:       r.Description,
		IDPPCPMasterQris:  &r.IDPPCPMasterQris,
		QrisImageURL:      &r.QrisImageURL,
		IsShow:            *r.IsShow,
	}

	dataDetail = entity.ProgramDonasiDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
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
	response = ReadProgramDonasi{
		IDPPCPProgramDonasi: data.IDPPCPProgramDonasi,
		Title:               data.Title,
		SubTitle:            data.SubTitle,
		Content:             data.Detail.Content,
		Tag:                 data.Tag,
		ThumbnailImageURL:   data.ThumbnailImageURL,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
		Target:              data.Target,
		DonationCollect:     data.DonationCollect,
		KitaBisaLink:        data.KitaBisaLink,
		AyoBantuLink:        data.AyoBantuLink,
		IDPPCPMasterQris:    data.IDPPCPMasterQris,
		QrisImageURL:        data.QrisImageURL,
		Description:         data.Description,
		Status:              data.Status,
		CreatedAt:           data.CreatedAt,
		CreatedBy:           data.CreatedBy,
		UpdatedAt:           data.UpdatedAt,
		UpdatedBy:           data.UpdatedBy,
		PublishedAt:         data.PublishedAt,
		PublishedBy:         data.PublishedBy,
		IsDeleted:           data.IsDeleted,
		IsShow:              data.IsShow,
	}
	return
}
