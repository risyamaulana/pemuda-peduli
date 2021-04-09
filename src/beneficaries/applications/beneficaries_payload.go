package applications

import (
	"encoding/json"
	"pemuda-peduli/src/beneficaries/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateBeneficaries struct {
	Title             string `json:"title" valid:"required"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	DeeplinkRight     string `json:"deeplink_right" valid:"url"`
	DeeplinkLeft      string `json:"deeplink_left" valid:"url"`
	Description       string `json:"description"`
}

type UpdateBeneficaries struct {
	Title             string `json:"title" valid:"required"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	DeeplinkRight     string `json:"deeplink_right" valid:"url"`
	DeeplinkLeft      string `json:"deeplink_left" valid:"url"`
	Description       string `json:"description"`
}

type BeneficariesQuery struct {
	Limit         string                    `json:"limit" valid:"required"`
	Offset        string                    `json:"offset" valid:"required"`
	Filter        []BeneficariesFilterQuery `json:"filters"`
	Order         string                    `json:"order"`
	Sort          string                    `json:"sort"`
	CreatedAtFrom string                    `json:"created_at_from"`
	CreatedAtTo   string                    `json:"created_at_to"`
	PublishAtFrom string                    `json:"publish_at_from"`
	PublishAtTo   string                    `json:"publish_at_to"`
}

type BeneficariesFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadBeneficaries struct {
	IDPPCPBeneficaries string     `json:"id"`
	Title              string     `json:"title"`
	ThumbnailImageURL  string     `json:"thumbnail_image_url"`
	DeeplinkRight      string     `json:"deeplink_right"`
	DeeplinkLeft       string     `json:"deeplink_left"`
	Description        string     `json:"description"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"created_at"`
	CreatedBy          *string    `json:"created_by"`
	UpdatedAt          *time.Time `json:"updated_at"`
	UpdatedBy          *string    `json:"updated_by"`
	PublishedAt        *time.Time `json:"published_at"`
	PublishedBy        *string    `json:"published_by"`
	IsDeleted          bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateBeneficaries, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateBeneficaries, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload BeneficariesQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateBeneficaries) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateBeneficaries) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r BeneficariesQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateBeneficaries) ToEntity() (data entity.BeneficariesEntity) {
	data = entity.BeneficariesEntity{
		Title:             r.Title,
		ThumbnailImageURL: r.ThumbnailImageURL,
		DeeplinkRight:     r.DeeplinkRight,
		DeeplinkLeft:      r.DeeplinkLeft,
		Description:       r.Description,
		CreatedAt:         time.Now(),
	}
	return
}

func (r UpdateBeneficaries) ToEntity() (data entity.BeneficariesEntity) {
	data = entity.BeneficariesEntity{
		Title:             r.Title,
		ThumbnailImageURL: r.ThumbnailImageURL,
		DeeplinkRight:     r.DeeplinkRight,
		DeeplinkLeft:      r.DeeplinkLeft,
		Description:       r.Description,
	}
	return
}

func (r BeneficariesQuery) ToEntity() (data entity.BeneficariesQueryEntity) {
	filters := []entity.BeneficariesFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.BeneficariesFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.BeneficariesQueryEntity{
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

func ToPayload(data entity.BeneficariesEntity) (response ReadBeneficaries) {
	response = ReadBeneficaries{
		IDPPCPBeneficaries: data.IDPPCPBeneficaries,
		Title:              data.Title,
		ThumbnailImageURL:  data.ThumbnailImageURL,
		DeeplinkRight:      data.DeeplinkRight,
		DeeplinkLeft:       data.DeeplinkLeft,
		Description:        data.Description,
		Status:             data.Status,
		CreatedAt:          data.CreatedAt,
		CreatedBy:          data.CreatedBy,
		UpdatedAt:          data.UpdatedAt,
		UpdatedBy:          data.UpdatedBy,
		PublishedAt:        data.PublishedAt,
		PublishedBy:        data.PublishedBy,
		IsDeleted:          data.IsDeleted,
	}
	return
}
