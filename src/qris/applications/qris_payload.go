package applications

import (
	"encoding/json"
	"pemuda-peduli/src/qris/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateQris struct {
	Title             string `json:"title" valid:"required"`
	Description       string `json:"description"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"required,url"`
}

type UpdateQris struct {
	Title             string `json:"title" valid:"required"`
	Description       string `json:"description"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"required,url"`
}

type QrisQuery struct {
	Limit         string            `json:"limit" valid:"required"`
	Offset        string            `json:"offset" valid:"required"`
	Filter        []QrisFilterQuery `json:"filters"`
	Order         string            `json:"order"`
	Sort          string            `json:"sort"`
	CreatedAtFrom string            `json:"created_at_from"`
	CreatedAtTo   string            `json:"created_at_to"`
}

type QrisFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadQris struct {
	IDPPCPQris        string     `json:"id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	ThumbnailImageURL string     `json:"thumbnail_image_url"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
}

func GetCreatePayload(body []byte) (payload CreateQris, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateQris, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload QrisQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateQris) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateQris) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r QrisQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateQris) ToEntity() (data entity.QrisEntity) {
	data = entity.QrisEntity{
		Title:             r.Title,
		Description:       r.Description,
		ThumbnailImageURL: r.ThumbnailImageURL,
	}
	return
}

func (r UpdateQris) ToEntity() (data entity.QrisEntity) {
	data = entity.QrisEntity{
		Title:             r.Title,
		Description:       r.Description,
		ThumbnailImageURL: r.ThumbnailImageURL,
	}
	return
}

func (r QrisQuery) ToEntity() (data entity.QrisQueryEntity) {
	filters := []entity.QrisFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.QrisFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.QrisQueryEntity{
		Limit:         r.Limit,
		Offset:        r.Offset,
		Filter:        filters,
		Order:         r.Order,
		Sort:          r.Sort,
		CreatedAtFrom: r.CreatedAtFrom,
		CreatedAtTo:   r.CreatedAtTo,
	}
	return
}

func ToPayload(data entity.QrisEntity) (response ReadQris) {
	response = ReadQris{
		IDPPCPQris:        data.IDPPCPQris,
		Title:             data.Title,
		Description:       data.Description,
		ThumbnailImageURL: data.ThumbnailImageURL,
		Status:            data.Status,
		CreatedAt:         data.CreatedAt,
		CreatedBy:         data.CreatedBy,
		UpdatedAt:         data.UpdatedAt,
		UpdatedBy:         data.UpdatedBy,
	}
	return
}
