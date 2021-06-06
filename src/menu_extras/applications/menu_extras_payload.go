package applications

import (
	"encoding/json"
	"pemuda-peduli/src/menu_extras/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateMenuExtras struct {
	Title string `json:"title" valid:"required"`
	Link  string `json:"link" valid:"required,url"`
}

type UpdateMenuExtras struct {
	Title string `json:"title" valid:"required"`
	Link  string `json:"link" valid:"required,url"`
}

type MenuExtrasQuery struct {
	Limit         string                  `json:"limit" valid:"required"`
	Offset        string                  `json:"offset" valid:"required"`
	Filter        []MenuExtrasFilterQuery `json:"filters"`
	Order         string                  `json:"order"`
	Sort          string                  `json:"sort"`
	CreatedAtFrom string                  `json:"created_at_from"`
	CreatedAtTo   string                  `json:"created_at_to"`
	PublishAtFrom string                  `json:"publish_at_from"`
	PublishAtTo   string                  `json:"publish_at_to"`
}

type MenuExtrasFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadMenuExtras struct {
	IDPPCPMenuExtras string     `json:"id"`
	Title            string     `json:"title"`
	Link             string     `json:"link"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        *string    `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        *string    `json:"updated_by"`
	PublishedAt      *time.Time `json:"published_at"`
	PublishedBy      *string    `json:"published_by"`
	IsDeleted        bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateMenuExtras, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateMenuExtras, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload MenuExtrasQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateMenuExtras) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateMenuExtras) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r MenuExtrasQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateMenuExtras) ToEntity() (data entity.MenuExtrasEntity) {
	data = entity.MenuExtrasEntity{
		Title:     r.Title,
		Link:      r.Link,
		CreatedAt: time.Now(),
	}
	return
}

func (r UpdateMenuExtras) ToEntity() (data entity.MenuExtrasEntity) {
	data = entity.MenuExtrasEntity{
		Title: r.Title,
		Link:  r.Link,
	}
	return
}

func (r MenuExtrasQuery) ToEntity() (data entity.MenuExtrasQueryEntity) {
	filters := []entity.MenuExtrasFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.MenuExtrasFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.MenuExtrasQueryEntity{
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

func ToPayload(data entity.MenuExtrasEntity) (response ReadMenuExtras) {
	response = ReadMenuExtras{
		IDPPCPMenuExtras: data.IDPPCPMenuExtras,
		Title:            data.Title,
		Link:             data.Link,
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
