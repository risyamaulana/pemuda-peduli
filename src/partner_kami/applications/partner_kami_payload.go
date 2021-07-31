package applications

import (
	"encoding/json"
	"pemuda-peduli/src/partner_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreatePartnerKami struct {
	Name              string `json:"name" valid:"required"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type UpdatePartnerKami struct {
	Name              string `json:"name" valid:"required"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type PartnerKamiQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []PartnerKamiFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type PartnerKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadPartnerKami struct {
	IDPPCPPartnerKami string     `json:"id"`
	Name              string     `json:"name"`
	ThumbnailImageURL string     `json:"thumbnail_image_url"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreatePartnerKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdatePartnerKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload PartnerKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreatePartnerKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdatePartnerKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r PartnerKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreatePartnerKami) ToEntity() (data entity.PartnerKamiEntity) {
	data = entity.PartnerKamiEntity{
		Name:              r.Name,
		ThumbnailImageURL: r.ThumbnailImageURL,
		CreatedAt:         time.Now().UTC(),
	}
	return
}

func (r UpdatePartnerKami) ToEntity() (data entity.PartnerKamiEntity) {
	data = entity.PartnerKamiEntity{
		Name:              r.Name,
		ThumbnailImageURL: r.ThumbnailImageURL,
	}
	return
}

func (r PartnerKamiQuery) ToEntity() (data entity.PartnerKamiQueryEntity) {
	filters := []entity.PartnerKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.PartnerKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.PartnerKamiQueryEntity{
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

func ToPayload(data entity.PartnerKamiEntity) (response ReadPartnerKami) {
	response = ReadPartnerKami{
		IDPPCPPartnerKami: data.IDPPCPPartnerKami,
		Name:              data.Name,
		ThumbnailImageURL: data.ThumbnailImageURL,
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
