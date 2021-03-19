package applications

import (
	"encoding/json"
	"pemuda-peduli/src/tujuan_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTujuanKami struct {
	Title       string `json:"title" valid:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon" valid:"required,url"`
}

type UpdateTujuanKami struct {
	Title       string `json:"title" valid:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon" valid:"required,url"`
}

type TujuanKamiQuery struct {
	Limit         string                  `json:"limit" valid:"required"`
	Offset        string                  `json:"offset" valid:"required"`
	Filter        []TujuanKamiFilterQuery `json:"filters"`
	Order         string                  `json:"order"`
	Sort          string                  `json:"sort"`
	CreatedAtFrom string                  `json:"created_at_from"`
	CreatedAtTo   string                  `json:"created_at_to"`
	PublishAtFrom string                  `json:"publish_at_from"`
	PublishAtTo   string                  `json:"publish_at_to"`
}

type TujuanKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadTujuanKami struct {
	IDPPCPTujuanKami string     `json:"id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Icon             string     `json:"icon"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        *string    `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        *string    `json:"updated_by"`
	PublishedAt      *time.Time `json:"published_at"`
	PublishedBy      *string    `json:"published_by"`
	IsDeleted        bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateTujuanKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateTujuanKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload TujuanKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateTujuanKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateTujuanKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r TujuanKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateTujuanKami) ToEntity() (data entity.TujuanKamiEntity) {
	data = entity.TujuanKamiEntity{
		Title:       r.Title,
		Description: r.Description,
		Icon:        r.Icon,
		CreatedAt:   time.Now(),
	}
	return
}

func (r UpdateTujuanKami) ToEntity() (data entity.TujuanKamiEntity) {
	data = entity.TujuanKamiEntity{
		Title:       r.Title,
		Description: r.Description,
		Icon:        r.Icon,
	}
	return
}

func (r TujuanKamiQuery) ToEntity() (data entity.TujuanKamiQueryEntity) {
	filters := []entity.TujuanKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.TujuanKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.TujuanKamiQueryEntity{
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

func ToPayload(data entity.TujuanKamiEntity) (response ReadTujuanKami) {
	response = ReadTujuanKami{
		IDPPCPTujuanKami: data.IDPPCPTujuanKami,
		Title:            data.Title,
		Description:      data.Description,
		Icon:             data.Icon,
		Status:           data.Status,
		CreatedAt:        data.CreatedAt,
		CreatedBy:        data.CreatedBy,
		UpdatedAt:        data.UpdatedAt,
		UpdatedBy:        data.UpdatedBy,
		PublishedAt:      data.PublishedAt,
		PublishedBy:      data.PublishedBy,
		IsDeleted:        data.IsDeleted,
	}
	return
}
