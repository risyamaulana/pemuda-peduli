package applications

import (
	"encoding/json"
	"pemuda-peduli/src/kontak_kami/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateKontakKami struct {
	SKLegalitas string `json:"sk_legalitas" valid:"required,url"`
	Address     string `json:"address" valid:"required"`
	ContactName string `json:"contact_name" valid:"required"`
	Icon        string `json:"icon" valid:"required,url"`
	ContactLink string `json:"contact_link" valid:"required,url"`
	Menu        string `json:"menu"`
}

type UpdateKontakKami struct {
	SKLegalitas string `json:"sk_legalitas" valid:"required,url"`
	Address     string `json:"address" valid:"required"`
	ContactName string `json:"contact_name" valid:"required"`
	Icon        string `json:"icon" valid:"required,url"`
	ContactLink string `json:"contact_link" valid:"required,url"`
	Menu        string `json:"menu"`
}

type KontakKamiQuery struct {
	Limit         string                  `json:"limit" valid:"required"`
	Offset        string                  `json:"offset" valid:"required"`
	Filter        []KontakKamiFilterQuery `json:"filters"`
	Order         string                  `json:"order"`
	Sort          string                  `json:"sort"`
	CreatedAtFrom string                  `json:"created_at_from"`
	CreatedAtTo   string                  `json:"created_at_to"`
	PublishAtFrom string                  `json:"publish_at_from"`
	PublishAtTo   string                  `json:"publish_at_to"`
}

type KontakKamiFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadKontakKami struct {
	IDPPCPKontakKami string     `json:"id"`
	SKLegalitas      string     `json:"sk_legalitas"`
	Address          string     `json:"address"`
	ContactName      string     `json:"contact_name"`
	Icon             string     `json:"icon"`
	ContactLink      string     `json:"contact_link"`
	Menu             string     `json:"menu"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        *string    `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        *string    `json:"updated_by"`
	PublishedAt      *time.Time `json:"published_at"`
	PublishedBy      *string    `json:"published_by"`
	IsDeleted        bool       `json:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreateKontakKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateKontakKami, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload KontakKamiQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateKontakKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateKontakKami) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r KontakKamiQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateKontakKami) ToEntity() (data entity.KontakKamiEntity) {
	data = entity.KontakKamiEntity{
		SKLegalitas: r.SKLegalitas,
		Address:     r.Address,
		ContactName: r.ContactName,
		Icon:        r.Icon,
		ContactLink: r.ContactLink,
		Menu:        r.Menu,
		CreatedAt:   time.Now(),
	}
	return
}

func (r UpdateKontakKami) ToEntity() (data entity.KontakKamiEntity) {
	data = entity.KontakKamiEntity{
		SKLegalitas: r.SKLegalitas,
		Address:     r.Address,
		ContactName: r.ContactName,
		Icon:        r.Icon,
		ContactLink: r.ContactLink,
		Menu:        r.Menu,
	}
	return
}

func (r KontakKamiQuery) ToEntity() (data entity.KontakKamiQueryEntity) {
	filters := []entity.KontakKamiFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.KontakKamiFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.KontakKamiQueryEntity{
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

func ToPayload(data entity.KontakKamiEntity) (response ReadKontakKami) {
	response = ReadKontakKami{
		IDPPCPKontakKami: data.IDPPCPKontakKami,
		SKLegalitas:      data.SKLegalitas,
		Address:          data.Address,
		ContactName:      data.ContactName,
		Icon:             data.Icon,
		ContactLink:      data.ContactLink,
		Menu:             data.Menu,
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
