package applications

import (
	"encoding/json"
	"pemuda-peduli/src/tentang_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTentangKami struct {
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type UpdateTentangKami struct {
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type TentangKamiQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []TentangKamiFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type TentangKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadTentangKami struct {
	IDPPCPTentangKami string     `json:"id"`
	Title             string     `json:"title"`
	SubTitle          string     `json:"sub_title"`
	Tag               string     `json:"tag"`
	ThumbnailImageURL string     `json:"thumbnail_image_url"`
	Description       string     `json:"description"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateTentangKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateTentangKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload TentangKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateTentangKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateTentangKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r TentangKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateTentangKami) ToEntity() (data entity.TentangKamiEntity) {
	data = entity.TentangKamiEntity{
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
		CreatedAt:         time.Now(),
	}
	return
}

func (r UpdateTentangKami) ToEntity() (data entity.TentangKamiEntity) {
	data = entity.TentangKamiEntity{
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
	}
	return
}

func (r TentangKamiQuery) ToEntity() (data entity.TentangKamiQueryEntity) {
	filters := []entity.TentangKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.TentangKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.TentangKamiQueryEntity{
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

func ToPayload(data entity.TentangKamiEntity) (response ReadTentangKami) {
	response = ReadTentangKami{
		IDPPCPTentangKami: data.IDPPCPTentangKami,
		ThumbnailImageURL: data.ThumbnailImageURL,
		Description:       data.Description,
		Status:            data.Status,
		CreatedAt:         data.CreatedAt,
		CreatedBy:         data.CreatedBy,
		UpdatedAt:         data.UpdatedAt,
		UpdatedBy:         data.UpdatedBy,
		PublishedAt:       data.PublishedAt,
		PublishedBy:       data.PublishedBy,
		IsDeleted:         data.IsDeleted,
	}
	return
}
