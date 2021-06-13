package applications

import (
	"encoding/json"
	"pemuda-peduli/src/hubungi_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateHubungiKami struct {
	Icon  string `json:"icon" valid:"required,url"`
	Link  string `json:"link" valid:"required,url"`
	Title string `json:"title" valid:"required"`
}

type UpdateHubungiKami struct {
	Icon  string `json:"icon" valid:"required,url"`
	Link  string `json:"link" valid:"required,url"`
	Title string `json:"title" valid:"required"`
}

type HubungiKamiQuery struct {
	Limit         string                   `json:"limit" valid:"required"`
	Offset        string                   `json:"offset" valid:"required"`
	Filter        []HubungiKamiFilterQuery `json:"filters"`
	Order         string                   `json:"order"`
	Sort          string                   `json:"sort"`
	CreatedAtFrom string                   `json:"created_at_from"`
	CreatedAtTo   string                   `json:"created_at_to"`
	PublishAtFrom string                   `json:"publish_at_from"`
	PublishAtTo   string                   `json:"publish_at_to"`
}

type HubungiKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadHubungiKami struct {
	IDPPCPHubungiKami string     `json:"id"`
	Icon              string     `json:"icon"`
	Link              string     `json:"link"`
	Title             string     `json:"title"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         *string    `json:"created_by"`
	UpdatedAt         *time.Time `json:"updated_at"`
	UpdatedBy         *string    `json:"updated_by"`
	PublishedAt       *time.Time `json:"published_at"`
	PublishedBy       *string    `json:"published_by"`
	IsDeleted         bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateHubungiKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateHubungiKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload HubungiKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateHubungiKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateHubungiKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r HubungiKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateHubungiKami) ToEntity() (data entity.HubungiKamiEntity) {
	data = entity.HubungiKamiEntity{
		Icon:      r.Icon,
		Link:      r.Link,
		Title:     r.Title,
		CreatedAt: time.Now(),
	}
	return
}

func (r UpdateHubungiKami) ToEntity() (data entity.HubungiKamiEntity) {
	data = entity.HubungiKamiEntity{
		Icon:  r.Icon,
		Link:  r.Link,
		Title: r.Title,
	}
	return
}

func (r HubungiKamiQuery) ToEntity() (data entity.HubungiKamiQueryEntity) {
	filters := []entity.HubungiKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.HubungiKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.HubungiKamiQueryEntity{
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

func ToPayload(data entity.HubungiKamiEntity) (response ReadHubungiKami) {
	response = ReadHubungiKami{
		IDPPCPHubungiKami: data.IDPPCPHubungiKami,
		Icon:              data.Icon,
		Link:              data.Link,
		Title:             data.Title,
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
