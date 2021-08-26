package applications

import (
	"context"
	"encoding/json"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreatePenggalangDana struct {
	Name              string `json:"name" valid:"required"`
	Description       string `json:"description"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type UpdatePenggalangDana struct {
	Name              string `json:"name" valid:"required"`
	Description       string `json:"description"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
}

type PenggalangDanaQuery struct {
	Limit         string                      `json:"limit" valid:"required"`
	Offset        string                      `json:"offset" valid:"required"`
	Filter        []PenggalangDanaFilterQuery `json:"filters"`
	Order         string                      `json:"order"`
	Sort          string                      `json:"sort"`
	CreatedAtFrom string                      `json:"created_at_from"`
	CreatedAtTo   string                      `json:"created_at_to"`
	PublishAtFrom string                      `json:"publish_at_from"`
	PublishAtTo   string                      `json:"publish_at_to"`
}

type PenggalangDanaFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadPenggalangDana struct {
	IDPPCPPenggalangDana string     `db:"id"`
	Name                 string     `db:"name"`
	Description          string     `db:"description"`
	ThumbnailImageURL    string     `db:"thumbnail_image_url"`
	IsVerified           bool       `db:"is_verified"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at"`
	IsDeleted            bool       `db:"is_deleted"`
}

func GetCreatePayload(body []byte) (payload CreatePenggalangDana, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdatePenggalangDana, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload PenggalangDanaQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreatePenggalangDana) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdatePenggalangDana) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r PenggalangDanaQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreatePenggalangDana) ToEntity() (data entity.PenggalangDanaEntity) {
	data = entity.PenggalangDanaEntity{
		Name:              r.Name,
		Description:       r.Description,
		ThumbnailImageURL: r.ThumbnailImageURL,
		IsVerified:        false,
		CreatedAt:         time.Now().UTC(),
		IsDeleted:         false,
	}
	return
}

func (r UpdatePenggalangDana) ToEntity(ctx context.Context) (data entity.PenggalangDanaEntity) {
	currentTime := time.Now().UTC()
	data = entity.PenggalangDanaEntity{
		IDPPCPPenggalangDana: ctx.Value("id").(string),
		Name:                 r.Name,
		Description:          r.Description,
		ThumbnailImageURL:    r.ThumbnailImageURL,
		UpdatedAt:            &currentTime,
	}
	return
}

func (r PenggalangDanaQuery) ToEntity() (data entity.PenggalangDanaQueryEntity) {
	filters := []entity.PenggalangDanaFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.PenggalangDanaFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.PenggalangDanaQueryEntity{
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

func ToPayload(data entity.PenggalangDanaEntity) (response ReadPenggalangDana) {
	response = ReadPenggalangDana{
		IDPPCPPenggalangDana: data.IDPPCPPenggalangDana,
		Name:                 data.Name,
		Description:          data.Description,
		ThumbnailImageURL:    data.ThumbnailImageURL,
		IsVerified:           data.IsVerified,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
		IsDeleted:            data.IsDeleted,
	}
	return
}
