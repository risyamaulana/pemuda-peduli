package applications

import (
	"encoding/json"
	"pemuda-peduli/src/berita/domain/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateBerita struct {
	Title             string `json:"title" valid:"required,alphanum"`
	SubTitle          string `json:"sub_title" valid:"required,alphanum"`
	Content           string `json:"content" valid:"required,alphanum"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type UpdateBerita struct {
	Title             string `json:"title" valid:"required,alphanum"`
	SubTitle          string `json:"sub_title" valid:"required,alphanum"`
	Content           string `json:"content" valid:"required,alphanum"`
	Tag               string `json:"tag"`
	ThumbnailImageURL string `json:"thumbnail_image_url" valid:"url"`
	Description       string `json:"description"`
}

type BeritaQuery struct {
	Limit         string              `json:"limit" valid:"required"`
	Offset        string              `json:"offset" valid:"required"`
	Filter        []BeritaFilterQuery `json:"filters"`
	Order         string              `json:"order"`
	Sort          string              `json:"sort"`
	CreatedAtFrom string              `json:"created_at_from"`
	CreatedAtTo   string              `json:"created_at_to"`
	PublishAtFrom string              `json:"publish_at_from"`
	PublishAtTo   string              `json:"publish_at_to"`
}

type BeritaFilterQuery struct {
	Field   string `json:"field"`
	Keyword string `json:"keyword"`
}

type ReadBerita struct {
	IDPPCPBerita      string            `json:"id"`
	Title             string            `json:"title"`
	SubTitle          string            `json:"sub_title"`
	Tag               string            `json:"tag"`
	ThumbnailImageURL string            `json:"thumbnail_image_url"`
	Detail            *ReadBeritaDetail `json:"detail,omitempty"`
	Description       string            `json:"description"`
	Status            string            `json:"status"`
	CreatedAt         time.Time         `json:"created_at"`
	CreatedBy         *string           `json:"created_by"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	UpdatedBy         *string           `json:"updated_by"`
	PublishedAt       *time.Time        `json:"published_at"`
	PublishedBy       *string           `json:"published_by"`
	IsDeleted         bool              `json:"is_deleted"`
}

type ReadBeritaDetail struct {
	Content string `json:"content"`
}

func GetCreatePayload(body []byte) (payload CreateBerita, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetUpdatePayload(body []byte) (payload UpdateBerita, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func GetQueryPayload(body []byte) (payload BeritaQuery, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func (r CreateBerita) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r UpdateBerita) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r BeritaQuery) Validate() (err error) {
	// Validate Payload
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return
	}
	return
}

func (r CreateBerita) ToEntity() (data entity.BeritaEntity, dataDetail entity.BeritaDetailEntity) {
	data = entity.BeritaEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
		CreatedAt:         time.Now(),
	}

	dataDetail = entity.BeritaDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r UpdateBerita) ToEntity() (data entity.BeritaEntity, dataDetail entity.BeritaDetailEntity) {
	data = entity.BeritaEntity{
		Title:             r.Title,
		SubTitle:          r.SubTitle,
		Tag:               r.Tag,
		ThumbnailImageURL: r.ThumbnailImageURL,
		Description:       r.Description,
	}

	dataDetail = entity.BeritaDetailEntity{
		Content: r.Content,
		Tag:     r.Tag,
	}
	return
}

func (r BeritaQuery) ToEntity() (data entity.BeritaQueryEntity) {
	filters := []entity.BeritaFilterQueryEntity{}
	for _, fil := range r.Filter {
		filter := entity.BeritaFilterQueryEntity{
			Field:   fil.Field,
			Keyword: fil.Keyword,
		}
		filters = append(filters, filter)
	}
	data = entity.BeritaQueryEntity{
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

func ToPayload(data entity.BeritaEntity, isDetail bool) (response ReadBerita) {
	var detail *ReadBeritaDetail
	if isDetail {
		detail = &ReadBeritaDetail{
			Content: data.Detail.Content,
		}
	}
	response = ReadBerita{
		IDPPCPBerita:      data.IDPPCPBerita,
		Title:             data.Title,
		SubTitle:          data.SubTitle,
		Tag:               data.Tag,
		ThumbnailImageURL: data.ThumbnailImageURL,
		Detail:            detail,
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
